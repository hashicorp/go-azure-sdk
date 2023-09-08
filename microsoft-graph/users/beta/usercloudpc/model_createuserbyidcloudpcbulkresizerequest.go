package usercloudpc

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdCloudPCBulkResizeRequest struct {
	CloudPCIds          *[]string `json:"cloudPcIds,omitempty"`
	TargetServicePlanId *string   `json:"targetServicePlanId,omitempty"`
}
