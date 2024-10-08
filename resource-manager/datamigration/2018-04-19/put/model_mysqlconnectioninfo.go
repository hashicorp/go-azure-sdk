package put

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ConnectionInfo = MySqlConnectionInfo{}

type MySqlConnectionInfo struct {
	Port       int64  `json:"port"`
	ServerName string `json:"serverName"`

	// Fields inherited from ConnectionInfo

	Password *string `json:"password,omitempty"`
	Type     string  `json:"type"`
	UserName *string `json:"userName,omitempty"`
}

func (s MySqlConnectionInfo) ConnectionInfo() BaseConnectionInfoImpl {
	return BaseConnectionInfoImpl{
		Password: s.Password,
		Type:     s.Type,
		UserName: s.UserName,
	}
}

var _ json.Marshaler = MySqlConnectionInfo{}

func (s MySqlConnectionInfo) MarshalJSON() ([]byte, error) {
	type wrapper MySqlConnectionInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MySqlConnectionInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MySqlConnectionInfo: %+v", err)
	}

	decoded["type"] = "MySqlConnectionInfo"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MySqlConnectionInfo: %+v", err)
	}

	return encoded, nil
}
