package siteonenotenotebook

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CopySiteOnenoteNotebookNotebookRequest struct {
	GroupId          nullable.Type[string] `json:"groupId,omitempty"`
	NotebookFolder   nullable.Type[string] `json:"notebookFolder,omitempty"`
	RenameAs         nullable.Type[string] `json:"renameAs,omitempty"`
	SiteCollectionId nullable.Type[string] `json:"siteCollectionId,omitempty"`
	SiteId           nullable.Type[string] `json:"siteId,omitempty"`
}
