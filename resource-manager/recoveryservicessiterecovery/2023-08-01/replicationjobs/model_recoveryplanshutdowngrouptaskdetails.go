package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GroupTaskDetails = RecoveryPlanShutdownGroupTaskDetails{}

type RecoveryPlanShutdownGroupTaskDetails struct {

	// Fields inherited from RecoveryPlanGroupTaskDetails
	GroupId     *string `json:"groupId,omitempty"`
	Name        *string `json:"name,omitempty"`
	RpGroupType *string `json:"rpGroupType,omitempty"`
}

var _ json.Marshaler = RecoveryPlanShutdownGroupTaskDetails{}

func (s RecoveryPlanShutdownGroupTaskDetails) MarshalJSON() ([]byte, error) {
	type wrapper RecoveryPlanShutdownGroupTaskDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RecoveryPlanShutdownGroupTaskDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RecoveryPlanShutdownGroupTaskDetails: %+v", err)
	}
	decoded["instanceType"] = "RecoveryPlanShutdownGroupTaskDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RecoveryPlanShutdownGroupTaskDetails: %+v", err)
	}

	return encoded, nil
}
