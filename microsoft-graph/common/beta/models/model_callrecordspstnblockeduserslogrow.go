package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsPstnBlockedUsersLogRow struct {
	BlockDateTime       *string                                         `json:"blockDateTime,omitempty"`
	BlockReason         *string                                         `json:"blockReason,omitempty"`
	ODataType           *string                                         `json:"@odata.type,omitempty"`
	RemediationId       *string                                         `json:"remediationId,omitempty"`
	UserBlockMode       *CallRecordsPstnBlockedUsersLogRowUserBlockMode `json:"userBlockMode,omitempty"`
	UserDisplayName     *string                                         `json:"userDisplayName,omitempty"`
	UserId              *string                                         `json:"userId,omitempty"`
	UserPrincipalName   *string                                         `json:"userPrincipalName,omitempty"`
	UserTelephoneNumber *string                                         `json:"userTelephoneNumber,omitempty"`
}
