package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyServicePrincipalHistoryItem struct {
	Activity                *RiskServicePrincipalActivity               `json:"activity,omitempty"`
	AppId                   *string                                     `json:"appId,omitempty"`
	DisplayName             *string                                     `json:"displayName,omitempty"`
	History                 *[]RiskyServicePrincipalHistoryItem         `json:"history,omitempty"`
	Id                      *string                                     `json:"id,omitempty"`
	InitiatedBy             *string                                     `json:"initiatedBy,omitempty"`
	IsEnabled               *bool                                       `json:"isEnabled,omitempty"`
	IsProcessing            *bool                                       `json:"isProcessing,omitempty"`
	ODataType               *string                                     `json:"@odata.type,omitempty"`
	RiskDetail              *RiskyServicePrincipalHistoryItemRiskDetail `json:"riskDetail,omitempty"`
	RiskLastUpdatedDateTime *string                                     `json:"riskLastUpdatedDateTime,omitempty"`
	RiskLevel               *RiskyServicePrincipalHistoryItemRiskLevel  `json:"riskLevel,omitempty"`
	RiskState               *RiskyServicePrincipalHistoryItemRiskState  `json:"riskState,omitempty"`
	ServicePrincipalType    *string                                     `json:"servicePrincipalType,omitempty"`
}
