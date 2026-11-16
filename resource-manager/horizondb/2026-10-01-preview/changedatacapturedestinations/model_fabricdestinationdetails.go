package changedatacapturedestinations

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ChangeDataCaptureDestinationDetails = FabricDestinationDetails{}

type FabricDestinationDetails struct {
	WorkspaceId string `json:"workspaceId"`

	// Fields inherited from ChangeDataCaptureDestinationDetails

	DestinationType ChangeDataCaptureDestinationType `json:"destinationType"`
	Protocol        ChangeDataCaptureProtocol        `json:"protocol"`
}

func (s FabricDestinationDetails) ChangeDataCaptureDestinationDetails() BaseChangeDataCaptureDestinationDetailsImpl {
	return BaseChangeDataCaptureDestinationDetailsImpl{
		DestinationType: s.DestinationType,
		Protocol:        s.Protocol,
	}
}

var _ json.Marshaler = FabricDestinationDetails{}

func (s FabricDestinationDetails) MarshalJSON() ([]byte, error) {
	type wrapper FabricDestinationDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FabricDestinationDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FabricDestinationDetails: %+v", err)
	}

	decoded["destinationType"] = "Fabric"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FabricDestinationDetails: %+v", err)
	}

	return encoded, nil
}
