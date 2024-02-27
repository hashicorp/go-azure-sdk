package manageddatabasemoveoperations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDatabaseMoveOperationResultProperties struct {
	ErrorCode                 *int64                    `json:"errorCode,omitempty"`
	ErrorDescription          *string                   `json:"errorDescription,omitempty"`
	ErrorSeverity             *int64                    `json:"errorSeverity,omitempty"`
	IsCancellable             *bool                     `json:"isCancellable,omitempty"`
	IsUserError               *bool                     `json:"isUserError,omitempty"`
	Operation                 *string                   `json:"operation,omitempty"`
	OperationFriendlyName     *string                   `json:"operationFriendlyName,omitempty"`
	OperationMode             *MoveOperationMode        `json:"operationMode,omitempty"`
	SourceDatabaseName        *string                   `json:"sourceDatabaseName,omitempty"`
	SourceManagedInstanceId   *string                   `json:"sourceManagedInstanceId,omitempty"`
	SourceManagedInstanceName *string                   `json:"sourceManagedInstanceName,omitempty"`
	StartTime                 *string                   `json:"startTime,omitempty"`
	State                     *ManagementOperationState `json:"state,omitempty"`
	TargetDatabaseName        *string                   `json:"targetDatabaseName,omitempty"`
	TargetManagedInstanceId   *string                   `json:"targetManagedInstanceId,omitempty"`
	TargetManagedInstanceName *string                   `json:"targetManagedInstanceName,omitempty"`
}

func (o *ManagedDatabaseMoveOperationResultProperties) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}
