package dscconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DscConfigurationProperties struct {
	CreationTime           *string                               `json:"creationTime,omitempty"`
	Description            *string                               `json:"description,omitempty"`
	JobCount               *int64                                `json:"jobCount,omitempty"`
	LastModifiedTime       *string                               `json:"lastModifiedTime,omitempty"`
	LogVerbose             *bool                                 `json:"logVerbose,omitempty"`
	NodeConfigurationCount *int64                                `json:"nodeConfigurationCount,omitempty"`
	Parameters             *map[string]DscConfigurationParameter `json:"parameters,omitempty"`
	ProvisioningState      *DscConfigurationProvisioningState    `json:"provisioningState,omitempty"`
	Source                 *ContentSource                        `json:"source,omitempty"`
	State                  *DscConfigurationState                `json:"state,omitempty"`
}
