package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TextClassificationRequest struct {
	ContentMetaData          *ClassificationRequestContentMetaData              `json:"contentMetaData,omitempty"`
	FileExtension            *string                                            `json:"fileExtension,omitempty"`
	Id                       *string                                            `json:"id,omitempty"`
	MatchTolerancesToInclude *TextClassificationRequestMatchTolerancesToInclude `json:"matchTolerancesToInclude,omitempty"`
	ODataType                *string                                            `json:"@odata.type,omitempty"`
	ScopesToRun              *TextClassificationRequestScopesToRun              `json:"scopesToRun,omitempty"`
	SensitiveTypeIds         *[]string                                          `json:"sensitiveTypeIds,omitempty"`
	Text                     *string                                            `json:"text,omitempty"`
}
