package main

import (
	"fmt"
	"github.com/Joker/jade"
	beego "github.com/beego/beego/v2/server/web"
	"html/template"
	//"os"
	"path/filepath"
	//"strings"
	_ "web_00/routers"
	"web_00/utils"
)

func addJadeTemplate() {
	beego.AddTemplateEngine("jade", func(root, path string, funcs template.FuncMap) (*template.Template, error) {
		jadePath := filepath.Join(root, path)
		content, err := utils.ReadFile(jadePath)
		//fmt.Println(content)
		if err != nil {
			return nil, fmt.Errorf("error loading jade template: %v", err)
		}
		tpl, err := jade.Parse("name_of_tpl", content)
		if err != nil {
			return nil, fmt.Errorf("error loading jade template: %v", err)
		}
		//fmt.Println("html:\n%s", tpl)
		tmp := template.New("Person template")
		tmp, err = tmp.Parse(tpl)
		if err != nil {
			return nil, fmt.Errorf("error loading jade template: %v", err)
		}
		//fmt.Println(tmp)
		return tmp, err

	})
}

func main() {
	addJadeTemplate()
	beego.Run()
}
