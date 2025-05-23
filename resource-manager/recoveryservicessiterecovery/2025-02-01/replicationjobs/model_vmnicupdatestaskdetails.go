package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TaskTypeDetails = VMNicUpdatesTaskDetails{}

type VMNicUpdatesTaskDetails struct {
	Name  *string `json:"name,omitempty"`
	NicId *string `json:"nicId,omitempty"`
	VMId  *string `json:"vmId,omitempty"`

	// Fields inherited from TaskTypeDetails

	InstanceType string `json:"instanceType"`
}

func (s VMNicUpdatesTaskDetails) TaskTypeDetails() BaseTaskTypeDetailsImpl {
	return BaseTaskTypeDetailsImpl{
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = VMNicUpdatesTaskDetails{}

func (s VMNicUpdatesTaskDetails) MarshalJSON() ([]byte, error) {
	type wrapper VMNicUpdatesTaskDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VMNicUpdatesTaskDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VMNicUpdatesTaskDetails: %+v", err)
	}

	decoded["instanceType"] = "VmNicUpdatesTaskDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VMNicUpdatesTaskDetails: %+v", err)
	}

	return encoded, nil
}
