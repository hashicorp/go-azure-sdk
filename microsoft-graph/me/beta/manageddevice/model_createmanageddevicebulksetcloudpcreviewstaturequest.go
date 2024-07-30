package manageddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateManagedDeviceBulkSetCloudPCReviewStatuRequest struct {
	ManagedDeviceIds *[]string            `json:"managedDeviceIds,omitempty"`
	ReviewStatus     *CloudPCReviewStatus `json:"reviewStatus,omitempty"`
}
