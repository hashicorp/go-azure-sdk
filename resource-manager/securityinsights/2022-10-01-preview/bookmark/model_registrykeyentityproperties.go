package bookmark

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryKeyEntityProperties struct {
	AdditionalData *map[string]interface{} `json:"additionalData,omitempty"`
	FriendlyName   *string                 `json:"friendlyName,omitempty"`
	Hive           *RegistryHive           `json:"hive,omitempty"`
	Key            *string                 `json:"key,omitempty"`
}
