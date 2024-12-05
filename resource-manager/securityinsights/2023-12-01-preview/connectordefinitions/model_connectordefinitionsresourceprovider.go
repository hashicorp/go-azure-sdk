package connectordefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorDefinitionsResourceProvider struct {
	PermissionsDisplayText string                              `json:"permissionsDisplayText"`
	Provider               string                              `json:"provider"`
	ProviderDisplayName    string                              `json:"providerDisplayName"`
	RequiredPermissions    ResourceProviderRequiredPermissions `json:"requiredPermissions"`
	Scope                  ProviderPermissionsScope            `json:"scope"`
}
