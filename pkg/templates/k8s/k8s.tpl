# This is the kubernetes template
# Customize this file as needed
{{define "deployment.yaml"}}
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .AppName }}
  labels:
    app: {{ .AppName }}
spec:
  replicas: {{ .Replicas }}
  selector:
    matchLabels:
      app: {{ .AppName }}
  template:
    metadata:
      labels:
        app: {{ .AppName }}
    spec:
      containers:
      - name: {{ .AppName }}
        image: {{ .Image }}
        ports:
        - containerPort: {{ .Port }}
{{end}}

{{define "service.yaml"}}
apiVersion: v1
kind: Service
metadata:
  name: {{ .AppName }}
spec:
  type: {{ .ServiceType }}
  selector:
    app: {{ .AppName }}
  ports:
  - protocol: TCP
    port: {{ .Port }}
    targetPort: {{ .Port }}
{{end}}

{{define "configmap.yaml"}}
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .AppName }}-config
data:
  APP_ENV: "production"
  APP_DEBUG: "false"
{{end}}



