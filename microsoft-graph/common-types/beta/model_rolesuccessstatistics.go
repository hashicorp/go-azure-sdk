package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleSuccessStatistics struct {
	ODataType        *string `json:"@odata.type,omitempty"`
	PermanentFail    *int64  `json:"permanentFail,omitempty"`
	PermanentSuccess *int64  `json:"permanentSuccess,omitempty"`
	RemoveFail       *int64  `json:"removeFail,omitempty"`
	RemoveSuccess    *int64  `json:"removeSuccess,omitempty"`
	RoleId           *string `json:"roleId,omitempty"`
	RoleName         *string `json:"roleName,omitempty"`
	TemporaryFail    *int64  `json:"temporaryFail,omitempty"`
	TemporarySuccess *int64  `json:"temporarySuccess,omitempty"`
	UnknownFail      *int64  `json:"unknownFail,omitempty"`
}
