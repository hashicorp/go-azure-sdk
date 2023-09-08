package mecloudpc

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeCloudPCByIdReprovisionRequest struct {
	OsVersion       *CreateMeCloudPCByIdReprovisionRequestOsVersion       `json:"osVersion,omitempty"`
	UserAccountType *CreateMeCloudPCByIdReprovisionRequestUserAccountType `json:"userAccountType,omitempty"`
}
