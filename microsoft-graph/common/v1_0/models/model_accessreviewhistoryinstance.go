package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewHistoryInstance struct {
	DownloadUri                      *string                            `json:"downloadUri,omitempty"`
	ExpirationDateTime               *string                            `json:"expirationDateTime,omitempty"`
	FulfilledDateTime                *string                            `json:"fulfilledDateTime,omitempty"`
	Id                               *string                            `json:"id,omitempty"`
	ODataType                        *string                            `json:"@odata.type,omitempty"`
	ReviewHistoryPeriodEndDateTime   *string                            `json:"reviewHistoryPeriodEndDateTime,omitempty"`
	ReviewHistoryPeriodStartDateTime *string                            `json:"reviewHistoryPeriodStartDateTime,omitempty"`
	RunDateTime                      *string                            `json:"runDateTime,omitempty"`
	Status                           *AccessReviewHistoryInstanceStatus `json:"status,omitempty"`
}
