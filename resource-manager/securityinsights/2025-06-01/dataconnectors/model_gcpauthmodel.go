package dataconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CcpAuthConfig = GCPAuthModel{}

type GCPAuthModel struct {
	ProjectNumber              string `json:"projectNumber"`
	ServiceAccountEmail        string `json:"serviceAccountEmail"`
	WorkloadIdentityProviderId string `json:"workloadIdentityProviderId"`

	// Fields inherited from CcpAuthConfig

	Type CcpAuthType `json:"type"`
}

func (s GCPAuthModel) CcpAuthConfig() BaseCcpAuthConfigImpl {
	return BaseCcpAuthConfigImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = GCPAuthModel{}

func (s GCPAuthModel) MarshalJSON() ([]byte, error) {
	type wrapper GCPAuthModel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GCPAuthModel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GCPAuthModel: %+v", err)
	}

	decoded["type"] = "GCP"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GCPAuthModel: %+v", err)
	}

	return encoded, nil
}
