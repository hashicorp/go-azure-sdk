package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseSettings struct {
	Id                  *string                                `json:"id,omitempty"`
	ODataType           *string                                `json:"@odata.type,omitempty"`
	Ocr                 *EdiscoveryOcrSettings                 `json:"ocr,omitempty"`
	RedundancyDetection *EdiscoveryRedundancyDetectionSettings `json:"redundancyDetection,omitempty"`
	TopicModeling       *EdiscoveryTopicModelingSettings       `json:"topicModeling,omitempty"`
}
