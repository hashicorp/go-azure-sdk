package replicationjobs

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobErrorDetails struct {
	CreationTime         *string        `json:"creationTime,omitempty"`
	ErrorLevel           *string        `json:"errorLevel,omitempty"`
	ProviderErrorDetails *ProviderError `json:"providerErrorDetails"`
	ServiceErrorDetails  *ServiceError  `json:"serviceErrorDetails"`
	TaskId               *string        `json:"taskId,omitempty"`
}

func (o *JobErrorDetails) GetCreationTimeAsTime() (*time.Time, error) {
	if o.CreationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *JobErrorDetails) SetCreationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationTime = &formatted
}
