package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdateCatalogItem struct {
	Classification   *WindowsQualityUpdateCatalogItemClassification `json:"classification,omitempty"`
	DisplayName      *string                                        `json:"displayName,omitempty"`
	EndOfSupportDate *string                                        `json:"endOfSupportDate,omitempty"`
	Id               *string                                        `json:"id,omitempty"`
	IsExpeditable    *bool                                          `json:"isExpeditable,omitempty"`
	KbArticleId      *string                                        `json:"kbArticleId,omitempty"`
	ODataType        *string                                        `json:"@odata.type,omitempty"`
	ReleaseDateTime  *string                                        `json:"releaseDateTime,omitempty"`
}
