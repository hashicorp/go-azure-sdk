package bms

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProtectionIntent = AzureResourceProtectionIntent{}

type AzureResourceProtectionIntent struct {
	FriendlyName *string `json:"friendlyName,omitempty"`

	// Fields inherited from ProtectionIntent

	BackupManagementType     *BackupManagementType    `json:"backupManagementType,omitempty"`
	ItemId                   *string                  `json:"itemId,omitempty"`
	PolicyId                 *string                  `json:"policyId,omitempty"`
	ProtectionIntentItemType ProtectionIntentItemType `json:"protectionIntentItemType"`
	ProtectionState          *ProtectionStatus        `json:"protectionState,omitempty"`
	SourceResourceId         *string                  `json:"sourceResourceId,omitempty"`
}

func (s AzureResourceProtectionIntent) ProtectionIntent() BaseProtectionIntentImpl {
	return BaseProtectionIntentImpl{
		BackupManagementType:     s.BackupManagementType,
		ItemId:                   s.ItemId,
		PolicyId:                 s.PolicyId,
		ProtectionIntentItemType: s.ProtectionIntentItemType,
		ProtectionState:          s.ProtectionState,
		SourceResourceId:         s.SourceResourceId,
	}
}

var _ json.Marshaler = AzureResourceProtectionIntent{}

func (s AzureResourceProtectionIntent) MarshalJSON() ([]byte, error) {
	type wrapper AzureResourceProtectionIntent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureResourceProtectionIntent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureResourceProtectionIntent: %+v", err)
	}

	decoded["protectionIntentItemType"] = "AzureResourceItem"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureResourceProtectionIntent: %+v", err)
	}

	return encoded, nil
}
