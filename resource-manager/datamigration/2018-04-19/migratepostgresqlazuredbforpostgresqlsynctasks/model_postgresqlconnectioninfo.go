package migratepostgresqlazuredbforpostgresqlsynctasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ConnectionInfo = PostgreSqlConnectionInfo{}

type PostgreSqlConnectionInfo struct {
	DatabaseName *string `json:"databaseName,omitempty"`
	Port         int64   `json:"port"`
	ServerName   string  `json:"serverName"`

	// Fields inherited from ConnectionInfo
	Password *string `json:"password,omitempty"`
	UserName *string `json:"userName,omitempty"`
}

var _ json.Marshaler = PostgreSqlConnectionInfo{}

func (s PostgreSqlConnectionInfo) MarshalJSON() ([]byte, error) {
	type wrapper PostgreSqlConnectionInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PostgreSqlConnectionInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PostgreSqlConnectionInfo: %+v", err)
	}
	decoded["type"] = "PostgreSqlConnectionInfo"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PostgreSqlConnectionInfo: %+v", err)
	}

	return encoded, nil
}
