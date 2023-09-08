package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityArticle struct {
	Body                *FormattedContent           `json:"body,omitempty"`
	CreatedDateTime     *string                     `json:"createdDateTime,omitempty"`
	Id                  *string                     `json:"id,omitempty"`
	ImageUrl            *string                     `json:"imageUrl,omitempty"`
	Indicators          *[]SecurityArticleIndicator `json:"indicators,omitempty"`
	IsFeatured          *bool                       `json:"isFeatured,omitempty"`
	LastUpdatedDateTime *string                     `json:"lastUpdatedDateTime,omitempty"`
	ODataType           *string                     `json:"@odata.type,omitempty"`
	Summary             *FormattedContent           `json:"summary,omitempty"`
	Tags                *[]string                   `json:"tags,omitempty"`
	Title               *string                     `json:"title,omitempty"`
}
