package main
import (
     "fmt"
     "github.com/hardyantz/go-hello-world/ext"
)
func main() {
     fmt.Println("hello world")
     count := ext.Counting(4, 2)
     fmt.Println(count)
}