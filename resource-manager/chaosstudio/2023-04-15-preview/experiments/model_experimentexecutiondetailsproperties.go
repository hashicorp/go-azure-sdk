package experiments

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExperimentExecutionDetailsProperties struct {
	CreatedDateTime    *string                                             `json:"createdDateTime,omitempty"`
	ExperimentId       *string                                             `json:"experimentId,omitempty"`
	FailureReason      *string                                             `json:"failureReason,omitempty"`
	LastActionDateTime *string                                             `json:"lastActionDateTime,omitempty"`
	RunInformation     *ExperimentExecutionDetailsPropertiesRunInformation `json:"runInformation,omitempty"`
	StartDateTime      *string                                             `json:"startDateTime,omitempty"`
	Status             *string                                             `json:"status,omitempty"`
	StopDateTime       *string                                             `json:"stopDateTime,omitempty"`
}

func (o *ExperimentExecutionDetailsProperties) GetCreatedDateTimeAsTime() (*time.Time, error) {
	if o.CreatedDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ExperimentExecutionDetailsProperties) GetLastActionDateTimeAsTime() (*time.Time, error) {
	if o.LastActionDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastActionDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ExperimentExecutionDetailsProperties) GetStartDateTimeAsTime() (*time.Time, error) {
	if o.StartDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ExperimentExecutionDetailsProperties) GetStopDateTimeAsTime() (*time.Time, error) {
	if o.StopDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StopDateTime, "2006-01-02T15:04:05Z07:00")
}
