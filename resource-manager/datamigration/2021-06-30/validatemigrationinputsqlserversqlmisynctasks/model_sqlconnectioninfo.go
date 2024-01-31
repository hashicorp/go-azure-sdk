package validatemigrationinputsqlserversqlmisynctasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ConnectionInfo = SqlConnectionInfo{}

type SqlConnectionInfo struct {
	AdditionalSettings     *string             `json:"additionalSettings,omitempty"`
	Authentication         *AuthenticationType `json:"authentication,omitempty"`
	DataSource             string              `json:"dataSource"`
	EncryptConnection      *bool               `json:"encryptConnection,omitempty"`
	Platform               *SqlSourcePlatform  `json:"platform,omitempty"`
	TrustServerCertificate *bool               `json:"trustServerCertificate,omitempty"`

	// Fields inherited from ConnectionInfo
	Password *string `json:"password,omitempty"`
	UserName *string `json:"userName,omitempty"`
}

var _ json.Marshaler = SqlConnectionInfo{}

func (s SqlConnectionInfo) MarshalJSON() ([]byte, error) {
	type wrapper SqlConnectionInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SqlConnectionInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SqlConnectionInfo: %+v", err)
	}
	decoded["type"] = "SqlConnectionInfo"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SqlConnectionInfo: %+v", err)
	}

	return encoded, nil
}
