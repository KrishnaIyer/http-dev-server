{{- define "http-dev-server.ingress-route-tls" -}}
apiVersion: traefik.containo.us/v1alpha1
kind: IngressRoute
metadata:
  name: {{ include "http-dev-server.name" . }}-httpsecure
  labels:
    {{- include "http-dev-server.labels" . | nindent 4 }}
spec:
  entryPoints: {{ .Values.global.ingress.httpsecure.entryPoints }}
  routes:
    - kind: Rule
      match: PathPrefix(`/`)
      services:
        - name: {{ include "http-dev-server.name" . }}
          port: 8080
          scheme: http
  tls:
    secretName: {{ required ".Values.ingress.tls.secretName is required!" .Values.ingress.tls.secretName }}
{{- end }}
