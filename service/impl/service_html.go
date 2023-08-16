package impl

import (
	"SQLResultExport/tool"
	"bytes"
	"text/template"
)

type ExportHtmlService struct {
}

func (service *ExportHtmlService) Export(ds []map[string]string) (string, error) {
	tmpl := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<table border="1" style="border-collapse:collapse">
{{range $_, $row := .ds}}
<tr>
{{range $key, $value := $row}}
  <td>{{$value}}</td>
{{end}}
</tr>
{{end}}
</table>
</body>
</html>`

	t, err := template.New("html").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	data := make(map[string]any)
	data["ds"] = ds
	err = t.Execute(buf, data)
	if err != nil {
		return "", err
	}

	var fileName = "Export.html"
	err = tool.WriteFile(fileName, buf.String())
	if err != nil {
		return "", err
	}
	return fileName, nil
}
