package restores

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RestoreRequest = AzureFileShareRestoreRequest{}

type AzureFileShareRestoreRequest struct {
	CopyOptions        *CopyOptions          `json:"copyOptions,omitempty"`
	RecoveryType       *RecoveryType         `json:"recoveryType,omitempty"`
	RestoreFileSpecs   *[]RestoreFileSpecs   `json:"restoreFileSpecs,omitempty"`
	RestoreRequestType *RestoreRequestType   `json:"restoreRequestType,omitempty"`
	SourceResourceId   *string               `json:"sourceResourceId,omitempty"`
	TargetDetails      *TargetAFSRestoreInfo `json:"targetDetails,omitempty"`

	// Fields inherited from RestoreRequest

	ObjectType string `json:"objectType"`
}

func (s AzureFileShareRestoreRequest) RestoreRequest() BaseRestoreRequestImpl {
	return BaseRestoreRequestImpl{
		ObjectType: s.ObjectType,
	}
}

var _ json.Marshaler = AzureFileShareRestoreRequest{}

func (s AzureFileShareRestoreRequest) MarshalJSON() ([]byte, error) {
	type wrapper AzureFileShareRestoreRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureFileShareRestoreRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureFileShareRestoreRequest: %+v", err)
	}

	decoded["objectType"] = "AzureFileShareRestoreRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureFileShareRestoreRequest: %+v", err)
	}

	return encoded, nil
}
