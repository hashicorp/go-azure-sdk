package backupprotectioncontainers

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProtectionContainer = AzureIaaSClassicComputeVMContainer{}

type AzureIaaSClassicComputeVMContainer struct {
	ResourceGroup         *string `json:"resourceGroup,omitempty"`
	VirtualMachineId      *string `json:"virtualMachineId,omitempty"`
	VirtualMachineVersion *string `json:"virtualMachineVersion,omitempty"`

	// Fields inherited from ProtectionContainer
	BackupManagementType  *BackupManagementType `json:"backupManagementType,omitempty"`
	FriendlyName          *string               `json:"friendlyName,omitempty"`
	HealthStatus          *string               `json:"healthStatus,omitempty"`
	ProtectableObjectType *string               `json:"protectableObjectType,omitempty"`
	RegistrationStatus    *string               `json:"registrationStatus,omitempty"`
}

var _ json.Marshaler = AzureIaaSClassicComputeVMContainer{}

func (s AzureIaaSClassicComputeVMContainer) MarshalJSON() ([]byte, error) {
	type wrapper AzureIaaSClassicComputeVMContainer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureIaaSClassicComputeVMContainer: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureIaaSClassicComputeVMContainer: %+v", err)
	}
	decoded["containerType"] = "Microsoft.ClassicCompute/virtualMachines"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureIaaSClassicComputeVMContainer: %+v", err)
	}

	return encoded, nil
}
