package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteHistory struct {
	AllowRedirect     *bool                                `json:"allowRedirect,omitempty"`
	Comment           *string                              `json:"comment,omitempty"`
	CompatibilityMode *BrowserSiteHistoryCompatibilityMode `json:"compatibilityMode,omitempty"`
	LastModifiedBy    *IdentitySet                         `json:"lastModifiedBy,omitempty"`
	MergeType         *BrowserSiteHistoryMergeType         `json:"mergeType,omitempty"`
	ODataType         *string                              `json:"@odata.type,omitempty"`
	PublishedDateTime *string                              `json:"publishedDateTime,omitempty"`
	TargetEnvironment *BrowserSiteHistoryTargetEnvironment `json:"targetEnvironment,omitempty"`
}
