package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalMeetingRegistrant struct {
	Id         *string `json:"id,omitempty"`
	JoinWebUrl *string `json:"joinWebUrl,omitempty"`
	ODataType  *string `json:"@odata.type,omitempty"`
	TenantId   *string `json:"tenantId,omitempty"`
	UserId     *string `json:"userId,omitempty"`
}
