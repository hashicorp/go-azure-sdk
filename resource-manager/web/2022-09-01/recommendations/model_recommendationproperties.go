package recommendations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationProperties struct {
	ActionName                 *string            `json:"actionName,omitempty"`
	BladeName                  *string            `json:"bladeName,omitempty"`
	CategoryTags               *[]string          `json:"categoryTags,omitempty"`
	Channels                   *Channels          `json:"channels,omitempty"`
	CreationTime               *string            `json:"creationTime,omitempty"`
	DisplayName                *string            `json:"displayName,omitempty"`
	Enabled                    *int64             `json:"enabled,omitempty"`
	EndTime                    *string            `json:"endTime,omitempty"`
	ExtensionName              *string            `json:"extensionName,omitempty"`
	ForwardLink                *string            `json:"forwardLink,omitempty"`
	IsDynamic                  *bool              `json:"isDynamic,omitempty"`
	Level                      *NotificationLevel `json:"level,omitempty"`
	Message                    *string            `json:"message,omitempty"`
	NextNotificationTime       *string            `json:"nextNotificationTime,omitempty"`
	NotificationExpirationTime *string            `json:"notificationExpirationTime,omitempty"`
	NotifiedTime               *string            `json:"notifiedTime,omitempty"`
	RecommendationId           *string            `json:"recommendationId,omitempty"`
	ResourceId                 *string            `json:"resourceId,omitempty"`
	ResourceScope              *ResourceScopeType `json:"resourceScope,omitempty"`
	RuleName                   *string            `json:"ruleName,omitempty"`
	Score                      *float64           `json:"score,omitempty"`
	StartTime                  *string            `json:"startTime,omitempty"`
	States                     *[]string          `json:"states,omitempty"`
}
