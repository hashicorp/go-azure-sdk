package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type A2AProtectionIntentManagedDiskInputDetails struct {
	DiskEncryptionInfo                      *DiskEncryptionInfo                 `json:"diskEncryptionInfo,omitempty"`
	DiskId                                  string                              `json:"diskId"`
	PrimaryStagingStorageAccountCustomInput *StorageAccountCustomDetails        `json:"primaryStagingStorageAccountCustomInput,omitempty"`
	RecoveryDiskEncryptionSetId             *string                             `json:"recoveryDiskEncryptionSetId,omitempty"`
	RecoveryReplicaDiskAccountType          *string                             `json:"recoveryReplicaDiskAccountType,omitempty"`
	RecoveryResourceGroupCustomInput        *RecoveryResourceGroupCustomDetails `json:"recoveryResourceGroupCustomInput,omitempty"`
	RecoveryTargetDiskAccountType           *string                             `json:"recoveryTargetDiskAccountType,omitempty"`
}

var _ json.Unmarshaler = &A2AProtectionIntentManagedDiskInputDetails{}

func (s *A2AProtectionIntentManagedDiskInputDetails) UnmarshalJSON(bytes []byte) error {
	type alias A2AProtectionIntentManagedDiskInputDetails
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into A2AProtectionIntentManagedDiskInputDetails: %+v", err)
	}

	s.DiskEncryptionInfo = decoded.DiskEncryptionInfo
	s.DiskId = decoded.DiskId
	s.RecoveryDiskEncryptionSetId = decoded.RecoveryDiskEncryptionSetId
	s.RecoveryReplicaDiskAccountType = decoded.RecoveryReplicaDiskAccountType
	s.RecoveryTargetDiskAccountType = decoded.RecoveryTargetDiskAccountType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling A2AProtectionIntentManagedDiskInputDetails into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["primaryStagingStorageAccountCustomInput"]; ok {
		impl, err := unmarshalStorageAccountCustomDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'PrimaryStagingStorageAccountCustomInput' for 'A2AProtectionIntentManagedDiskInputDetails': %+v", err)
		}
		s.PrimaryStagingStorageAccountCustomInput = &impl
	}

	if v, ok := temp["recoveryResourceGroupCustomInput"]; ok {
		impl, err := unmarshalRecoveryResourceGroupCustomDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RecoveryResourceGroupCustomInput' for 'A2AProtectionIntentManagedDiskInputDetails': %+v", err)
		}
		s.RecoveryResourceGroupCustomInput = &impl
	}
	return nil
}
