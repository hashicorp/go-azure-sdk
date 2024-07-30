package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetentionLabelSettings struct {
	BehaviorDuringRetentionPeriod *RetentionLabelSettingsBehaviorDuringRetentionPeriod `json:"behaviorDuringRetentionPeriod,omitempty"`
	IsContentUpdateAllowed        *bool                                                `json:"isContentUpdateAllowed,omitempty"`
	IsDeleteAllowed               *bool                                                `json:"isDeleteAllowed,omitempty"`
	IsLabelUpdateAllowed          *bool                                                `json:"isLabelUpdateAllowed,omitempty"`
	IsMetadataUpdateAllowed       *bool                                                `json:"isMetadataUpdateAllowed,omitempty"`
	IsRecordLocked                *bool                                                `json:"isRecordLocked,omitempty"`
	ODataType                     *string                                              `json:"@odata.type,omitempty"`
}
