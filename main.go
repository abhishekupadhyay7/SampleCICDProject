package main

import "fmt"

func main() {
    msg := printHi("Abhishek.")
    fmt.Println(msg)
}

func printHi(name string) string {
    return fmt.Sprintf("Hello %s", name)
}