package invoice

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os/exec"
	"time"
)

// creating template from local "html-template.html" file
var myHtmlTemplate, _ = template.ParseFiles("./template.html")

func Generate() error {
	// creating a buffer
	buff := bytes.NewBufferString("")

	// compiling the template into the buffer
	err := myHtmlTemplate.Execute(buff, "Hello")
	if err != nil {
		return err
	}

	// writing the compiled template to the file
	err = ioutil.WriteFile("compiled.html", buff.Bytes(), 0666)
	if err != nil {
		return err
	}

	// converting "compiled.html" file to "my.pdf"
	name := fmt.Sprintf("%d%d.pdf", time.Now().Unix(), time.Nanosecond)
	return exec.Command("wkhtmltopdf", "compiled.html", name).Run()
}
