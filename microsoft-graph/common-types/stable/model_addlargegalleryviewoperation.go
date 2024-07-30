package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddLargeGalleryViewOperation struct {
	ClientContext *string                             `json:"clientContext,omitempty"`
	Id            *string                             `json:"id,omitempty"`
	ODataType     *string                             `json:"@odata.type,omitempty"`
	ResultInfo    *ResultInfo                         `json:"resultInfo,omitempty"`
	Status        *AddLargeGalleryViewOperationStatus `json:"status,omitempty"`
}
