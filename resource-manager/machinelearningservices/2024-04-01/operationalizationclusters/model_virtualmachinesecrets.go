package operationalizationclusters

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ComputeSecrets = VirtualMachineSecrets{}

type VirtualMachineSecrets struct {
	AdministratorAccount *VirtualMachineSshCredentials `json:"administratorAccount,omitempty"`

	// Fields inherited from ComputeSecrets

	ComputeType ComputeType `json:"computeType"`
}

func (s VirtualMachineSecrets) ComputeSecrets() BaseComputeSecretsImpl {
	return BaseComputeSecretsImpl{
		ComputeType: s.ComputeType,
	}
}

var _ json.Marshaler = VirtualMachineSecrets{}

func (s VirtualMachineSecrets) MarshalJSON() ([]byte, error) {
	type wrapper VirtualMachineSecrets
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VirtualMachineSecrets: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualMachineSecrets: %+v", err)
	}

	decoded["computeType"] = "VirtualMachine"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualMachineSecrets: %+v", err)
	}

	return encoded, nil
}
