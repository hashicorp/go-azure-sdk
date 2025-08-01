package monitors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountInfo struct {
	AccountId   *string `json:"accountId,omitempty"`
	CompanyName *string `json:"companyName,omitempty"`
	RegionId    *string `json:"regionId,omitempty"`
}
