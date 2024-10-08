package sourcecontrol

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceControlCreateOrUpdateProperties struct {
	AutoSync       *bool                                 `json:"autoSync,omitempty"`
	Branch         *string                               `json:"branch,omitempty"`
	Description    *string                               `json:"description,omitempty"`
	FolderPath     *string                               `json:"folderPath,omitempty"`
	PublishRunbook *bool                                 `json:"publishRunbook,omitempty"`
	RepoURL        *string                               `json:"repoUrl,omitempty"`
	SecurityToken  *SourceControlSecurityTokenProperties `json:"securityToken,omitempty"`
	SourceType     *SourceType                           `json:"sourceType,omitempty"`
}
