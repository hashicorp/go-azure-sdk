package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTopicModelingSettingsOperationPredicate struct {
	DynamicallyAdjustTopicCount *bool
	IgnoreNumbers               *bool
	IsEnabled                   *bool
	ODataType                   *string
	TopicCount                  *int64
}

func (p SecurityTopicModelingSettingsOperationPredicate) Matches(input SecurityTopicModelingSettings) bool {

	if p.DynamicallyAdjustTopicCount != nil && (input.DynamicallyAdjustTopicCount == nil || *p.DynamicallyAdjustTopicCount != *input.DynamicallyAdjustTopicCount) {
		return false
	}

	if p.IgnoreNumbers != nil && (input.IgnoreNumbers == nil || *p.IgnoreNumbers != *input.IgnoreNumbers) {
		return false
	}

	if p.IsEnabled != nil && (input.IsEnabled == nil || *p.IsEnabled != *input.IsEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TopicCount != nil && (input.TopicCount == nil || *p.TopicCount != *input.TopicCount) {
		return false
	}

	return true
}
