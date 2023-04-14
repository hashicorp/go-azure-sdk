package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CreateProtectionIntentProviderSpecificDetails = A2ACreateProtectionIntentInput{}

type A2ACreateProtectionIntentInput struct {
	AgentAutoUpdateStatus                      *AgentAutoUpdateStatus                        `json:"agentAutoUpdateStatus,omitempty"`
	AutoProtectionOfDataDisk                   *AutoProtectionOfDataDisk                     `json:"autoProtectionOfDataDisk,omitempty"`
	AutomationAccountArmId                     *string                                       `json:"automationAccountArmId,omitempty"`
	AutomationAccountAuthenticationType        *AutomationAccountAuthenticationType          `json:"automationAccountAuthenticationType,omitempty"`
	DiskEncryptionInfo                         *DiskEncryptionInfo                           `json:"diskEncryptionInfo,omitempty"`
	FabricObjectId                             string                                        `json:"fabricObjectId"`
	MultiVMGroupId                             *string                                       `json:"multiVmGroupId,omitempty"`
	MultiVMGroupName                           *string                                       `json:"multiVmGroupName,omitempty"`
	PrimaryLocation                            string                                        `json:"primaryLocation"`
	PrimaryStagingStorageAccountCustomInput    *StorageAccountCustomDetails                  `json:"primaryStagingStorageAccountCustomInput,omitempty"`
	ProtectionProfileCustomInput               *ProtectionProfileCustomDetails               `json:"protectionProfileCustomInput,omitempty"`
	RecoveryAvailabilitySetCustomInput         *RecoveryAvailabilitySetCustomDetails         `json:"recoveryAvailabilitySetCustomInput,omitempty"`
	RecoveryAvailabilityType                   A2ARecoveryAvailabilityType                   `json:"recoveryAvailabilityType"`
	RecoveryAvailabilityZone                   *string                                       `json:"recoveryAvailabilityZone,omitempty"`
	RecoveryBootDiagStorageAccount             *StorageAccountCustomDetails                  `json:"recoveryBootDiagStorageAccount,omitempty"`
	RecoveryLocation                           string                                        `json:"recoveryLocation"`
	RecoveryProximityPlacementGroupCustomInput *RecoveryProximityPlacementGroupCustomDetails `json:"recoveryProximityPlacementGroupCustomInput,omitempty"`
	RecoveryResourceGroupId                    string                                        `json:"recoveryResourceGroupId"`
	RecoverySubscriptionId                     string                                        `json:"recoverySubscriptionId"`
	RecoveryVirtualNetworkCustomInput          *RecoveryVirtualNetworkCustomDetails          `json:"recoveryVirtualNetworkCustomInput,omitempty"`
	VMDisks                                    *[]A2AProtectionIntentDiskInputDetails        `json:"vmDisks,omitempty"`
	VMManagedDisks                             *[]A2AProtectionIntentManagedDiskInputDetails `json:"vmManagedDisks,omitempty"`

	// Fields inherited from CreateProtectionIntentProviderSpecificDetails
}

var _ json.Marshaler = A2ACreateProtectionIntentInput{}

func (s A2ACreateProtectionIntentInput) MarshalJSON() ([]byte, error) {
	type wrapper A2ACreateProtectionIntentInput
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling A2ACreateProtectionIntentInput: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling A2ACreateProtectionIntentInput: %+v", err)
	}
	decoded["instanceType"] = "A2A"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling A2ACreateProtectionIntentInput: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &A2ACreateProtectionIntentInput{}

func (s *A2ACreateProtectionIntentInput) UnmarshalJSON(bytes []byte) error {
	type alias A2ACreateProtectionIntentInput
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into A2ACreateProtectionIntentInput: %+v", err)
	}

	s.AgentAutoUpdateStatus = decoded.AgentAutoUpdateStatus
	s.AutoProtectionOfDataDisk = decoded.AutoProtectionOfDataDisk
	s.AutomationAccountArmId = decoded.AutomationAccountArmId
	s.AutomationAccountAuthenticationType = decoded.AutomationAccountAuthenticationType
	s.DiskEncryptionInfo = decoded.DiskEncryptionInfo
	s.FabricObjectId = decoded.FabricObjectId
	s.MultiVMGroupId = decoded.MultiVMGroupId
	s.MultiVMGroupName = decoded.MultiVMGroupName
	s.PrimaryLocation = decoded.PrimaryLocation
	s.RecoveryAvailabilityType = decoded.RecoveryAvailabilityType
	s.RecoveryAvailabilityZone = decoded.RecoveryAvailabilityZone
	s.RecoveryLocation = decoded.RecoveryLocation
	s.RecoveryResourceGroupId = decoded.RecoveryResourceGroupId
	s.RecoverySubscriptionId = decoded.RecoverySubscriptionId
	s.VMDisks = decoded.VMDisks
	s.VMManagedDisks = decoded.VMManagedDisks

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling A2ACreateProtectionIntentInput into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["primaryStagingStorageAccountCustomInput"]; ok {
		impl, err := unmarshalStorageAccountCustomDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'PrimaryStagingStorageAccountCustomInput' for 'A2ACreateProtectionIntentInput': %+v", err)
		}
		s.PrimaryStagingStorageAccountCustomInput = impl
	}

	if v, ok := temp["protectionProfileCustomInput"]; ok {
		impl, err := unmarshalProtectionProfileCustomDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProtectionProfileCustomInput' for 'A2ACreateProtectionIntentInput': %+v", err)
		}
		s.ProtectionProfileCustomInput = impl
	}

	if v, ok := temp["recoveryAvailabilitySetCustomInput"]; ok {
		impl, err := unmarshalRecoveryAvailabilitySetCustomDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RecoveryAvailabilitySetCustomInput' for 'A2ACreateProtectionIntentInput': %+v", err)
		}
		s.RecoveryAvailabilitySetCustomInput = impl
	}

	if v, ok := temp["recoveryBootDiagStorageAccount"]; ok {
		impl, err := unmarshalStorageAccountCustomDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RecoveryBootDiagStorageAccount' for 'A2ACreateProtectionIntentInput': %+v", err)
		}
		s.RecoveryBootDiagStorageAccount = impl
	}

	if v, ok := temp["recoveryProximityPlacementGroupCustomInput"]; ok {
		impl, err := unmarshalRecoveryProximityPlacementGroupCustomDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RecoveryProximityPlacementGroupCustomInput' for 'A2ACreateProtectionIntentInput': %+v", err)
		}
		s.RecoveryProximityPlacementGroupCustomInput = impl
	}

	if v, ok := temp["recoveryVirtualNetworkCustomInput"]; ok {
		impl, err := unmarshalRecoveryVirtualNetworkCustomDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RecoveryVirtualNetworkCustomInput' for 'A2ACreateProtectionIntentInput': %+v", err)
		}
		s.RecoveryVirtualNetworkCustomInput = impl
	}
	return nil
}
