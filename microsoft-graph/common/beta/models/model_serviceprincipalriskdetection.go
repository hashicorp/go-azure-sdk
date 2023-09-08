package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalRiskDetection struct {
	Activity                    *ServicePrincipalRiskDetectionActivity            `json:"activity,omitempty"`
	ActivityDateTime            *string                                           `json:"activityDateTime,omitempty"`
	AdditionalInfo              *string                                           `json:"additionalInfo,omitempty"`
	AppId                       *string                                           `json:"appId,omitempty"`
	CorrelationId               *string                                           `json:"correlationId,omitempty"`
	DetectedDateTime            *string                                           `json:"detectedDateTime,omitempty"`
	DetectionTimingType         *ServicePrincipalRiskDetectionDetectionTimingType `json:"detectionTimingType,omitempty"`
	Id                          *string                                           `json:"id,omitempty"`
	IpAddress                   *string                                           `json:"ipAddress,omitempty"`
	KeyIds                      *[]string                                         `json:"keyIds,omitempty"`
	LastUpdatedDateTime         *string                                           `json:"lastUpdatedDateTime,omitempty"`
	Location                    *SignInLocation                                   `json:"location,omitempty"`
	ODataType                   *string                                           `json:"@odata.type,omitempty"`
	RequestId                   *string                                           `json:"requestId,omitempty"`
	RiskDetail                  *ServicePrincipalRiskDetectionRiskDetail          `json:"riskDetail,omitempty"`
	RiskEventType               *string                                           `json:"riskEventType,omitempty"`
	RiskLevel                   *ServicePrincipalRiskDetectionRiskLevel           `json:"riskLevel,omitempty"`
	RiskState                   *ServicePrincipalRiskDetectionRiskState           `json:"riskState,omitempty"`
	ServicePrincipalDisplayName *string                                           `json:"servicePrincipalDisplayName,omitempty"`
	ServicePrincipalId          *string                                           `json:"servicePrincipalId,omitempty"`
	Source                      *string                                           `json:"source,omitempty"`
	TokenIssuerType             *ServicePrincipalRiskDetectionTokenIssuerType     `json:"tokenIssuerType,omitempty"`
}
