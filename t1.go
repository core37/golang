package main
import (
    "fmt"
    "math/rand"
    "time"
)
var s[100] int;

func qs(l, r int){
    if (l >= r) { return; }
    var i, j, pivot int = l, r, s[l];
    for i < j{
        for (i < j && s[j] >= pivot){ j--; }
        s[i] = s[j];
        for (i < j && s[i] <= pivot) { i++; }
        s[j] = s[i];
    }
    s[i] = pivot;
    qs(l, i - 1);
    qs(i + 1, r);
}
func main() {
    var i int;
    rand.Seed(time.Now().UnixNano())
    for i = 0; i < 100; i++ {
        s[i] = rand.Intn(100);
    }
    qs(0, 99);
    for i = 0; i < 100; i++ {
        fmt.Printf("%d ", s[i]);
    }

}