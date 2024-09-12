package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TeamsAppInstallationScopeInfo = GroupChatTeamsAppInstallationScopeInfo{}

type GroupChatTeamsAppInstallationScopeInfo struct {
	ChatId nullable.Type[string] `json:"chatId,omitempty"`

	// Fields inherited from TeamsAppInstallationScopeInfo

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Scope *TeamsAppInstallationScopes `json:"scope,omitempty"`
}

func (s GroupChatTeamsAppInstallationScopeInfo) TeamsAppInstallationScopeInfo() BaseTeamsAppInstallationScopeInfoImpl {
	return BaseTeamsAppInstallationScopeInfoImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		Scope:     s.Scope,
	}
}

var _ json.Marshaler = GroupChatTeamsAppInstallationScopeInfo{}

func (s GroupChatTeamsAppInstallationScopeInfo) MarshalJSON() ([]byte, error) {
	type wrapper GroupChatTeamsAppInstallationScopeInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupChatTeamsAppInstallationScopeInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupChatTeamsAppInstallationScopeInfo: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.groupChatTeamsAppInstallationScopeInfo"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupChatTeamsAppInstallationScopeInfo: %+v", err)
	}

	return encoded, nil
}
