package python3package

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModuleProperties struct {
	ActivityCount     *int64                   `json:"activityCount,omitempty"`
	ContentLink       *ContentLink             `json:"contentLink,omitempty"`
	CreationTime      *string                  `json:"creationTime,omitempty"`
	Description       *string                  `json:"description,omitempty"`
	Error             *ModuleErrorInfo         `json:"error,omitempty"`
	IsComposite       *bool                    `json:"isComposite,omitempty"`
	IsGlobal          *bool                    `json:"isGlobal,omitempty"`
	LastModifiedTime  *string                  `json:"lastModifiedTime,omitempty"`
	ProvisioningState *ModuleProvisioningState `json:"provisioningState,omitempty"`
	SizeInBytes       *int64                   `json:"sizeInBytes,omitempty"`
	Version           *string                  `json:"version,omitempty"`
}
