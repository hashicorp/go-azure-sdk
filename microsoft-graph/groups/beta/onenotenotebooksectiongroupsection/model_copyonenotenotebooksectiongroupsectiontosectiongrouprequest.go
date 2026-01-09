package onenotenotebooksectiongroupsection

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CopyOnenoteNotebookSectionGroupSectionToSectionGroupRequest struct {
	GroupId          nullable.Type[string] `json:"groupId,omitempty"`
	Id               nullable.Type[string] `json:"id,omitempty"`
	RenameAs         nullable.Type[string] `json:"renameAs,omitempty"`
	SiteCollectionId nullable.Type[string] `json:"siteCollectionId,omitempty"`
	SiteId           nullable.Type[string] `json:"siteId,omitempty"`
}
