package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventPresenter struct {
	Email            *string                       `json:"email,omitempty"`
	Id               *string                       `json:"id,omitempty"`
	Identity         *CommunicationsUserIdentity   `json:"identity,omitempty"`
	ODataType        *string                       `json:"@odata.type,omitempty"`
	PresenterDetails *VirtualEventPresenterDetails `json:"presenterDetails,omitempty"`
	ProfilePhoto     *string                       `json:"profilePhoto,omitempty"`
	Sessions         *[]VirtualEventSession        `json:"sessions,omitempty"`
}
