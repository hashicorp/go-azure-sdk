package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskDefinition struct {
	Category        *IdentityGovernanceTaskDefinitionCategory `json:"category,omitempty"`
	ContinueOnError *bool                                     `json:"continueOnError,omitempty"`
	Description     *string                                   `json:"description,omitempty"`
	DisplayName     *string                                   `json:"displayName,omitempty"`
	Id              *string                                   `json:"id,omitempty"`
	ODataType       *string                                   `json:"@odata.type,omitempty"`
	Parameters      *[]IdentityGovernanceParameter            `json:"parameters,omitempty"`
	Version         *int64                                    `json:"version,omitempty"`
}
