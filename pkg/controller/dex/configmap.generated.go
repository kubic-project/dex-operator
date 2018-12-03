// +build !ignore_autogenerated

/*
Copyright 2018 SUSE LINUX GmbH, Nuernberg, Germany..

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated from "configmap.yaml.in". DO NOT EDIT.

package dex

const (
	configMapTemplate = `
kind: ConfigMap
apiVersion: v1
metadata:
  name: {{ .DexName }}
  namespace: {{ .DexNamespace }}
data:
  {{ .DexConfigMapFilename | basename }}: |
    issuer: "https://{{ .DexAddress }}:{{ .DexPort }}"

    storage:
      type: kubernetes
      config:
        inCluster: true
    web:
      https: 0.0.0.0:5556
      tlsCert: {{ .DexCertsDir }}/tls.crt
      tlsKey: {{ .DexCertsDir }}/tls.key

    frontend:
      dir: /usr/share/caasp-dex/web
      theme: caasp

{{- if .LDAPConnectors }}
    connectors:
  {{- range $Con := .LDAPConnectors }}
    - type: ldap
      id: {{ $Con.Spec.ID }}
      name: {{ $Con.Spec.Name }}
      config:
        # Host and optional port of the LDAP server in the form "host:port".
        # If the port is not supplied, it will be guessed based on "insecureNoSSL",
        # and "startTLS" flags. 389 for insecure or StartTLS connections, 636
        # otherwise.
        host: {{ $Con.Spec.Server }}

	  {{- if $Con.Spec.StartTLS }}
        # When connecting to the server, connect using the ldap:// protocol then issue
        # a StartTLS command. If unspecified, connections will use the ldaps:// protocol
        startTLS: {{ $Con.Spec.StartTLS }}
    {{- end }}

	  {{- if and $Con.Spec.BindDN $Con.Spec.BindPW }}
        # The DN and password for an application service account. The connector uses
        # these credentials to search for users and groups. Not required if the LDAP
        # server provides access for anonymous auth.
        # Please note that if the bind password contains a $, it has to be saved in an
        # environment variable which should be given as the value to bindPW.
        # bindDN: uid=seviceaccount,cn=users,dc=example,dc=com
        # bindPW: password
        bindDN: {{ $Con.Spec.BindDN }}
        bindPW: {{ $Con.Spec.BindPW }}
    {{- else }}
        # bindDN and bindPW not present; anonymous bind will be used
    {{- end }}

	  {{- if $Con.Spec.UsernamePrompt }}
        usernamePrompt: {{ $Con.Spec.UsernamePrompt }}
	  {{- end }}

	  {{- if $Con.Spec.RootCAData }}
        # A raw certificate file can also be provided inline.
        rootCAData: {{ $Con.Spec.RootCAData | replace "\n" "" | base64encode }}
      {{- else }}
        # Path to a trusted root certificate file.
        # The CA certificate bundle is automatically mounted into pods using
        # the default service account at this path. If you are not using the default
        # service account, we should build a configmap containing
        # the certificate bundle that you have access to use...
        rootCA: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
      {{- end }}
    {{- if $Con.Spec.User.BaseDN }}
        userSearch:
          # BaseDN to start the search from. It will translate to the query
          # "(&(objectClass=person)(uid=<username>))".
          baseDN: {{ $Con.Spec.User.BaseDN }}

          # Optional filter to apply when searching the directory.
          filter: {{ $Con.Spec.User.Filter }}

          # username attribute used for comparing user entries. This will be translated
          # and combined with the other filter as "(<attr>=<username>)".
          username: {{ $Con.Spec.User.Username }}
          idAttr: {{ $Con.Spec.User.IDAttr }}
		  {{- if $Con.Spec.User.EmailAttr }}
          # Required. Attribute to map to Email.
          emailAttr: {{ $Con.Spec.User.EmailAttr }}
		  {{- end }}

      {{- if $Con.Spec.User.NameAttr }}
          # Maps to display name of users. No default value.
          nameAttr: {{ $Con.Spec.User.NameAttr }}
		  {{- end }}
    {{- end }}

    {{- if $Con.Spec.Group.BaseDN }}
        # Group search queries for groups given a user entry.
        groupSearch:
          # BaseDN to start the search from. It will translate to the query
          # "(&(objectClass=group)(member=<user uid>))".
          baseDN: {{ $Con.Spec.Group.BaseDN }}

          # Optional filter to apply when searching the directory.
          filter: {{ $Con.Spec.Group.Filter }}

          # Following two fields are used to match a user to a group. It adds an additional
          # requirement to the filter that an attribute in the group must match the user's
          # attribute value.
          userAttr: {{ $Con.Spec.Group.UserAttr }}
          groupAttr: {{ $Con.Spec.Group.GroupAttr }}

    	{{- if $Con.Spec.Group.NameAttr }}
          # Represents group name.
          nameAttr: {{ $Con.Spec.Group.NameAttr }}
		  {{- end }}

    {{- end }}
  {{- end }}
{{- end }}

    oauth2:
      skipApprovalScreen: true

    staticClients:
    # The 'name' must match the k8s API server's 'oidc-client-id'
    - id: kubernetes
      redirectURIs:
      - 'urn:ietf:wg:oauth:2.0:oob'
      name: "Kubernetes"
      secret: "{{ index .DexSharedPasswords "kubernetes" }}"

{{- if .StaticClients }}
      trustedPeers:
    {{- range $Client := .StaticClients }}
      {{- $id = $Client.Name | safeYAMLId }}
      - {{ $id }}
    {{- end }}

  {{- range $Client := .StaticClients }}
    {{- $id = $Client.Name | safeYAMLId }}
    - id: {{ $id }}
      redirectURIs:
      {{- range $URL := $Client.RedirectURLs }}
        - '{{ $URL }}'
      {{- end }}
      name: "{{ $Client.Name }}"
      secret: "{{ index .DexSharedPasswords $id }}"
      public: true
  {{- end }}

{{- end }}
`)

