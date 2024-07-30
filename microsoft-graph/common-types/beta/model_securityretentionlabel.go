package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRetentionLabel struct {
	ActionAfterRetentionPeriod    *SecurityRetentionLabelActionAfterRetentionPeriod    `json:"actionAfterRetentionPeriod,omitempty"`
	BehaviorDuringRetentionPeriod *SecurityRetentionLabelBehaviorDuringRetentionPeriod `json:"behaviorDuringRetentionPeriod,omitempty"`
	CreatedBy                     *IdentitySet                                         `json:"createdBy,omitempty"`
	CreatedDateTime               *string                                              `json:"createdDateTime,omitempty"`
	DefaultRecordBehavior         *SecurityRetentionLabelDefaultRecordBehavior         `json:"defaultRecordBehavior,omitempty"`
	DescriptionForAdmins          *string                                              `json:"descriptionForAdmins,omitempty"`
	DescriptionForUsers           *string                                              `json:"descriptionForUsers,omitempty"`
	Descriptors                   *SecurityFilePlanDescriptor                          `json:"descriptors,omitempty"`
	DisplayName                   *string                                              `json:"displayName,omitempty"`
	DispositionReviewStages       *[]SecurityDispositionReviewStage                    `json:"dispositionReviewStages,omitempty"`
	Id                            *string                                              `json:"id,omitempty"`
	IsInUse                       *bool                                                `json:"isInUse,omitempty"`
	LabelToBeApplied              *string                                              `json:"labelToBeApplied,omitempty"`
	LastModifiedBy                *IdentitySet                                         `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime          *string                                              `json:"lastModifiedDateTime,omitempty"`
	ODataType                     *string                                              `json:"@odata.type,omitempty"`
	RetentionDuration             *SecurityRetentionDuration                           `json:"retentionDuration,omitempty"`
	RetentionEventType            *SecurityRetentionEventType                          `json:"retentionEventType,omitempty"`
	RetentionTrigger              *SecurityRetentionLabelRetentionTrigger              `json:"retentionTrigger,omitempty"`
}
