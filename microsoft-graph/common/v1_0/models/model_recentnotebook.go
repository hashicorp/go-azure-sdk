package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecentNotebook struct {
	DisplayName      *string                      `json:"displayName,omitempty"`
	LastAccessedTime *string                      `json:"lastAccessedTime,omitempty"`
	Links            *RecentNotebookLinks         `json:"links,omitempty"`
	ODataType        *string                      `json:"@odata.type,omitempty"`
	SourceService    *RecentNotebookSourceService `json:"sourceService,omitempty"`
}
