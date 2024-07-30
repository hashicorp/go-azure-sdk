package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerSharedWithContainer struct {
	AccessLevel *PlannerSharedWithContainerAccessLevel `json:"accessLevel,omitempty"`
	ContainerId *string                                `json:"containerId,omitempty"`
	ODataType   *string                                `json:"@odata.type,omitempty"`
	Type        *PlannerSharedWithContainerType        `json:"type,omitempty"`
	Url         *string                                `json:"url,omitempty"`
}
