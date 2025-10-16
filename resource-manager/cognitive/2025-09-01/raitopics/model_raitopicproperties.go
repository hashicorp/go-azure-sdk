package raitopics

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RaiTopicProperties struct {
	CreatedAt      *string `json:"createdAt,omitempty"`
	Description    *string `json:"description,omitempty"`
	FailedReason   *string `json:"failedReason,omitempty"`
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`
	SampleBlobURL  *string `json:"sampleBlobUrl,omitempty"`
	Status         *string `json:"status,omitempty"`
	TopicId        *string `json:"topicId,omitempty"`
	TopicName      *string `json:"topicName,omitempty"`
}

func (o *RaiTopicProperties) GetCreatedAtAsTime() (*time.Time, error) {
	if o.CreatedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *RaiTopicProperties) SetCreatedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedAt = &formatted
}

func (o *RaiTopicProperties) GetLastModifiedAtAsTime() (*time.Time, error) {
	if o.LastModifiedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *RaiTopicProperties) SetLastModifiedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModifiedAt = &formatted
}
