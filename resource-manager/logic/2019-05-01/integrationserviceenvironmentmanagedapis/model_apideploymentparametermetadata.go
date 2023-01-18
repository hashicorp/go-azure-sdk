package integrationserviceenvironmentmanagedapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiDeploymentParameterMetadata struct {
	Description *string                           `json:"description,omitempty"`
	DisplayName *string                           `json:"displayName,omitempty"`
	IsRequired  *bool                             `json:"isRequired,omitempty"`
	Type        *string                           `json:"type,omitempty"`
	Visibility  *ApiDeploymentParameterVisibility `json:"visibility,omitempty"`
}
