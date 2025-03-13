package sourcecontrols

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Deployment struct {
	DeploymentId      *string           `json:"deploymentId,omitempty"`
	DeploymentLogsURL *string           `json:"deploymentLogsUrl,omitempty"`
	DeploymentResult  *DeploymentResult `json:"deploymentResult,omitempty"`
	DeploymentState   *DeploymentState  `json:"deploymentState,omitempty"`
	DeploymentTime    *string           `json:"deploymentTime,omitempty"`
}

func (o *Deployment) GetDeploymentTimeAsTime() (*time.Time, error) {
	if o.DeploymentTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DeploymentTime, "2006-01-02T15:04:05Z07:00")
}

func (o *Deployment) SetDeploymentTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.DeploymentTime = &formatted
}
