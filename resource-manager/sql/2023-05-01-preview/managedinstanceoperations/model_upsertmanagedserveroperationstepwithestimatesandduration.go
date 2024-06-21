package managedinstanceoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpsertManagedServerOperationStepWithEstimatesAndDuration struct {
	Name          *string `json:"name,omitempty"`
	Order         *int64  `json:"order,omitempty"`
	Status        *Status `json:"status,omitempty"`
	StepEndTime   *string `json:"stepEndTime,omitempty"`
	StepStartTime *string `json:"stepStartTime,omitempty"`
	TimeElapsed   *string `json:"timeElapsed,omitempty"`
}
