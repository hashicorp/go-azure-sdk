package workflowversions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowVersionProperties struct {
	AccessControl          *FlowAccessControlConfiguration `json:"accessControl,omitempty"`
	AccessEndpoint         *string                         `json:"accessEndpoint,omitempty"`
	ChangedTime            *string                         `json:"changedTime,omitempty"`
	CreatedTime            *string                         `json:"createdTime,omitempty"`
	Definition             *interface{}                    `json:"definition,omitempty"`
	EndpointsConfiguration *FlowEndpointsConfiguration     `json:"endpointsConfiguration,omitempty"`
	IntegrationAccount     *ResourceReference              `json:"integrationAccount,omitempty"`
	Parameters             *map[string]WorkflowParameter   `json:"parameters,omitempty"`
	ProvisioningState      *WorkflowProvisioningState      `json:"provisioningState,omitempty"`
	Sku                    *WorkflowSku                    `json:"sku,omitempty"`
	State                  *WorkflowState                  `json:"state,omitempty"`
	Version                *string                         `json:"version,omitempty"`
}

func (o *WorkflowVersionProperties) GetChangedTimeAsTime() (*time.Time, error) {
	if o.ChangedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ChangedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *WorkflowVersionProperties) GetCreatedTimeAsTime() (*time.Time, error) {
	if o.CreatedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedTime, "2006-01-02T15:04:05Z07:00")
}
