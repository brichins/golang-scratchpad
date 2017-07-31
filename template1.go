package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

type Employee struct {
	Name       string
	EmployeeID int32
}

func main() {
	emp := Employee{"Bob", 123456789}
	tmpl, _ := template.New("t").Parse("Hello {{.}}")
	tmpl.Execute(os.Stdout, emp.Name)

	//or it can use a byte buffer to convert back to a string
	tmpl2, _ := template.New("t").Parse("Hello {{.Name}}")
	buf := new(bytes.Buffer)
	tmpl2.Execute(buf, emp)
	fmt.Println("\nThe template text is ", buf.String())
}
