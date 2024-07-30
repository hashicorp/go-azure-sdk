package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewScheduleDefinition struct {
	AdditionalNotificationRecipients *[]AccessReviewNotificationRecipientItem `json:"additionalNotificationRecipients,omitempty"`
	BackupReviewers                  *[]AccessReviewReviewerScope             `json:"backupReviewers,omitempty"`
	CreatedBy                        *UserIdentity                            `json:"createdBy,omitempty"`
	CreatedDateTime                  *string                                  `json:"createdDateTime,omitempty"`
	DescriptionForAdmins             *string                                  `json:"descriptionForAdmins,omitempty"`
	DescriptionForReviewers          *string                                  `json:"descriptionForReviewers,omitempty"`
	DisplayName                      *string                                  `json:"displayName,omitempty"`
	FallbackReviewers                *[]AccessReviewReviewerScope             `json:"fallbackReviewers,omitempty"`
	Id                               *string                                  `json:"id,omitempty"`
	InstanceEnumerationScope         *AccessReviewScope                       `json:"instanceEnumerationScope,omitempty"`
	Instances                        *[]AccessReviewInstance                  `json:"instances,omitempty"`
	LastModifiedDateTime             *string                                  `json:"lastModifiedDateTime,omitempty"`
	ODataType                        *string                                  `json:"@odata.type,omitempty"`
	Reviewers                        *[]AccessReviewReviewerScope             `json:"reviewers,omitempty"`
	Scope                            *AccessReviewScope                       `json:"scope,omitempty"`
	Settings                         *AccessReviewScheduleSettings            `json:"settings,omitempty"`
	StageSettings                    *[]AccessReviewStageSettings             `json:"stageSettings,omitempty"`
	Status                           *string                                  `json:"status,omitempty"`
}
