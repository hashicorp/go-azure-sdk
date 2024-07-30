package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HasPayloadLinkResultItem struct {
	Error     *string                          `json:"error,omitempty"`
	HasLink   *bool                            `json:"hasLink,omitempty"`
	ODataType *string                          `json:"@odata.type,omitempty"`
	PayloadId *string                          `json:"payloadId,omitempty"`
	Sources   *HasPayloadLinkResultItemSources `json:"sources,omitempty"`
}
