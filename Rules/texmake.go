package main

import (
	f "fmt"
	o "os"
)

func main() {
	f.Println("File name: ")
	var name string
	var ext = ".tex"
	_, _ = f.Scan(&name)

    fname := f.Sprintf("%s%s",name,ext)
	file, err := o.Create(fname)
	if err != nil {
		panic(err)
	}
	f.Println("CREATED")
	f.Println("Title: ")
	var title string
	_, _ =f.Scanln(&title)

	start := f.Sprintf("\\documentclass[12pt, letterpaper]{article}\n\\title{%s}\n\\author{Mustafif Khan}\n", title)
	var toc = "\\usepackage{color}\n\\usepackage{hyperref}\n\\hypersetup{\ncolorlinks=true, % make the links colored\nlinkcolor=black, % color TOC links in blue\nurlcolor=red, % color URLs in red\nlinktoc=all % 'all' will create links for everything in the TOC\n}\n"
	var page = "\\begin{document}\n\\maketitle\n\\pagenumbering{arabic}\n\\newpage\n\\tableofcontents\n\\newpage\n\n\n\\end{document}"
	_, _ = file.WriteString(start)
	_, _ = file.WriteString(toc)
	_, _ = file.WriteString(page)
}







