package recommendations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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

func (o *RecommendationProperties) GetCreationTimeAsTime() (*time.Time, error) {
	if o.CreationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RecommendationProperties) SetCreationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationTime = &formatted
}

func (o *RecommendationProperties) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RecommendationProperties) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *RecommendationProperties) GetNextNotificationTimeAsTime() (*time.Time, error) {
	if o.NextNotificationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.NextNotificationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RecommendationProperties) SetNextNotificationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.NextNotificationTime = &formatted
}

func (o *RecommendationProperties) GetNotificationExpirationTimeAsTime() (*time.Time, error) {
	if o.NotificationExpirationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.NotificationExpirationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RecommendationProperties) SetNotificationExpirationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.NotificationExpirationTime = &formatted
}

func (o *RecommendationProperties) GetNotifiedTimeAsTime() (*time.Time, error) {
	if o.NotifiedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.NotifiedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RecommendationProperties) SetNotifiedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.NotifiedTime = &formatted
}

func (o *RecommendationProperties) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RecommendationProperties) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}
