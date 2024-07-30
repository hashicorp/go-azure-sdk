package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TabUpdatedEventMessageDetail struct {
	Initiator *IdentitySet `json:"initiator,omitempty"`
	ODataType *string      `json:"@odata.type,omitempty"`
	TabId     *string      `json:"tabId,omitempty"`
}
