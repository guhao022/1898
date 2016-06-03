package doc

import (
	"go/token"
	"go/parser"
	"go/doc"
	"regexp"
	"strings"
)

// 文档结构
type Doc struct {
	path string 	//需要生成文档的文件夹
	Host 	string    `json:"host"`
	Port string    `json:"port"`
	Comments []Comment    `json:"comments"`	//文档结构
}

type Comment struct {
	Name	string        `json:"name"`
	Method	string        `json:"method"`
	Uri		string        `json:"uri"`
	Params	map[string]string    `json:"params"`
	Response 	map[string]string    `json:"response"`
}

func New(path string) *Doc {
	return &Doc{path:path, Host:"127.0.0.1", Port:"9900"}
}

func (d *Doc) SetHost(host, port string) *Doc {
	d.Host = host
	d.Port = port

	return d
}

// 获取注释
func (d *Doc) Analyze() *Doc {

	fset := token.NewFileSet()

	astPkgs, err := parser.ParseDir(fset, d.path, nil, parser.ParseComments)

	if err != nil {
		panic(err)
	}

	var fcs []Comment

	for _, v := range astPkgs {

		p := doc.New(v, d.path, 0)

		for _, t := range p.Funcs {

			fc := match(t.Doc, "name", "uri", "method", "param", "response")

			fcs = append(fcs, fc)
		}
	}

	d.Comments = fcs

	return d

}

func match(doc string, matchs ...string) Comment {
	if len(matchs) < 1 {
		matchs = []string{"name"}
	}

	var co Comment

	for _, m := range matchs {

		reg := regexp.MustCompile("@" + m + "{1} (.*)")
		vals := reg.FindAllString(doc, -1)

		for _, v := range vals {

			kVal := strings.SplitAfterN(v, " ", 2)[1]

			if m == "name" {
				co.Name = kVal
			}

			if m == "method" {
				co.Method = kVal
			}

			if m == "uri" {
				co.Uri = kVal
			}

			if m == "param" {
				co.Params = comp(kVal)
			}

			if m == "response" {
				co.Response = comp(kVal)
			}

		}

	}

	return co
}

func comp(p string) map[string]string {
	var pm = make(map[string]string)

	ps := strings.SplitAfterN(p, " ", 2)

	pm[ps[0]] = ps[1]

	return pm
}

