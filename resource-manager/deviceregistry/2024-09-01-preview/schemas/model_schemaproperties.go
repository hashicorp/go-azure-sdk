package schemas

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SchemaProperties struct {
	Description       *string            `json:"description,omitempty"`
	DisplayName       *string            `json:"displayName,omitempty"`
	Format            Format             `json:"format"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	SchemaType        SchemaType         `json:"schemaType"`
	Tags              *map[string]string `json:"tags,omitempty"`
	Uuid              *string            `json:"uuid,omitempty"`
}
