package metadata

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetadataProperties struct {
	Author           *MetadataAuthor       `json:"author,omitempty"`
	Categories       *MetadataCategories   `json:"categories,omitempty"`
	ContentId        *string               `json:"contentId,omitempty"`
	Dependencies     *MetadataDependencies `json:"dependencies,omitempty"`
	FirstPublishDate *string               `json:"firstPublishDate,omitempty"`
	Kind             Kind                  `json:"kind"`
	LastPublishDate  *string               `json:"lastPublishDate,omitempty"`
	ParentId         string                `json:"parentId"`
	Providers        *[]string             `json:"providers,omitempty"`
	Source           *MetadataSource       `json:"source,omitempty"`
	Support          *MetadataSupport      `json:"support,omitempty"`
	Version          *string               `json:"version,omitempty"`
}
