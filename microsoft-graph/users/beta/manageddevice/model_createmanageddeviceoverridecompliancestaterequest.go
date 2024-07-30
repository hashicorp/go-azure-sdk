package manageddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateManagedDeviceOverrideComplianceStateRequest struct {
	ComplianceState *CreateManagedDeviceOverrideComplianceStateRequestComplianceState `json:"complianceState,omitempty"`
	RemediationUrl  *string                                                           `json:"remediationUrl,omitempty"`
}
