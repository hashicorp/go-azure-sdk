package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertRecord struct {
	AlertImpact         *DeviceManagementAlertImpact                  `json:"alertImpact,omitempty"`
	AlertRuleId         *string                                       `json:"alertRuleId,omitempty"`
	AlertRuleTemplate   *DeviceManagementAlertRecordAlertRuleTemplate `json:"alertRuleTemplate,omitempty"`
	DetectedDateTime    *string                                       `json:"detectedDateTime,omitempty"`
	DisplayName         *string                                       `json:"displayName,omitempty"`
	Id                  *string                                       `json:"id,omitempty"`
	LastUpdatedDateTime *string                                       `json:"lastUpdatedDateTime,omitempty"`
	ODataType           *string                                       `json:"@odata.type,omitempty"`
	ResolvedDateTime    *string                                       `json:"resolvedDateTime,omitempty"`
	Severity            *DeviceManagementAlertRecordSeverity          `json:"severity,omitempty"`
	Status              *DeviceManagementAlertRecordStatus            `json:"status,omitempty"`
}
