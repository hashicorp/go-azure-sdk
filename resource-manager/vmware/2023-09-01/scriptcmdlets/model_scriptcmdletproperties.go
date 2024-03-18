package scriptcmdlets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptCmdletProperties struct {
	Audience          *ScriptCmdletAudience          `json:"audience,omitempty"`
	Description       *string                        `json:"description,omitempty"`
	Parameters        *[]ScriptParameter             `json:"parameters,omitempty"`
	ProvisioningState *ScriptCmdletProvisioningState `json:"provisioningState,omitempty"`
	Timeout           *string                        `json:"timeout,omitempty"`
}
