package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RelyingPartyDetailedSummary struct {
	FailedSignInCount          *int64                                      `json:"failedSignInCount,omitempty"`
	Id                         *string                                     `json:"id,omitempty"`
	MigrationStatus            *RelyingPartyDetailedSummaryMigrationStatus `json:"migrationStatus,omitempty"`
	MigrationValidationDetails *[]KeyValuePair                             `json:"migrationValidationDetails,omitempty"`
	ODataType                  *string                                     `json:"@odata.type,omitempty"`
	RelyingPartyId             *string                                     `json:"relyingPartyId,omitempty"`
	RelyingPartyName           *string                                     `json:"relyingPartyName,omitempty"`
	ReplyUrls                  *[]string                                   `json:"replyUrls,omitempty"`
	ServiceId                  *string                                     `json:"serviceId,omitempty"`
	SuccessfulSignInCount      *int64                                      `json:"successfulSignInCount,omitempty"`
	TotalSignInCount           *int64                                      `json:"totalSignInCount,omitempty"`
	UniqueUserCount            *int64                                      `json:"uniqueUserCount,omitempty"`
}
