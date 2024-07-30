package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsPstnCallLogRow struct {
	CallDurationSource      *CallRecordsPstnCallLogRowCallDurationSource `json:"callDurationSource,omitempty"`
	CallId                  *string                                      `json:"callId,omitempty"`
	CallType                *string                                      `json:"callType,omitempty"`
	CalleeNumber            *string                                      `json:"calleeNumber,omitempty"`
	CallerNumber            *string                                      `json:"callerNumber,omitempty"`
	Charge                  *float64                                     `json:"charge,omitempty"`
	ClientLocalIpV4Address  *string                                      `json:"clientLocalIpV4Address,omitempty"`
	ClientLocalIpV6Address  *string                                      `json:"clientLocalIpV6Address,omitempty"`
	ClientPublicIpV4Address *string                                      `json:"clientPublicIpV4Address,omitempty"`
	ClientPublicIpV6Address *string                                      `json:"clientPublicIpV6Address,omitempty"`
	ConferenceId            *string                                      `json:"conferenceId,omitempty"`
	ConnectionCharge        *float64                                     `json:"connectionCharge,omitempty"`
	Currency                *string                                      `json:"currency,omitempty"`
	DestinationContext      *string                                      `json:"destinationContext,omitempty"`
	DestinationName         *string                                      `json:"destinationName,omitempty"`
	Duration                *int64                                       `json:"duration,omitempty"`
	EndDateTime             *string                                      `json:"endDateTime,omitempty"`
	Id                      *string                                      `json:"id,omitempty"`
	InventoryType           *string                                      `json:"inventoryType,omitempty"`
	LicenseCapability       *string                                      `json:"licenseCapability,omitempty"`
	ODataType               *string                                      `json:"@odata.type,omitempty"`
	Operator                *string                                      `json:"operator,omitempty"`
	OtherPartyCountryCode   *string                                      `json:"otherPartyCountryCode,omitempty"`
	StartDateTime           *string                                      `json:"startDateTime,omitempty"`
	TenantCountryCode       *string                                      `json:"tenantCountryCode,omitempty"`
	UsageCountryCode        *string                                      `json:"usageCountryCode,omitempty"`
	UserDisplayName         *string                                      `json:"userDisplayName,omitempty"`
	UserId                  *string                                      `json:"userId,omitempty"`
	UserPrincipalName       *string                                      `json:"userPrincipalName,omitempty"`
}
