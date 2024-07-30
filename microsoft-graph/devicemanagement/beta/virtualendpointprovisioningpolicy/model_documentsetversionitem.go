package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DocumentSetVersionItem struct {
	ItemId    *string `json:"itemId,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Title     *string `json:"title,omitempty"`
	VersionId *string `json:"versionId,omitempty"`
}
