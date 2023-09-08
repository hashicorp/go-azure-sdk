package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryCaseSettings struct {
	Id                  *string                              `json:"id,omitempty"`
	ODataType           *string                              `json:"@odata.type,omitempty"`
	Ocr                 *SecurityOcrSettings                 `json:"ocr,omitempty"`
	RedundancyDetection *SecurityRedundancyDetectionSettings `json:"redundancyDetection,omitempty"`
	TopicModeling       *SecurityTopicModelingSettings       `json:"topicModeling,omitempty"`
}
