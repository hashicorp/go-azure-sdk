package sourcecontrol

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceControlProperties struct {
	AutoSync         *bool       `json:"autoSync,omitempty"`
	Branch           *string     `json:"branch,omitempty"`
	CreationTime     *string     `json:"creationTime,omitempty"`
	Description      *string     `json:"description,omitempty"`
	FolderPath       *string     `json:"folderPath,omitempty"`
	LastModifiedTime *string     `json:"lastModifiedTime,omitempty"`
	PublishRunbook   *bool       `json:"publishRunbook,omitempty"`
	RepoUrl          *string     `json:"repoUrl,omitempty"`
	SourceType       *SourceType `json:"sourceType,omitempty"`
}
