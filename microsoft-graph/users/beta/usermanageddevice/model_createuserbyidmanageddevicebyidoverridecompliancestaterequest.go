package usermanageddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdManagedDeviceByIdOverrideComplianceStateRequest struct {
	ComplianceState *CreateUserByIdManagedDeviceByIdOverrideComplianceStateRequestComplianceState `json:"complianceState,omitempty"`
	RemediationUrl  *string                                                                       `json:"remediationUrl,omitempty"`
}
