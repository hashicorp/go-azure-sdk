package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudOffering = CspmMonitorGcpOffering{}

type CspmMonitorGcpOffering struct {
	NativeCloudConnection *CspmMonitorGcpOfferingNativeCloudConnection `json:"nativeCloudConnection,omitempty"`

	// Fields inherited from CloudOffering
	Description *string `json:"description,omitempty"`
}

var _ json.Marshaler = CspmMonitorGcpOffering{}

func (s CspmMonitorGcpOffering) MarshalJSON() ([]byte, error) {
	type wrapper CspmMonitorGcpOffering
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CspmMonitorGcpOffering: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CspmMonitorGcpOffering: %+v", err)
	}
	decoded["offeringType"] = "CspmMonitorGcp"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CspmMonitorGcpOffering: %+v", err)
	}

	return encoded, nil
}
