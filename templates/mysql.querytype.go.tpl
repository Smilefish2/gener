{{- $table := (schema .Schema .Table.TableName) -}}
{{- if .Comment -}}
// {{ .Comment }}
{{- else -}}
// {{ .Name }} represents a row from '{{ $table }}'.
{{- end }}
//proteus:generate
type {{ .Name }} struct {
{{- range .Fields }}
	{{ .Name }} {{ retype .Type }} // {{ .Col.ColumnName }}
{{- end }}
}

