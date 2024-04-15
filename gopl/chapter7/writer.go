package main

import (
	"io"
)

// writeString writes s to w.
// If w has a WriteString method, it is invoked instead of w.Write.
func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	if sw, ok := w.(stringWriter); ok {
		return sw.WriteString(s) // avoid a copy
	}
	return w.Write([]byte(s)) // allocate temporary copy
}

func writeHeader(w io.Writer, contentType string) error {
	if _, err := writeString(w, "Content-Type: "); err != nil {
		return err
	}
	if _, err := writeString(w, contentType); err != nil {
		return err
	}
	// ...
}

func main() {
	//var w io.Writer
	//w = os.Stdout
	//f := w.(*os.File)      // success: f == os.Stdout
	//c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
	//fmt.Println(f)
	//fmt.Println(c)

	//var w io.Writer
	//w = os.Stdout
	//rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
	//w = new(ByteCounter)
	//rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method
	//
	//fmt.Println(w)
	//fmt.Println(rw)

}
