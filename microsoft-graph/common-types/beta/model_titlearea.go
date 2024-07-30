package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TitleArea struct {
	AlternativeText         *string                 `json:"alternativeText,omitempty"`
	EnableGradientEffect    *bool                   `json:"enableGradientEffect,omitempty"`
	ImageWebUrl             *string                 `json:"imageWebUrl,omitempty"`
	Layout                  *TitleAreaLayout        `json:"layout,omitempty"`
	ODataType               *string                 `json:"@odata.type,omitempty"`
	ServerProcessedContent  *ServerProcessedContent `json:"serverProcessedContent,omitempty"`
	ShowAuthor              *bool                   `json:"showAuthor,omitempty"`
	ShowPublishedDate       *bool                   `json:"showPublishedDate,omitempty"`
	ShowTextBlockAboveTitle *bool                   `json:"showTextBlockAboveTitle,omitempty"`
	TextAboveTitle          *string                 `json:"textAboveTitle,omitempty"`
	TextAlignment           *TitleAreaTextAlignment `json:"textAlignment,omitempty"`
}
