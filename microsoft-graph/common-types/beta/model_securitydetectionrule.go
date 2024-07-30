package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDetectionRule struct {
	CreatedBy            *string                  `json:"createdBy,omitempty"`
	CreatedDateTime      *string                  `json:"createdDateTime,omitempty"`
	DetectionAction      *SecurityDetectionAction `json:"detectionAction,omitempty"`
	DisplayName          *string                  `json:"displayName,omitempty"`
	Id                   *string                  `json:"id,omitempty"`
	IsEnabled            *bool                    `json:"isEnabled,omitempty"`
	LastModifiedBy       *string                  `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                  `json:"lastModifiedDateTime,omitempty"`
	LastRunDetails       *SecurityRunDetails      `json:"lastRunDetails,omitempty"`
	ODataType            *string                  `json:"@odata.type,omitempty"`
	QueryCondition       *SecurityQueryCondition  `json:"queryCondition,omitempty"`
	Schedule             *SecurityRuleSchedule    `json:"schedule,omitempty"`
}
