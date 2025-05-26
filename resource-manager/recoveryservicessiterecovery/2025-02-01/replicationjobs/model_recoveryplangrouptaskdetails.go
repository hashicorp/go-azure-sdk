package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GroupTaskDetails = RecoveryPlanGroupTaskDetails{}

type RecoveryPlanGroupTaskDetails struct {
	GroupId     *string `json:"groupId,omitempty"`
	Name        *string `json:"name,omitempty"`
	RpGroupType *string `json:"rpGroupType,omitempty"`

	// Fields inherited from GroupTaskDetails

	ChildTasks   *[]ASRTask `json:"childTasks,omitempty"`
	InstanceType string     `json:"instanceType"`
}

func (s RecoveryPlanGroupTaskDetails) GroupTaskDetails() BaseGroupTaskDetailsImpl {
	return BaseGroupTaskDetailsImpl{
		ChildTasks:   s.ChildTasks,
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = RecoveryPlanGroupTaskDetails{}

func (s RecoveryPlanGroupTaskDetails) MarshalJSON() ([]byte, error) {
	type wrapper RecoveryPlanGroupTaskDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RecoveryPlanGroupTaskDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RecoveryPlanGroupTaskDetails: %+v", err)
	}

	decoded["instanceType"] = "RecoveryPlanGroupTaskDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RecoveryPlanGroupTaskDetails: %+v", err)
	}

	return encoded, nil
}
