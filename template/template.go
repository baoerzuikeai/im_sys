package template

import (
	"bytes"
	"html/template"
	"log"

	"github.com/baoer/im_sys/util"
)

// sendcode template
func ParseSendCodetemplate() []byte {
	type TemplateData struct {
		Code string
	}
	data := TemplateData{
		Code: util.CreateCode(),
	}

	// 解析和执行模板

	tmpl, err := template.ParseFiles("./template/email.html")
	if err != nil {
		log.Fatalf("解析模板失败: %v", err)
	}

	// 渲染模板
	var tplBuffer bytes.Buffer
	err = tmpl.Execute(&tplBuffer, data)
	if err != nil {
		log.Fatalf("渲染模板失败: %v", err)
	}

	return tplBuffer.Bytes()
}
