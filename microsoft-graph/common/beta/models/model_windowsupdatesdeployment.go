package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDeployment struct {
	Audience             *WindowsUpdatesDeploymentAudience `json:"audience,omitempty"`
	Content              *WindowsUpdatesDeployableContent  `json:"content,omitempty"`
	CreatedDateTime      *string                           `json:"createdDateTime,omitempty"`
	Id                   *string                           `json:"id,omitempty"`
	LastModifiedDateTime *string                           `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                           `json:"@odata.type,omitempty"`
	Settings             *WindowsUpdatesDeploymentSettings `json:"settings,omitempty"`
	State                *WindowsUpdatesDeploymentState    `json:"state,omitempty"`
}
