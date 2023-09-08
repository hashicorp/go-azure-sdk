package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewHistoryDefinition struct {
	CreatedBy                        *UserIdentity                             `json:"createdBy,omitempty"`
	CreatedDateTime                  *string                                   `json:"createdDateTime,omitempty"`
	Decisions                        *[]AccessReviewHistoryDefinitionDecisions `json:"decisions,omitempty"`
	DisplayName                      *string                                   `json:"displayName,omitempty"`
	DownloadUri                      *string                                   `json:"downloadUri,omitempty"`
	FulfilledDateTime                *string                                   `json:"fulfilledDateTime,omitempty"`
	Id                               *string                                   `json:"id,omitempty"`
	Instances                        *[]AccessReviewHistoryInstance            `json:"instances,omitempty"`
	ODataType                        *string                                   `json:"@odata.type,omitempty"`
	ReviewHistoryPeriodEndDateTime   *string                                   `json:"reviewHistoryPeriodEndDateTime,omitempty"`
	ReviewHistoryPeriodStartDateTime *string                                   `json:"reviewHistoryPeriodStartDateTime,omitempty"`
	ScheduleSettings                 *AccessReviewHistoryScheduleSettings      `json:"scheduleSettings,omitempty"`
	Scopes                           *[]AccessReviewScope                      `json:"scopes,omitempty"`
	Status                           *AccessReviewHistoryDefinitionStatus      `json:"status,omitempty"`
}
