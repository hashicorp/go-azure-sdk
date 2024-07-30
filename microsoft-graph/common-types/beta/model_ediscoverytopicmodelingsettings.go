package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryTopicModelingSettings struct {
	DynamicallyAdjustTopicCount *bool   `json:"dynamicallyAdjustTopicCount,omitempty"`
	IgnoreNumbers               *bool   `json:"ignoreNumbers,omitempty"`
	IsEnabled                   *bool   `json:"isEnabled,omitempty"`
	ODataType                   *string `json:"@odata.type,omitempty"`
	TopicCount                  *int64  `json:"topicCount,omitempty"`
}
