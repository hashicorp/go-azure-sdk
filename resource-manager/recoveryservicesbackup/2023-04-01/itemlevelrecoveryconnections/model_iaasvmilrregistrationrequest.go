package itemlevelrecoveryconnections

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ILRRequest = IaasVMILRRegistrationRequest{}

type IaasVMILRRegistrationRequest struct {
	InitiatorName             *string `json:"initiatorName,omitempty"`
	RecoveryPointId           *string `json:"recoveryPointId,omitempty"`
	RenewExistingRegistration *bool   `json:"renewExistingRegistration,omitempty"`
	VirtualMachineId          *string `json:"virtualMachineId,omitempty"`

	// Fields inherited from ILRRequest
}

var _ json.Marshaler = IaasVMILRRegistrationRequest{}

func (s IaasVMILRRegistrationRequest) MarshalJSON() ([]byte, error) {
	type wrapper IaasVMILRRegistrationRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IaasVMILRRegistrationRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IaasVMILRRegistrationRequest: %+v", err)
	}
	decoded["objectType"] = "IaasVMILRRegistrationRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IaasVMILRRegistrationRequest: %+v", err)
	}

	return encoded, nil
}
