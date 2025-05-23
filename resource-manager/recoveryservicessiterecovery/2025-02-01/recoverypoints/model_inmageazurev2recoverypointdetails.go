package recoverypoints

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProviderSpecificRecoveryPointDetails = InMageAzureV2RecoveryPointDetails{}

type InMageAzureV2RecoveryPointDetails struct {
	IsMultiVMSyncPoint *string `json:"isMultiVmSyncPoint,omitempty"`

	// Fields inherited from ProviderSpecificRecoveryPointDetails

	InstanceType string `json:"instanceType"`
}

func (s InMageAzureV2RecoveryPointDetails) ProviderSpecificRecoveryPointDetails() BaseProviderSpecificRecoveryPointDetailsImpl {
	return BaseProviderSpecificRecoveryPointDetailsImpl{
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = InMageAzureV2RecoveryPointDetails{}

func (s InMageAzureV2RecoveryPointDetails) MarshalJSON() ([]byte, error) {
	type wrapper InMageAzureV2RecoveryPointDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InMageAzureV2RecoveryPointDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InMageAzureV2RecoveryPointDetails: %+v", err)
	}

	decoded["instanceType"] = "InMageAzureV2"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InMageAzureV2RecoveryPointDetails: %+v", err)
	}

	return encoded, nil
}
