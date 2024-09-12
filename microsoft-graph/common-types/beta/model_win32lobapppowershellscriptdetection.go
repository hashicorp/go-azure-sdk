package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Win32LobAppDetection = Win32LobAppPowerShellScriptDetection{}

type Win32LobAppPowerShellScriptDetection struct {
	// A value indicating whether signature check is enforced
	EnforceSignatureCheck *bool `json:"enforceSignatureCheck,omitempty"`

	// A value indicating whether this script should run as 32-bit
	RunAs32Bit *bool `json:"runAs32Bit,omitempty"`

	// The base64 encoded script content to detect Win32 Line of Business (LoB) app
	ScriptContent nullable.Type[string] `json:"scriptContent,omitempty"`

	// Fields inherited from Win32LobAppDetection

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s Win32LobAppPowerShellScriptDetection) Win32LobAppDetection() BaseWin32LobAppDetectionImpl {
	return BaseWin32LobAppDetectionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Win32LobAppPowerShellScriptDetection{}

func (s Win32LobAppPowerShellScriptDetection) MarshalJSON() ([]byte, error) {
	type wrapper Win32LobAppPowerShellScriptDetection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Win32LobAppPowerShellScriptDetection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32LobAppPowerShellScriptDetection: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.win32LobAppPowerShellScriptDetection"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Win32LobAppPowerShellScriptDetection: %+v", err)
	}

	return encoded, nil
}
