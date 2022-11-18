package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TaskTypeDetails = VmNicUpdatesTaskDetails{}

type VmNicUpdatesTaskDetails struct {
	Name  *string `json:"name,omitempty"`
	NicId *string `json:"nicId,omitempty"`
	VmId  *string `json:"vmId,omitempty"`

	// Fields inherited from TaskTypeDetails
}

var _ json.Marshaler = VmNicUpdatesTaskDetails{}

func (s VmNicUpdatesTaskDetails) MarshalJSON() ([]byte, error) {
	type wrapper VmNicUpdatesTaskDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VmNicUpdatesTaskDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VmNicUpdatesTaskDetails: %+v", err)
	}
	decoded["instanceType"] = "VmNicUpdatesTaskDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VmNicUpdatesTaskDetails: %+v", err)
	}

	return encoded, nil
}
