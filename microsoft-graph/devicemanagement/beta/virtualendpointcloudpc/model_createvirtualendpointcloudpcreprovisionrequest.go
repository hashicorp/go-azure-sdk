package virtualendpointcloudpc

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateVirtualEndpointCloudPCReprovisionRequest struct {
	OsVersion       *CreateVirtualEndpointCloudPCReprovisionRequestOsVersion       `json:"osVersion,omitempty"`
	UserAccountType *CreateVirtualEndpointCloudPCReprovisionRequestUserAccountType `json:"userAccountType,omitempty"`
}
