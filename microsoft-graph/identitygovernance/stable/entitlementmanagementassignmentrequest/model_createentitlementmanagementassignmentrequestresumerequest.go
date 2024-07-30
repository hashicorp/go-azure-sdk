package entitlementmanagementassignmentrequest

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEntitlementManagementAssignmentRequestResumeRequest struct {
	Data   *CustomExtensionData `json:"data,omitempty"`
	Source *string              `json:"source,omitempty"`
	Type   *string              `json:"type,omitempty"`
}
