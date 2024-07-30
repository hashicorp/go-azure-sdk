package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotPolicyStatusDetail struct {
	ComplianceStatus          *DeviceManagementAutopilotPolicyStatusDetailComplianceStatus `json:"complianceStatus,omitempty"`
	DisplayName               *string                                                      `json:"displayName,omitempty"`
	ErrorCode                 *int64                                                       `json:"errorCode,omitempty"`
	Id                        *string                                                      `json:"id,omitempty"`
	LastReportedDateTime      *string                                                      `json:"lastReportedDateTime,omitempty"`
	ODataType                 *string                                                      `json:"@odata.type,omitempty"`
	PolicyType                *DeviceManagementAutopilotPolicyStatusDetailPolicyType       `json:"policyType,omitempty"`
	TrackedOnEnrollmentStatus *bool                                                        `json:"trackedOnEnrollmentStatus,omitempty"`
}
