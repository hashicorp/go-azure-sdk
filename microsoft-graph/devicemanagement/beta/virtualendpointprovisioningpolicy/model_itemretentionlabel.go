package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemRetentionLabel struct {
	Id                       *string                 `json:"id,omitempty"`
	IsLabelAppliedExplicitly *bool                   `json:"isLabelAppliedExplicitly,omitempty"`
	LabelAppliedBy           *IdentitySet            `json:"labelAppliedBy,omitempty"`
	LabelAppliedDateTime     *string                 `json:"labelAppliedDateTime,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	ODataType                *string                 `json:"@odata.type,omitempty"`
	RetentionSettings        *RetentionLabelSettings `json:"retentionSettings,omitempty"`
}
