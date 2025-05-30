package scripts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptProperties struct {
	ContinueOnErrors           *bool                       `json:"continueOnErrors,omitempty"`
	ForceUpdateTag             *string                     `json:"forceUpdateTag,omitempty"`
	PrincipalPermissionsAction *PrincipalPermissionsAction `json:"principalPermissionsAction,omitempty"`
	ProvisioningState          *ProvisioningState          `json:"provisioningState,omitempty"`
	ScriptContent              *string                     `json:"scriptContent,omitempty"`
	ScriptLevel                *ScriptLevel                `json:"scriptLevel,omitempty"`
	ScriptURL                  *string                     `json:"scriptUrl,omitempty"`
	ScriptURLSasToken          *string                     `json:"scriptUrlSasToken,omitempty"`
}
