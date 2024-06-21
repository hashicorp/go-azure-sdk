package policyrestrictionsvalidations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationResultContractProperties struct {
	ActionLog  *[]OperationResultLogItemContract `json:"actionLog,omitempty"`
	Error      *ErrorResponseBody                `json:"error,omitempty"`
	Id         *string                           `json:"id,omitempty"`
	ResultInfo *string                           `json:"resultInfo,omitempty"`
	Started    *string                           `json:"started,omitempty"`
	Status     *AsyncOperationStatus             `json:"status,omitempty"`
	Updated    *string                           `json:"updated,omitempty"`
}
