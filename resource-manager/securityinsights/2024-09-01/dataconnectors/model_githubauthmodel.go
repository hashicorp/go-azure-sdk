package dataconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CcpAuthConfig = GitHubAuthModel{}

type GitHubAuthModel struct {
	InstallationId *string `json:"installationId,omitempty"`

	// Fields inherited from CcpAuthConfig

	Type CcpAuthType `json:"type"`
}

func (s GitHubAuthModel) CcpAuthConfig() BaseCcpAuthConfigImpl {
	return BaseCcpAuthConfigImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = GitHubAuthModel{}

func (s GitHubAuthModel) MarshalJSON() ([]byte, error) {
	type wrapper GitHubAuthModel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GitHubAuthModel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GitHubAuthModel: %+v", err)
	}

	decoded["type"] = "GitHub"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GitHubAuthModel: %+v", err)
	}

	return encoded, nil
}
