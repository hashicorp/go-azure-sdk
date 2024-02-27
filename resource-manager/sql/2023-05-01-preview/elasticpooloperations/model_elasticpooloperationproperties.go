package elasticpooloperations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticPoolOperationProperties struct {
	Description             *string `json:"description,omitempty"`
	ElasticPoolName         *string `json:"elasticPoolName,omitempty"`
	ErrorCode               *int64  `json:"errorCode,omitempty"`
	ErrorDescription        *string `json:"errorDescription,omitempty"`
	ErrorSeverity           *int64  `json:"errorSeverity,omitempty"`
	EstimatedCompletionTime *string `json:"estimatedCompletionTime,omitempty"`
	IsCancellable           *bool   `json:"isCancellable,omitempty"`
	IsUserError             *bool   `json:"isUserError,omitempty"`
	Operation               *string `json:"operation,omitempty"`
	OperationFriendlyName   *string `json:"operationFriendlyName,omitempty"`
	PercentComplete         *int64  `json:"percentComplete,omitempty"`
	ServerName              *string `json:"serverName,omitempty"`
	StartTime               *string `json:"startTime,omitempty"`
	State                   *string `json:"state,omitempty"`
}

func (o *ElasticPoolOperationProperties) GetEstimatedCompletionTimeAsTime() (*time.Time, error) {
	if o.EstimatedCompletionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EstimatedCompletionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ElasticPoolOperationProperties) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}
