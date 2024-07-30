package cloudpc

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateCloudPCReprovisionRequest struct {
	OsVersion       *CreateCloudPCReprovisionRequestOsVersion       `json:"osVersion,omitempty"`
	UserAccountType *CreateCloudPCReprovisionRequestUserAccountType `json:"userAccountType,omitempty"`
}
