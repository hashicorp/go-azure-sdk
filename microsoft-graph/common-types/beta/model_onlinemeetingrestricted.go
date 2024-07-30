package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingRestricted struct {
	ContentSharingDisabled *OnlineMeetingRestrictedContentSharingDisabled `json:"contentSharingDisabled,omitempty"`
	ODataType              *string                                        `json:"@odata.type,omitempty"`
	VideoDisabled          *OnlineMeetingRestrictedVideoDisabled          `json:"videoDisabled,omitempty"`
}
