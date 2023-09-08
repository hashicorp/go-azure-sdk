package groupsiteonenotenotebook

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateGroupByIdSiteByIdOnenoteNotebookByIdCopyNotebookRequest struct {
	GroupId          *string `json:"groupId,omitempty"`
	NotebookFolder   *string `json:"notebookFolder,omitempty"`
	RenameAs         *string `json:"renameAs,omitempty"`
	SiteCollectionId *string `json:"siteCollectionId,omitempty"`
	SiteId           *string `json:"siteId,omitempty"`
}
