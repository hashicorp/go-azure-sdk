package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedSignupStatus struct {
	Id           *string                       `json:"id,omitempty"`
	IsRegistered *bool                         `json:"isRegistered,omitempty"`
	ODataType    *string                       `json:"@odata.type,omitempty"`
	Status       *PrivilegedSignupStatusStatus `json:"status,omitempty"`
}
