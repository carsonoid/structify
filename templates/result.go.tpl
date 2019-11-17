package main

import (
	{{- range $k, $v := .Metadata.PackageAliases }}
	{{ $k }} "{{ $v }}"
	{{- end }}
)
{{ range $k, $v := .Metadata.PointerAliases }}
{{ $k }} := {{ $v }}
{{- end }}

r := {{ .Dump }}
