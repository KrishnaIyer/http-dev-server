{{/*
Expand the name of the chart.
*/}}
{{- define "http-dev-server.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "http-dev-server.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "http-dev-server.labels" -}}
http-dev-server.sh/chart: {{ include "http-dev-server.chart" . }}
{{ include "http-dev-server.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "http-dev-server.selectorLabels" -}}
app.kubernetes.io/name: {{ include "http-dev-server.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}
