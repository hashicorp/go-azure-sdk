package namespacedevices

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StatusConfig struct {
	Error              *StatusError `json:"error,omitempty"`
	LastTransitionTime *string      `json:"lastTransitionTime,omitempty"`
	Version            *int64       `json:"version,omitempty"`
}

func (o *StatusConfig) GetLastTransitionTimeAsTime() (*time.Time, error) {
	if o.LastTransitionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastTransitionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *StatusConfig) SetLastTransitionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastTransitionTime = &formatted
}
