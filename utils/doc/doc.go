package doc

import (
	"go/token"
	"go/parser"
	"fmt"
	"go/doc"
)

// 文档结构
type Doc struct {
	path string 	//需要生成文档的文件夹
	comments []Comments	//文档结构
}

type Comments struct {
	Method	string
	Uri		string
	Params	[]map[string]string
}

var default_path = "./dll"

func New(path ...string) *Doc {
	if len(path) < 1 {
		path = default_path
	}

	return &Doc{path}
}

// 获取注释
func (d *Doc) comment() {

	fset := token.NewFileSet()

	astPkgs, err := parser.ParseDir(fset, d.path, nil, parser.ParseComments)

	if err != nil {
		println(err.Error())
		return
	}

	for k, v := range astPkgs {
		fmt.Println("package", k)

		p := doc.New(v, "./", 0)

		for _, t := range p.Funcs {
			fmt.Println("docs:", t.Doc)
		}
	}

}

func (d *Doc) parser() {

}


