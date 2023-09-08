package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistrantBase struct {
	Id         *string `json:"id,omitempty"`
	JoinWebUrl *string `json:"joinWebUrl,omitempty"`
	ODataType  *string `json:"@odata.type,omitempty"`
}
