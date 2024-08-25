package main

import (
    "crypto/aes"
    "crypto/cipher"
    "encoding/hex"
    "errors"
    "fmt"
)

// Decrypt decrypts the AES-encrypted data with the given key
func Decrypt(encryptedText string, key []byte) (string, error) {
    // Decode the hex string
    cipherText, err := hex.DecodeString(encryptedText)
    if err != nil {
        return "", err
    }

    // Create AES cipher
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    // Extract the IV from the start of the cipherText
    if len(cipherText) < aes.BlockSize {
        return "", errors.New("ciphertext too short")
    }
    iv := cipherText[:aes.BlockSize]
    cipherText = cipherText[aes.BlockSize:]

    // Create a new AES decryption cipher
    stream := cipher.NewCFBDecrypter(block, iv)

    // Decrypt the data
    decrypted := make([]byte, len(cipherText))
    stream.XORKeyStream(decrypted, cipherText)

    return string(decrypted), nil
}

func main() {
    // Example key and encrypted text
    // Use a proper 32-byte key (256-bit) for AES-256, encoded as hex
    key, _ := hex.DecodeString("6368616e676520746869732070617373")
    encryptedText := "faefebd0347b9ef7621a3033a1d715b7c4189736408873b7efb90911d5ae3cf8" 

    decryptedText, err := Decrypt(encryptedText, key)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Decrypted Text:", decryptedText)
}
