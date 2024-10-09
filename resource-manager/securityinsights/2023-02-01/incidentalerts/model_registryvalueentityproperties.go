package incidentalerts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryValueEntityProperties struct {
	AdditionalData *interface{}       `json:"additionalData,omitempty"`
	FriendlyName   *string            `json:"friendlyName,omitempty"`
	KeyEntityId    *string            `json:"keyEntityId,omitempty"`
	ValueData      *string            `json:"valueData,omitempty"`
	ValueName      *string            `json:"valueName,omitempty"`
	ValueType      *RegistryValueKind `json:"valueType,omitempty"`
}
