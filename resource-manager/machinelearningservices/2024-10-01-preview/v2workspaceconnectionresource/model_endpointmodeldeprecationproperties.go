package v2workspaceconnectionresource

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointModelDeprecationProperties struct {
	FineTune  *string `json:"fineTune,omitempty"`
	Inference *string `json:"inference,omitempty"`
}

func (o *EndpointModelDeprecationProperties) GetFineTuneAsTime() (*time.Time, error) {
	if o.FineTune == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.FineTune, "2006-01-02T15:04:05Z07:00")
}

func (o *EndpointModelDeprecationProperties) SetFineTuneAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.FineTune = &formatted
}

func (o *EndpointModelDeprecationProperties) GetInferenceAsTime() (*time.Time, error) {
	if o.Inference == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Inference, "2006-01-02T15:04:05Z07:00")
}

func (o *EndpointModelDeprecationProperties) SetInferenceAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Inference = &formatted
}
