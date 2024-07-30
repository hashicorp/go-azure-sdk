package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddressBookAccountTargetContent struct {
	AccountTargetEmails *[]string                            `json:"accountTargetEmails,omitempty"`
	ODataType           *string                              `json:"@odata.type,omitempty"`
	Type                *AddressBookAccountTargetContentType `json:"type,omitempty"`
}
