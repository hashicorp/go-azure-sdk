package deploymentoperations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentOperationProperties struct {
	Duration              *string                `json:"duration,omitempty"`
	ProvisioningOperation *ProvisioningOperation `json:"provisioningOperation,omitempty"`
	ProvisioningState     *string                `json:"provisioningState,omitempty"`
	Request               *HTTPMessage           `json:"request,omitempty"`
	Response              *HTTPMessage           `json:"response,omitempty"`
	ServiceRequestId      *string                `json:"serviceRequestId,omitempty"`
	StatusCode            *string                `json:"statusCode,omitempty"`
	StatusMessage         *StatusMessage         `json:"statusMessage,omitempty"`
	TargetResource        *TargetResource        `json:"targetResource,omitempty"`
	Timestamp             *string                `json:"timestamp,omitempty"`
}

func (o *DeploymentOperationProperties) GetTimestampAsTime() (*time.Time, error) {
	if o.Timestamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Timestamp, "2006-01-02T15:04:05Z07:00")
}

func (o *DeploymentOperationProperties) SetTimestampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Timestamp = &formatted
}
