package mecloudpc

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateMeCloudPCsBulkResizeRequest struct {
	CloudPCIds          *[]string `json:"cloudPcIds,omitempty"`
	TargetServicePlanId *string   `json:"targetServicePlanId,omitempty"`
}
