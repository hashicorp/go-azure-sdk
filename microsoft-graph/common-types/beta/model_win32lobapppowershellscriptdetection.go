package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptDetection struct {
	EnforceSignatureCheck *bool   `json:"enforceSignatureCheck,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
	RunAs32Bit            *bool   `json:"runAs32Bit,omitempty"`
	ScriptContent         *string `json:"scriptContent,omitempty"`
}
