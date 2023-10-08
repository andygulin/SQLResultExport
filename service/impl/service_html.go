package impl

import (
	. "SQLResultExport/service"
	"SQLResultExport/tool"
	"bytes"
	"path/filepath"
	"text/template"
)

type ExportHtmlService struct {
}

func (service *ExportHtmlService) Export(rs ResultSet) (File, error) {
	tmpl := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<table border="1" style="border-collapse:collapse">
{{range $_, $row := .rs}}
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
	data["rs"] = rs
	err = t.Execute(buf, data)
	if err != nil {
		return "", err
	}

	var fileName = "Export.html"
	err = tool.WriteFile(fileName, buf.String())
	if err != nil {
		return "", err
	}

	output, _ := filepath.Abs(fileName)
	return File(output), nil
}
