package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyUserHistoryItem struct {
	Activity                *RiskUserActivity               `json:"activity,omitempty"`
	History                 *[]RiskyUserHistoryItem         `json:"history,omitempty"`
	Id                      *string                         `json:"id,omitempty"`
	InitiatedBy             *string                         `json:"initiatedBy,omitempty"`
	IsDeleted               *bool                           `json:"isDeleted,omitempty"`
	IsProcessing            *bool                           `json:"isProcessing,omitempty"`
	ODataType               *string                         `json:"@odata.type,omitempty"`
	RiskDetail              *RiskyUserHistoryItemRiskDetail `json:"riskDetail,omitempty"`
	RiskLastUpdatedDateTime *string                         `json:"riskLastUpdatedDateTime,omitempty"`
	RiskLevel               *RiskyUserHistoryItemRiskLevel  `json:"riskLevel,omitempty"`
	RiskState               *RiskyUserHistoryItemRiskState  `json:"riskState,omitempty"`
	UserDisplayName         *string                         `json:"userDisplayName,omitempty"`
	UserId                  *string                         `json:"userId,omitempty"`
	UserPrincipalName       *string                         `json:"userPrincipalName,omitempty"`
}
