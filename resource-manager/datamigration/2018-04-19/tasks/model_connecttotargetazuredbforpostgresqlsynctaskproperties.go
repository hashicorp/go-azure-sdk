package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = ConnectToTargetAzureDbForPostgreSqlSyncTaskProperties{}

type ConnectToTargetAzureDbForPostgreSqlSyncTaskProperties struct {
	Input  *ConnectToTargetAzureDbForPostgreSqlSyncTaskInput    `json:"input,omitempty"`
	Output *[]ConnectToTargetAzureDbForPostgreSqlSyncTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	Commands *[]CommandProperties `json:"commands,omitempty"`
	Errors   *[]ODataError        `json:"errors,omitempty"`
	State    *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = ConnectToTargetAzureDbForPostgreSqlSyncTaskProperties{}

func (s ConnectToTargetAzureDbForPostgreSqlSyncTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper ConnectToTargetAzureDbForPostgreSqlSyncTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConnectToTargetAzureDbForPostgreSqlSyncTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConnectToTargetAzureDbForPostgreSqlSyncTaskProperties: %+v", err)
	}
	decoded["taskType"] = "ConnectToTarget.AzureDbForPostgreSql.Sync"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConnectToTargetAzureDbForPostgreSqlSyncTaskProperties: %+v", err)
	}

	return encoded, nil
}
