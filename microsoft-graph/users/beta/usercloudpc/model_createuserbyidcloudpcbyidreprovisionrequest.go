package usercloudpc

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdCloudPCByIdReprovisionRequest struct {
	OsVersion       *CreateUserByIdCloudPCByIdReprovisionRequestOsVersion       `json:"osVersion,omitempty"`
	UserAccountType *CreateUserByIdCloudPCByIdReprovisionRequestUserAccountType `json:"userAccountType,omitempty"`
}
