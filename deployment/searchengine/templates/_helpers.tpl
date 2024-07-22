


{{/*
Expand the name of the chart.
*/}}
{{- define "readApi.name" -}}
{{- default .Chart.Name .Values.readApi.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}



{{- define "readApi.chart" -}}
{{- printf "%s-%s" .Values.readApi.name  .Values.readApi.version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}



{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "readApi.fullname" -}}
{{- if .Values.readApi.fullnameOverride }}
{{- .Values.readApi.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.readApi.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}


{{/*
Selector labels
*/}}
{{- define "readApi.selectorLabels" -}}
app.kubernetes.io/name: {{ include "readApi.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}


{{/*
Common labels
*/}}
{{- define "readApi.labels" -}}
helm.sh/chart: {{ include "readApi.chart" . }}
{{ include "readApi.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}


{{/*
Create the name of the service account to use
*/}}
{{- define "readApi.serviceAccountName" -}}
{{- if .Values.readApi.serviceAccount.create }}
{{- default (include "readApi.fullname" .) .Values.readApi.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.readApi.serviceAccount.name }}
{{- end }}
{{- end }}








{{/*
Expand the name of the chart.
*/}}
{{- define "writeApi.name" -}}
{{- default .Chart.Name .Values.writeApi.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}



{{- define "writeApi.chart" -}}
{{- printf "%s-%s" .Values.writeApi.name  .Values.writeApi.version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}



{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "writeApi.fullname" -}}
{{- if .Values.writeApi.fullnameOverride }}
{{- .Values.writeApi.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.writeApi.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}


{{/*
Selector labels
*/}}
{{- define "writeApi.selectorLabels" -}}
app.kubernetes.io/name: {{ include "writeApi.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}


{{/*
Common labels
*/}}
{{- define "writeApi.labels" -}}
helm.sh/chart: {{ include "writeApi.chart" . }}
{{ include "writeApi.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}


{{/*
Create the name of the service account to use
*/}}
{{- define "writeApi.serviceAccountName" -}}
{{- if .Values.writeApi.serviceAccount.create }}
{{- default (include "writeApi.fullname" .) .Values.writeApi.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.writeApi.serviceAccount.name }}
{{- end }}
{{- end }}
























{{/*
Expand the name of the chart.
*/}}
{{- define "app.name" -}}
{{- default .Chart.Name .Values.app.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}



{{- define "app.chart" -}}
{{- printf "%s-%s" .Values.app.name  .Values.app.version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}



{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "app.fullname" -}}
{{- if .Values.readApi.fullnameOverride }}
{{- .Values.readApi.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Values.app.name .Values.app.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}


{{/*
Selector labels
*/}}
{{- define "app.selectorLabels" -}}
app.kubernetes.io/name: {{ include "app.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}


{{/*
Common labels
*/}}
{{- define "app.labels" -}}
helm.sh/chart: {{ include "app.chart" . }}
{{ include "app.selectorLabels" . }}
{{- if .Values.app.version }}
app.kubernetes.io/version: {{ .Values.app.version | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}


{{/*
Create the name of the service account to use
*/}}
{{- define "app.serviceAccountName" -}}
{{- if .Values.readApi.serviceAccount.create }}
{{- default (include "app.fullname" .) .Values.app.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.readApi.serviceAccount.name }}
{{- end }}
{{- end }}








