package connection

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionProperties struct {
	ConnectionType        *ConnectionTypeAssociationProperty `json:"connectionType,omitempty"`
	CreationTime          *string                            `json:"creationTime,omitempty"`
	Description           *string                            `json:"description,omitempty"`
	FieldDefinitionValues *map[string]string                 `json:"fieldDefinitionValues,omitempty"`
	LastModifiedTime      *string                            `json:"lastModifiedTime,omitempty"`
}
