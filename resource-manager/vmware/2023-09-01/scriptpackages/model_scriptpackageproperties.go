package scriptpackages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptPackageProperties struct {
	Company           *string                         `json:"company,omitempty"`
	Description       *string                         `json:"description,omitempty"`
	ProvisioningState *ScriptPackageProvisioningState `json:"provisioningState,omitempty"`
	Uri               *string                         `json:"uri,omitempty"`
	Version           *string                         `json:"version,omitempty"`
}
