package jitrequests

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JitRequestProperties struct {
	ApplicationResourceId    string                     `json:"applicationResourceId"`
	CreatedBy                *ApplicationClientDetails  `json:"createdBy,omitempty"`
	JitAuthorizationPolicies []JitAuthorizationPolicies `json:"jitAuthorizationPolicies"`
	JitRequestState          *JitRequestState           `json:"jitRequestState,omitempty"`
	JitSchedulingPolicy      JitSchedulingPolicy        `json:"jitSchedulingPolicy"`
	ProvisioningState        *ProvisioningState         `json:"provisioningState,omitempty"`
	PublisherTenantId        *string                    `json:"publisherTenantId,omitempty"`
	UpdatedBy                *ApplicationClientDetails  `json:"updatedBy,omitempty"`
}
