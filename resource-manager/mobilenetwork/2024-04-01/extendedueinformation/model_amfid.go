package extendedueinformation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AmfId struct {
	Pointer  int64 `json:"pointer"`
	RegionId int64 `json:"regionId"`
	SetId    int64 `json:"setId"`
}
