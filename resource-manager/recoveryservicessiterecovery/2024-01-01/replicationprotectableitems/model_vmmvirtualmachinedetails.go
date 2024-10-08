package replicationprotectableitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ConfigurationSettings = VMmVirtualMachineDetails{}

type VMmVirtualMachineDetails struct {

	// Fields inherited from ConfigurationSettings

	InstanceType string `json:"instanceType"`
}

func (s VMmVirtualMachineDetails) ConfigurationSettings() BaseConfigurationSettingsImpl {
	return BaseConfigurationSettingsImpl{
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = VMmVirtualMachineDetails{}

func (s VMmVirtualMachineDetails) MarshalJSON() ([]byte, error) {
	type wrapper VMmVirtualMachineDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VMmVirtualMachineDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VMmVirtualMachineDetails: %+v", err)
	}

	decoded["instanceType"] = "VmmVirtualMachine"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VMmVirtualMachineDetails: %+v", err)
	}

	return encoded, nil
}
