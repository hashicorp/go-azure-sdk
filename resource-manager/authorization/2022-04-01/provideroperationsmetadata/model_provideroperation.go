package provideroperationsmetadata

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProviderOperation struct {
	Description  *string      `json:"description,omitempty"`
	DisplayName  *string      `json:"displayName,omitempty"`
	IsDataAction *bool        `json:"isDataAction,omitempty"`
	Name         *string      `json:"name,omitempty"`
	Origin       *string      `json:"origin,omitempty"`
	Properties   *interface{} `json:"properties,omitempty"`
}
