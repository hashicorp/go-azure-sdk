package assessments

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ResourceDetails = OnPremiseResourceDetails{}

type OnPremiseResourceDetails struct {
	MachineName      string `json:"machineName"`
	SourceComputerId string `json:"sourceComputerId"`
	VMuuid           string `json:"vmuuid"`
	WorkspaceId      string `json:"workspaceId"`

	// Fields inherited from ResourceDetails

	Source Source `json:"source"`
}

func (s OnPremiseResourceDetails) ResourceDetails() BaseResourceDetailsImpl {
	return BaseResourceDetailsImpl{
		Source: s.Source,
	}
}

var _ json.Marshaler = OnPremiseResourceDetails{}

func (s OnPremiseResourceDetails) MarshalJSON() ([]byte, error) {
	type wrapper OnPremiseResourceDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnPremiseResourceDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnPremiseResourceDetails: %+v", err)
	}

	decoded["source"] = "OnPremise"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnPremiseResourceDetails: %+v", err)
	}

	return encoded, nil
}
