package integrationserviceenvironmentmanagedapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiReference struct {
	BrandColor                    *string            `json:"brandColor,omitempty"`
	Category                      *ApiTier           `json:"category,omitempty"`
	Description                   *string            `json:"description,omitempty"`
	DisplayName                   *string            `json:"displayName,omitempty"`
	IconUri                       *string            `json:"iconUri,omitempty"`
	Id                            *string            `json:"id,omitempty"`
	IntegrationServiceEnvironment *ResourceReference `json:"integrationServiceEnvironment,omitempty"`
	Name                          *string            `json:"name,omitempty"`
	Swagger                       *interface{}       `json:"swagger,omitempty"`
	Type                          *string            `json:"type,omitempty"`
}
