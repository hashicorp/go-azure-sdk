package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type A2AProtectionIntentDiskInputDetails struct {
	DiskUri                                 string                      `json:"diskUri"`
	PrimaryStagingStorageAccountCustomInput StorageAccountCustomDetails `json:"primaryStagingStorageAccountCustomInput"`
	RecoveryAzureStorageAccountCustomInput  StorageAccountCustomDetails `json:"recoveryAzureStorageAccountCustomInput"`
}

var _ json.Unmarshaler = &A2AProtectionIntentDiskInputDetails{}

func (s *A2AProtectionIntentDiskInputDetails) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DiskUri string `json:"diskUri"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DiskUri = decoded.DiskUri

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling A2AProtectionIntentDiskInputDetails into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["primaryStagingStorageAccountCustomInput"]; ok {
		impl, err := UnmarshalStorageAccountCustomDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'PrimaryStagingStorageAccountCustomInput' for 'A2AProtectionIntentDiskInputDetails': %+v", err)
		}
		s.PrimaryStagingStorageAccountCustomInput = impl
	}

	if v, ok := temp["recoveryAzureStorageAccountCustomInput"]; ok {
		impl, err := UnmarshalStorageAccountCustomDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RecoveryAzureStorageAccountCustomInput' for 'A2AProtectionIntentDiskInputDetails': %+v", err)
		}
		s.RecoveryAzureStorageAccountCustomInput = impl
	}

	return nil
}
