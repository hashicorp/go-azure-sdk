package featuresetversion

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeaturesetJob struct {
	CreatedDate   *string              `json:"createdDate,omitempty"`
	DisplayName   *string              `json:"displayName,omitempty"`
	Duration      *string              `json:"duration,omitempty"`
	ExperimentId  *string              `json:"experimentId,omitempty"`
	FeatureWindow *FeatureWindow       `json:"featureWindow,omitempty"`
	JobId         *string              `json:"jobId,omitempty"`
	Status        *JobStatus           `json:"status,omitempty"`
	Tags          *map[string]string   `json:"tags,omitempty"`
	Type          *FeaturestoreJobType `json:"type,omitempty"`
}

func (o *FeaturesetJob) GetCreatedDateAsTime() (*time.Time, error) {
	if o.CreatedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedDate, "2006-01-02T15:04:05Z07:00")
}

func (o *FeaturesetJob) SetCreatedDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedDate = &formatted
}
