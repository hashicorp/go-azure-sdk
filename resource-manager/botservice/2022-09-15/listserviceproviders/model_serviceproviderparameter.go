package listserviceproviders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceProviderParameter struct {
	Default     *string                           `json:"default,omitempty"`
	Description *string                           `json:"description,omitempty"`
	DisplayName *string                           `json:"displayName,omitempty"`
	HelpURL     *string                           `json:"helpUrl,omitempty"`
	Metadata    *ServiceProviderParameterMetadata `json:"metadata,omitempty"`
	Name        *string                           `json:"name,omitempty"`
	Type        *string                           `json:"type,omitempty"`
}
