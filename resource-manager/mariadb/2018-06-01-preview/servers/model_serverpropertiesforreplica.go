package servers

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ServerPropertiesForCreate = ServerPropertiesForReplica{}

type ServerPropertiesForReplica struct {
	SourceServerId string `json:"sourceServerId"`

	// Fields inherited from ServerPropertiesForCreate
	MinimalTlsVersion *MinimalTlsVersionEnum `json:"minimalTlsVersion,omitempty"`
	SslEnforcement    *SslEnforcementEnum    `json:"sslEnforcement,omitempty"`
	StorageProfile    *StorageProfile        `json:"storageProfile,omitempty"`
	Version           *ServerVersion         `json:"version,omitempty"`
}

var _ json.Marshaler = ServerPropertiesForReplica{}

func (s ServerPropertiesForReplica) MarshalJSON() ([]byte, error) {
	type wrapper ServerPropertiesForReplica
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServerPropertiesForReplica: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServerPropertiesForReplica: %+v", err)
	}
	decoded["createMode"] = "Replica"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServerPropertiesForReplica: %+v", err)
	}

	return encoded, nil
}
