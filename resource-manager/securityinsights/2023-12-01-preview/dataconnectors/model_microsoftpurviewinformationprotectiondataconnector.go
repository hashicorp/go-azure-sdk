package dataconnectors

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataConnector = MicrosoftPurviewInformationProtectionDataConnector{}

type MicrosoftPurviewInformationProtectionDataConnector struct {
	Properties *MicrosoftPurviewInformationProtectionDataConnectorProperties `json:"properties,omitempty"`

	// Fields inherited from DataConnector

	Etag       *string                `json:"etag,omitempty"`
	Id         *string                `json:"id,omitempty"`
	Kind       DataConnectorKind      `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s MicrosoftPurviewInformationProtectionDataConnector) DataConnector() BaseDataConnectorImpl {
	return BaseDataConnectorImpl{
		Etag:       s.Etag,
		Id:         s.Id,
		Kind:       s.Kind,
		Name:       s.Name,
		SystemData: s.SystemData,
		Type:       s.Type,
	}
}

var _ json.Marshaler = MicrosoftPurviewInformationProtectionDataConnector{}

func (s MicrosoftPurviewInformationProtectionDataConnector) MarshalJSON() ([]byte, error) {
	type wrapper MicrosoftPurviewInformationProtectionDataConnector
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MicrosoftPurviewInformationProtectionDataConnector: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MicrosoftPurviewInformationProtectionDataConnector: %+v", err)
	}

	decoded["kind"] = "MicrosoftPurviewInformationProtection"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MicrosoftPurviewInformationProtectionDataConnector: %+v", err)
	}

	return encoded, nil
}
