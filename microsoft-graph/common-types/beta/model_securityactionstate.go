package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityActionState struct {
	AppId           *string                    `json:"appId,omitempty"`
	ODataType       *string                    `json:"@odata.type,omitempty"`
	Status          *SecurityActionStateStatus `json:"status,omitempty"`
	UpdatedDateTime *string                    `json:"updatedDateTime,omitempty"`
	User            *string                    `json:"user,omitempty"`
}
