package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemSource struct {
	Application *DriveItemSourceApplication `json:"application,omitempty"`
	ExternalId  *string                     `json:"externalId,omitempty"`
	ODataType   *string                     `json:"@odata.type,omitempty"`
}
