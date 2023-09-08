package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDeploymentAudience struct {
	ApplicableContent *[]WindowsUpdatesApplicableContent `json:"applicableContent,omitempty"`
	Exclusions        *[]WindowsUpdatesUpdatableAsset    `json:"exclusions,omitempty"`
	Id                *string                            `json:"id,omitempty"`
	Members           *[]WindowsUpdatesUpdatableAsset    `json:"members,omitempty"`
	ODataType         *string                            `json:"@odata.type,omitempty"`
}
