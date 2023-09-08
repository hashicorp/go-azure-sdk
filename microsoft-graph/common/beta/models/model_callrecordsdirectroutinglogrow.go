package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsDirectRoutingLogRow struct {
	CallEndSubReason              *int64  `json:"callEndSubReason,omitempty"`
	CallType                      *string `json:"callType,omitempty"`
	CalleeNumber                  *string `json:"calleeNumber,omitempty"`
	CallerNumber                  *string `json:"callerNumber,omitempty"`
	CorrelationId                 *string `json:"correlationId,omitempty"`
	Duration                      *int64  `json:"duration,omitempty"`
	EndDateTime                   *string `json:"endDateTime,omitempty"`
	FailureDateTime               *string `json:"failureDateTime,omitempty"`
	FinalSipCode                  *int64  `json:"finalSipCode,omitempty"`
	FinalSipCodePhrase            *string `json:"finalSipCodePhrase,omitempty"`
	Id                            *string `json:"id,omitempty"`
	InviteDateTime                *string `json:"inviteDateTime,omitempty"`
	MediaBypassEnabled            *bool   `json:"mediaBypassEnabled,omitempty"`
	MediaPathLocation             *string `json:"mediaPathLocation,omitempty"`
	ODataType                     *string `json:"@odata.type,omitempty"`
	OtherPartyCountryCode         *string `json:"otherPartyCountryCode,omitempty"`
	SignalingLocation             *string `json:"signalingLocation,omitempty"`
	StartDateTime                 *string `json:"startDateTime,omitempty"`
	SuccessfulCall                *bool   `json:"successfulCall,omitempty"`
	TrunkFullyQualifiedDomainName *string `json:"trunkFullyQualifiedDomainName,omitempty"`
	UserCountryCode               *string `json:"userCountryCode,omitempty"`
	UserDisplayName               *string `json:"userDisplayName,omitempty"`
	UserId                        *string `json:"userId,omitempty"`
	UserPrincipalName             *string `json:"userPrincipalName,omitempty"`
}
