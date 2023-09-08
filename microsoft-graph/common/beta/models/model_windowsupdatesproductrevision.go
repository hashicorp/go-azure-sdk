package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesProductRevision struct {
	DisplayName          *string                             `json:"displayName,omitempty"`
	Id                   *string                             `json:"id,omitempty"`
	KnowledgeBaseArticle *WindowsUpdatesKnowledgeBaseArticle `json:"knowledgeBaseArticle,omitempty"`
	ODataType            *string                             `json:"@odata.type,omitempty"`
	OsBuild              *BuildVersionDetails                `json:"osBuild,omitempty"`
	Product              *string                             `json:"product,omitempty"`
	ReleaseDateTime      *string                             `json:"releaseDateTime,omitempty"`
	Version              *string                             `json:"version,omitempty"`
}
