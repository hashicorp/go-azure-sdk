package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EnvironmentData = GitlabScopeEnvironmentData{}

type GitlabScopeEnvironmentData struct {

	// Fields inherited from EnvironmentData
}

var _ json.Marshaler = GitlabScopeEnvironmentData{}

func (s GitlabScopeEnvironmentData) MarshalJSON() ([]byte, error) {
	type wrapper GitlabScopeEnvironmentData
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GitlabScopeEnvironmentData: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GitlabScopeEnvironmentData: %+v", err)
	}
	decoded["environmentType"] = "GitlabScope"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GitlabScopeEnvironmentData: %+v", err)
	}

	return encoded, nil
}
