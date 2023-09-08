package groupsitelistcontenttype

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateGroupByIdSiteByIdListByIdContentTypeByIdCopyToDefaultContentLocationRequest struct {
	DestinationFileName *string        `json:"destinationFileName,omitempty"`
	SourceFile          *ItemReference `json:"sourceFile,omitempty"`
}
