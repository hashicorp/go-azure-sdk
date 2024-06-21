package portalrevision

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PortalRevisionContractProperties struct {
	CreatedDateTime   *string               `json:"createdDateTime,omitempty"`
	Description       *string               `json:"description,omitempty"`
	IsCurrent         *bool                 `json:"isCurrent,omitempty"`
	ProvisioningState *string               `json:"provisioningState,omitempty"`
	Status            *PortalRevisionStatus `json:"status,omitempty"`
	StatusDetails     *string               `json:"statusDetails,omitempty"`
	UpdatedDateTime   *string               `json:"updatedDateTime,omitempty"`
}
