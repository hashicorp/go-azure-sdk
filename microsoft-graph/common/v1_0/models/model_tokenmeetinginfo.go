package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TokenMeetingInfo struct {
	ODataType *string `json:"@odata.type,omitempty"`
	Token     *string `json:"token,omitempty"`
}
