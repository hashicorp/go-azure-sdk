package itemlevelrecoveryconnections

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ILRRequest = AzureFileShareProvisionILRRequest{}

type AzureFileShareProvisionILRRequest struct {
	RecoveryPointId  *string `json:"recoveryPointId,omitempty"`
	SourceResourceId *string `json:"sourceResourceId,omitempty"`

	// Fields inherited from ILRRequest
}

var _ json.Marshaler = AzureFileShareProvisionILRRequest{}

func (s AzureFileShareProvisionILRRequest) MarshalJSON() ([]byte, error) {
	type wrapper AzureFileShareProvisionILRRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureFileShareProvisionILRRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureFileShareProvisionILRRequest: %+v", err)
	}
	decoded["objectType"] = "AzureFileShareProvisionILRRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureFileShareProvisionILRRequest: %+v", err)
	}

	return encoded, nil
}
