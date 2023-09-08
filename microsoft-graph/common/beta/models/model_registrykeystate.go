package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryKeyState struct {
	Hive         *RegistryKeyStateHive      `json:"hive,omitempty"`
	Key          *string                    `json:"key,omitempty"`
	ODataType    *string                    `json:"@odata.type,omitempty"`
	OldKey       *string                    `json:"oldKey,omitempty"`
	OldValueData *string                    `json:"oldValueData,omitempty"`
	OldValueName *string                    `json:"oldValueName,omitempty"`
	Operation    *RegistryKeyStateOperation `json:"operation,omitempty"`
	ProcessId    *int64                     `json:"processId,omitempty"`
	ValueData    *string                    `json:"valueData,omitempty"`
	ValueName    *string                    `json:"valueName,omitempty"`
	ValueType    *RegistryKeyStateValueType `json:"valueType,omitempty"`
}
