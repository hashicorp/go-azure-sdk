package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskDetection struct {
	Activity            *RiskDetectionActivity            `json:"activity,omitempty"`
	ActivityDateTime    *string                           `json:"activityDateTime,omitempty"`
	AdditionalInfo      *string                           `json:"additionalInfo,omitempty"`
	CorrelationId       *string                           `json:"correlationId,omitempty"`
	DetectedDateTime    *string                           `json:"detectedDateTime,omitempty"`
	DetectionTimingType *RiskDetectionDetectionTimingType `json:"detectionTimingType,omitempty"`
	Id                  *string                           `json:"id,omitempty"`
	IpAddress           *string                           `json:"ipAddress,omitempty"`
	LastUpdatedDateTime *string                           `json:"lastUpdatedDateTime,omitempty"`
	Location            *SignInLocation                   `json:"location,omitempty"`
	ODataType           *string                           `json:"@odata.type,omitempty"`
	RequestId           *string                           `json:"requestId,omitempty"`
	RiskDetail          *RiskDetectionRiskDetail          `json:"riskDetail,omitempty"`
	RiskEventType       *string                           `json:"riskEventType,omitempty"`
	RiskLevel           *RiskDetectionRiskLevel           `json:"riskLevel,omitempty"`
	RiskState           *RiskDetectionRiskState           `json:"riskState,omitempty"`
	Source              *string                           `json:"source,omitempty"`
	TokenIssuerType     *RiskDetectionTokenIssuerType     `json:"tokenIssuerType,omitempty"`
	UserDisplayName     *string                           `json:"userDisplayName,omitempty"`
	UserId              *string                           `json:"userId,omitempty"`
	UserPrincipalName   *string                           `json:"userPrincipalName,omitempty"`
}
