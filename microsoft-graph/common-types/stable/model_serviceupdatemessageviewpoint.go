package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceUpdateMessageViewpoint struct {
	IsArchived  *bool   `json:"isArchived,omitempty"`
	IsFavorited *bool   `json:"isFavorited,omitempty"`
	IsRead      *bool   `json:"isRead,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
}
