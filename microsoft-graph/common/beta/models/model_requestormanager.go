package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequestorManager struct {
	IsBackup     *bool   `json:"isBackup,omitempty"`
	ManagerLevel *int64  `json:"managerLevel,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
}
