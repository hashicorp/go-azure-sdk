package schemaversions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SchemaVersionProperties struct {
	Description       *string            `json:"description,omitempty"`
	Hash              *string            `json:"hash,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	SchemaContent     string             `json:"schemaContent"`
	Uuid              *string            `json:"uuid,omitempty"`
}
