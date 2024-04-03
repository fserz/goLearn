package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"

	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "123 123"

}
