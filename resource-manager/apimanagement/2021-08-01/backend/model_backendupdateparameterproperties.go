package backend

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackendUpdateParameterProperties struct {
	Credentials *BackendCredentialsContract `json:"credentials"`
	Description *string                     `json:"description,omitempty"`
	Properties  *BackendProperties          `json:"properties"`
	Protocol    *BackendProtocol            `json:"protocol,omitempty"`
	Proxy       *BackendProxyContract       `json:"proxy"`
	ResourceId  *string                     `json:"resourceId,omitempty"`
	Title       *string                     `json:"title,omitempty"`
	Tls         *BackendTlsProperties       `json:"tls"`
	Url         *string                     `json:"url,omitempty"`
}
