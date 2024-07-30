package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataCollectionInfo struct {
	Entitlements *EntitlementsDataCollectionInfo `json:"entitlements,omitempty"`
	Id           *string                         `json:"id,omitempty"`
	ODataType    *string                         `json:"@odata.type,omitempty"`
}
