package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EnvironmentData = GithubScopeEnvironmentData{}

type GithubScopeEnvironmentData struct {

	// Fields inherited from EnvironmentData

	EnvironmentType EnvironmentType `json:"environmentType"`
}

func (s GithubScopeEnvironmentData) EnvironmentData() BaseEnvironmentDataImpl {
	return BaseEnvironmentDataImpl{
		EnvironmentType: s.EnvironmentType,
	}
}

var _ json.Marshaler = GithubScopeEnvironmentData{}

func (s GithubScopeEnvironmentData) MarshalJSON() ([]byte, error) {
	type wrapper GithubScopeEnvironmentData
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GithubScopeEnvironmentData: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GithubScopeEnvironmentData: %+v", err)
	}

	decoded["environmentType"] = "GithubScope"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GithubScopeEnvironmentData: %+v", err)
	}

	return encoded, nil
}
