package containerappssourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceControlOperationState string

const (
	SourceControlOperationStateCanceled   SourceControlOperationState = "Canceled"
	SourceControlOperationStateFailed     SourceControlOperationState = "Failed"
	SourceControlOperationStateInProgress SourceControlOperationState = "InProgress"
	SourceControlOperationStateSucceeded  SourceControlOperationState = "Succeeded"
)

func PossibleValuesForSourceControlOperationState() []string {
	return []string{
		string(SourceControlOperationStateCanceled),
		string(SourceControlOperationStateFailed),
		string(SourceControlOperationStateInProgress),
		string(SourceControlOperationStateSucceeded),
	}
}
