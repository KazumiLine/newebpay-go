package newebpay

import (
	"bytes"
	"html/template"
)

var OrderTemplateText = `<form id="order_form" action="{{.Action}}" method="post">{{range $key,$element := .Values}}<input type="hidden" name="{{$key}}" id="{{$key}}" value="{{$element}}" />{{end -}}</form><script>document.querySelector("#order_form").submit();</script>`
var OrderTmpl = template.Must(template.New("AutoPostOrder").Parse(OrderTemplateText))

func GenerateAutoSubmitHtmlForm(targetUrl string, params map[string]string) (string, error) {
	var result bytes.Buffer
	err := OrderTmpl.Execute(&result, struct {
		Values map[string]string
		Action string
	}{
		Values: params,
		Action: targetUrl,
	})
	if err != nil {
		return "", err
	}
	return result.String(), nil
}
