package schemaregistry

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SchemaGroupProperties struct {
	CreatedAtUtc        *string              `json:"createdAtUtc,omitempty"`
	ETag                *string              `json:"eTag,omitempty"`
	GroupProperties     *map[string]string   `json:"groupProperties,omitempty"`
	SchemaCompatibility *SchemaCompatibility `json:"schemaCompatibility,omitempty"`
	SchemaType          *SchemaType          `json:"schemaType,omitempty"`
	UpdatedAtUtc        *string              `json:"updatedAtUtc,omitempty"`
}
