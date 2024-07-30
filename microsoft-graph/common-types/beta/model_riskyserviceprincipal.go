package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyServicePrincipal struct {
	AccountEnabled          *bool                               `json:"accountEnabled,omitempty"`
	AppId                   *string                             `json:"appId,omitempty"`
	DisplayName             *string                             `json:"displayName,omitempty"`
	History                 *[]RiskyServicePrincipalHistoryItem `json:"history,omitempty"`
	Id                      *string                             `json:"id,omitempty"`
	IsEnabled               *bool                               `json:"isEnabled,omitempty"`
	IsProcessing            *bool                               `json:"isProcessing,omitempty"`
	ODataType               *string                             `json:"@odata.type,omitempty"`
	RiskDetail              *RiskyServicePrincipalRiskDetail    `json:"riskDetail,omitempty"`
	RiskLastUpdatedDateTime *string                             `json:"riskLastUpdatedDateTime,omitempty"`
	RiskLevel               *RiskyServicePrincipalRiskLevel     `json:"riskLevel,omitempty"`
	RiskState               *RiskyServicePrincipalRiskState     `json:"riskState,omitempty"`
	ServicePrincipalType    *string                             `json:"servicePrincipalType,omitempty"`
}
