// Small hello world program not using fmt package as a test to see how big that is.
package main

import "os"

func main() {
	os.Stdout.WriteString("Simple Hello")
}
