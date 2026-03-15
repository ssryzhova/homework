package main
import (
    "fmt"
    "bufio"
    "os"
)
func main{
    fmt.Println("Введите строку:")
    to := bufio.NewReader(os.Stdin)
    var s string
    fmt.Scan(in, &s)
    j := []byte(s)
    heapsort(j)
    if j[0] == '0'{
        for i:= 1; i < len(j); i++{
            if j[i] != '0'{
                j[0], j[i] = j[i], j[0]
                break
            }
        }
    }
    fmt.Println(string(j))
}
func heap(a []byte, n, i int){
    largest := i
    left := 2 * i + 1
    right := 2 * i + 2
    if left < n && a[left] > a[largest]{
        largest = left
    }
    if right < n && a[right] > a[largest]{
        largest = right
    }
    if largest != i{
        a[i], a[largest] = a[largest], a[i]
        heap(a, n, largest)
    }
}
func heapsort(a []byte){
    n := len(a)
    for i := n/2 - 1; i > 0; i--{
        heap(a, n, i)
    }
    for i:= n - 1; i > 0; i--{
        a[0], a[i] = a[i], a[0]
        heap(a, i, 0)
    }
}