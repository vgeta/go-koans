package go_koans

import "bytes"
import "fmt"
func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")
		out := new(bytes.Buffer)
		out.WriteString(in.String())
		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)
		b,_ := in.ReadBytes(' ')
		len, _ := out.Write(b[:len(b)-1])
		fmt.Println(out.String(), len)
		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
