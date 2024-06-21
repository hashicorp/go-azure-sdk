package integrationruntime

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedIntegrationRuntimeOperationResult struct {
	ActivityId *string   `json:"activityId,omitempty"`
	ErrorCode  *string   `json:"errorCode,omitempty"`
	Parameters *[]string `json:"parameters,omitempty"`
	Result     *string   `json:"result,omitempty"`
	StartTime  *string   `json:"startTime,omitempty"`
	Type       *string   `json:"type,omitempty"`
}
