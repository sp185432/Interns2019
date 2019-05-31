package main
import "fmt"
func main() {

    s := make([]string, 3)
    fmt.Println("emp:", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("append:", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    l := s[2:5]
    fmt.Println("slice1:", l)

    l = s[:5]
    fmt.Println("slice2:", l)

    l = s[2:]
    fmt.Println("slice3:", l)
//initialization and declaration
    t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
}