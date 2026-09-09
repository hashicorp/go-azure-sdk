package recommendations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationRuleProperties struct {
	ActionName         *string            `json:"actionName,omitempty"`
	BladeName          *string            `json:"bladeName,omitempty"`
	CategoryTags       *[]string          `json:"categoryTags,omitempty"`
	Channels           *Channels          `json:"channels,omitempty"`
	Description        *string            `json:"description,omitempty"`
	DisplayName        *string            `json:"displayName,omitempty"`
	ExtensionName      *string            `json:"extensionName,omitempty"`
	ForwardLink        *string            `json:"forwardLink,omitempty"`
	IsDynamic          *bool              `json:"isDynamic,omitempty"`
	Level              *NotificationLevel `json:"level,omitempty"`
	Message            *string            `json:"message,omitempty"`
	RecommendationId   *string            `json:"recommendationId,omitempty"`
	RecommendationName *string            `json:"recommendationName,omitempty"`
}
