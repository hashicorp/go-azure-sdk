package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignIn struct {
	AppDisplayName                   *string                           `json:"appDisplayName,omitempty"`
	AppId                            *string                           `json:"appId,omitempty"`
	AppliedConditionalAccessPolicies *[]AppliedConditionalAccessPolicy `json:"appliedConditionalAccessPolicies,omitempty"`
	ClientAppUsed                    *string                           `json:"clientAppUsed,omitempty"`
	ConditionalAccessStatus          *SignInConditionalAccessStatus    `json:"conditionalAccessStatus,omitempty"`
	CorrelationId                    *string                           `json:"correlationId,omitempty"`
	CreatedDateTime                  *string                           `json:"createdDateTime,omitempty"`
	DeviceDetail                     *DeviceDetail                     `json:"deviceDetail,omitempty"`
	Id                               *string                           `json:"id,omitempty"`
	IpAddress                        *string                           `json:"ipAddress,omitempty"`
	IsInteractive                    *bool                             `json:"isInteractive,omitempty"`
	Location                         *SignInLocation                   `json:"location,omitempty"`
	ODataType                        *string                           `json:"@odata.type,omitempty"`
	ResourceDisplayName              *string                           `json:"resourceDisplayName,omitempty"`
	ResourceId                       *string                           `json:"resourceId,omitempty"`
	RiskDetail                       *SignInRiskDetail                 `json:"riskDetail,omitempty"`
	RiskEventTypes                   *[]SignInRiskEventTypes           `json:"riskEventTypes,omitempty"`
	RiskEventTypesv2                 *[]string                         `json:"riskEventTypes_v2,omitempty"`
	RiskLevelAggregated              *SignInRiskLevelAggregated        `json:"riskLevelAggregated,omitempty"`
	RiskLevelDuringSignIn            *SignInRiskLevelDuringSignIn      `json:"riskLevelDuringSignIn,omitempty"`
	RiskState                        *SignInRiskState                  `json:"riskState,omitempty"`
	Status                           *SignInStatus                     `json:"status,omitempty"`
	UserDisplayName                  *string                           `json:"userDisplayName,omitempty"`
	UserId                           *string                           `json:"userId,omitempty"`
	UserPrincipalName                *string                           `json:"userPrincipalName,omitempty"`
}
