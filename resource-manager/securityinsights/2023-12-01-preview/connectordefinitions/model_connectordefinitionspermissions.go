package connectordefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorDefinitionsPermissions struct {
	Customs          *[]CustomPermissionDetails              `json:"customs,omitempty"`
	Licenses         *[]string                               `json:"licenses,omitempty"`
	ResourceProvider *[]ConnectorDefinitionsResourceProvider `json:"resourceProvider,omitempty"`
	Tenant           *[]string                               `json:"tenant,omitempty"`
}
