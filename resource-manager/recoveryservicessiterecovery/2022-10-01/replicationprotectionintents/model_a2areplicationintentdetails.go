package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ReplicationProtectionIntentProviderSpecificSettings = A2AReplicationIntentDetails{}

type A2AReplicationIntentDetails struct {
	AgentAutoUpdateStatus               *AgentAutoUpdateStatus                        `json:"agentAutoUpdateStatus,omitempty"`
	AutoProtectionOfDataDisk            *AutoProtectionOfDataDisk                     `json:"autoProtectionOfDataDisk,omitempty"`
	AutomationAccountArmId              *string                                       `json:"automationAccountArmId,omitempty"`
	AutomationAccountAuthenticationType *AutomationAccountAuthenticationType          `json:"automationAccountAuthenticationType,omitempty"`
	DiskEncryptionInfo                  *DiskEncryptionInfo                           `json:"diskEncryptionInfo"`
	FabricObjectId                      *string                                       `json:"fabricObjectId,omitempty"`
	MultiVmGroupId                      *string                                       `json:"multiVmGroupId,omitempty"`
	MultiVmGroupName                    *string                                       `json:"multiVmGroupName,omitempty"`
	PrimaryLocation                     *string                                       `json:"primaryLocation,omitempty"`
	PrimaryStagingStorageAccount        StorageAccountCustomDetails                   `json:"primaryStagingStorageAccount"`
	ProtectionProfile                   ProtectionProfileCustomDetails                `json:"protectionProfile"`
	RecoveryAvailabilitySet             RecoveryAvailabilitySetCustomDetails          `json:"recoveryAvailabilitySet"`
	RecoveryAvailabilityType            string                                        `json:"recoveryAvailabilityType"`
	RecoveryAvailabilityZone            *string                                       `json:"recoveryAvailabilityZone,omitempty"`
	RecoveryBootDiagStorageAccount      StorageAccountCustomDetails                   `json:"recoveryBootDiagStorageAccount"`
	RecoveryLocation                    *string                                       `json:"recoveryLocation,omitempty"`
	RecoveryProximityPlacementGroup     RecoveryProximityPlacementGroupCustomDetails  `json:"recoveryProximityPlacementGroup"`
	RecoveryResourceGroupId             *string                                       `json:"recoveryResourceGroupId,omitempty"`
	RecoverySubscriptionId              *string                                       `json:"recoverySubscriptionId,omitempty"`
	RecoveryVirtualNetwork              RecoveryVirtualNetworkCustomDetails           `json:"recoveryVirtualNetwork"`
	VmDisks                             *[]A2AProtectionIntentDiskInputDetails        `json:"vmDisks,omitempty"`
	VmManagedDisks                      *[]A2AProtectionIntentManagedDiskInputDetails `json:"vmManagedDisks,omitempty"`

	// Fields inherited from ReplicationProtectionIntentProviderSpecificSettings
}

var _ json.Marshaler = A2AReplicationIntentDetails{}

func (s A2AReplicationIntentDetails) MarshalJSON() ([]byte, error) {
	type wrapper A2AReplicationIntentDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling A2AReplicationIntentDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling A2AReplicationIntentDetails: %+v", err)
	}
	decoded["instanceType"] = "A2A"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling A2AReplicationIntentDetails: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &A2AReplicationIntentDetails{}

func (s *A2AReplicationIntentDetails) UnmarshalJSON(bytes []byte) error {
	type alias A2AReplicationIntentDetails
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into A2AReplicationIntentDetails: %+v", err)
	}

	s.AgentAutoUpdateStatus = decoded.AgentAutoUpdateStatus
	s.AutoProtectionOfDataDisk = decoded.AutoProtectionOfDataDisk
	s.AutomationAccountArmId = decoded.AutomationAccountArmId
	s.AutomationAccountAuthenticationType = decoded.AutomationAccountAuthenticationType
	s.DiskEncryptionInfo = decoded.DiskEncryptionInfo
	s.FabricObjectId = decoded.FabricObjectId
	s.MultiVmGroupId = decoded.MultiVmGroupId
	s.MultiVmGroupName = decoded.MultiVmGroupName
	s.PrimaryLocation = decoded.PrimaryLocation
	s.RecoveryAvailabilityType = decoded.RecoveryAvailabilityType
	s.RecoveryAvailabilityZone = decoded.RecoveryAvailabilityZone
	s.RecoveryLocation = decoded.RecoveryLocation
	s.RecoveryResourceGroupId = decoded.RecoveryResourceGroupId
	s.RecoverySubscriptionId = decoded.RecoverySubscriptionId
	s.VmDisks = decoded.VmDisks
	s.VmManagedDisks = decoded.VmManagedDisks

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling A2AReplicationIntentDetails into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["primaryStagingStorageAccount"]; ok {
		impl, err := unmarshalStorageAccountCustomDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'PrimaryStagingStorageAccount' for 'A2AReplicationIntentDetails': %+v", err)
		}
		s.PrimaryStagingStorageAccount = impl
	}

	if v, ok := temp["protectionProfile"]; ok {
		impl, err := unmarshalProtectionProfileCustomDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProtectionProfile' for 'A2AReplicationIntentDetails': %+v", err)
		}
		s.ProtectionProfile = impl
	}

	if v, ok := temp["recoveryAvailabilitySet"]; ok {
		impl, err := unmarshalRecoveryAvailabilitySetCustomDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RecoveryAvailabilitySet' for 'A2AReplicationIntentDetails': %+v", err)
		}
		s.RecoveryAvailabilitySet = impl
	}

	if v, ok := temp["recoveryBootDiagStorageAccount"]; ok {
		impl, err := unmarshalStorageAccountCustomDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RecoveryBootDiagStorageAccount' for 'A2AReplicationIntentDetails': %+v", err)
		}
		s.RecoveryBootDiagStorageAccount = impl
	}

	if v, ok := temp["recoveryProximityPlacementGroup"]; ok {
		impl, err := unmarshalRecoveryProximityPlacementGroupCustomDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RecoveryProximityPlacementGroup' for 'A2AReplicationIntentDetails': %+v", err)
		}
		s.RecoveryProximityPlacementGroup = impl
	}

	if v, ok := temp["recoveryVirtualNetwork"]; ok {
		impl, err := unmarshalRecoveryVirtualNetworkCustomDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RecoveryVirtualNetwork' for 'A2AReplicationIntentDetails': %+v", err)
		}
		s.RecoveryVirtualNetwork = impl
	}
	return nil
}
