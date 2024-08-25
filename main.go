
package main

import (
    "fmt"
    "strconv"
    "os"
    "bufio"
)


func main() {


    scan := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter the number of blocks you want to create")
    scan.Scan()
    num, _ := strconv.Atoi(scan.Text())
    
    var Blockchain []Block

    HossineEthBlock := HossineEthBlock()
    Blockchain = append(Blockchain, HossineEthBlock)


   for i := 0; i < num; i++ {
        previousBlock := Blockchain[len(Blockchain)-1]
        newBlock := createblock(previousBlock, "Hossine Eth Block " + strconv.Itoa(i))
        Blockchain = append(Blockchain, newBlock)
    }
    for _, block := range Blockchain {
        fmt.Printf("Index: %d\n", block.Index)
        fmt.Printf("Timestamp: %s\n", block.Timestamp)
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("PreviousHash: %s\n", block.PreviousHash)
        fmt.Printf("Hash: %s\n", block.Hash)
        fmt.Println()
    }
}
