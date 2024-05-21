package replicationprotectionclusters

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ApplyClusterRecoveryPointProviderSpecificInput = A2AApplyClusterRecoveryPointInput{}

type A2AApplyClusterRecoveryPointInput struct {

	// Fields inherited from ApplyClusterRecoveryPointProviderSpecificInput
}

var _ json.Marshaler = A2AApplyClusterRecoveryPointInput{}

func (s A2AApplyClusterRecoveryPointInput) MarshalJSON() ([]byte, error) {
	type wrapper A2AApplyClusterRecoveryPointInput
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling A2AApplyClusterRecoveryPointInput: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling A2AApplyClusterRecoveryPointInput: %+v", err)
	}
	decoded["instanceType"] = "A2A"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling A2AApplyClusterRecoveryPointInput: %+v", err)
	}

	return encoded, nil
}
