package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ StorageAccountCustomDetails = ExistingStorageAccount{}

type ExistingStorageAccount struct {
	AzureStorageAccountId string `json:"azureStorageAccountId"`

	// Fields inherited from StorageAccountCustomDetails
}

var _ json.Marshaler = ExistingStorageAccount{}

func (s ExistingStorageAccount) MarshalJSON() ([]byte, error) {
	type wrapper ExistingStorageAccount
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExistingStorageAccount: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExistingStorageAccount: %+v", err)
	}
	decoded["resourceType"] = "Existing"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExistingStorageAccount: %+v", err)
	}

	return encoded, nil
}
