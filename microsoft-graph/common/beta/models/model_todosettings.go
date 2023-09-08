package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoSettings struct {
	IsExternalJoinEnabled     *bool   `json:"isExternalJoinEnabled,omitempty"`
	IsExternalShareEnabled    *bool   `json:"isExternalShareEnabled,omitempty"`
	IsPushNotificationEnabled *bool   `json:"isPushNotificationEnabled,omitempty"`
	ODataType                 *string `json:"@odata.type,omitempty"`
}
