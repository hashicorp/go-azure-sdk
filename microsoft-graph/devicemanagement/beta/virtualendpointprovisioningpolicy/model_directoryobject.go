package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryObject struct {
	DeletedDateTime *string `json:"deletedDateTime,omitempty"`
	Id              *string `json:"id,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
}
