package managedinstanceoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceOperationSteps struct {
	CurrentStep *int64                              `json:"currentStep,omitempty"`
	StepsList   *[]UpsertManagedServerOperationStep `json:"stepsList,omitempty"`
	TotalSteps  *string                             `json:"totalSteps,omitempty"`
}
