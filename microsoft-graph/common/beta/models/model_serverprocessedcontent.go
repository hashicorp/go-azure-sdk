package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerProcessedContent struct {
	ComponentDependencies *[]MetaDataKeyStringPair `json:"componentDependencies,omitempty"`
	CustomMetadata        *[]MetaDataKeyValuePair  `json:"customMetadata,omitempty"`
	HtmlStrings           *[]MetaDataKeyStringPair `json:"htmlStrings,omitempty"`
	ImageSources          *[]MetaDataKeyStringPair `json:"imageSources,omitempty"`
	Links                 *[]MetaDataKeyStringPair `json:"links,omitempty"`
	ODataType             *string                  `json:"@odata.type,omitempty"`
	SearchablePlainTexts  *[]MetaDataKeyStringPair `json:"searchablePlainTexts,omitempty"`
}
