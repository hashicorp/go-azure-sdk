package recoverypoint

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RehydrationStatus string

const (
	RehydrationStatusCOMPLETED        RehydrationStatus = "COMPLETED"
	RehydrationStatusCREATEINPROGRESS RehydrationStatus = "CREATE_IN_PROGRESS"
	RehydrationStatusDELETED          RehydrationStatus = "DELETED"
	RehydrationStatusDELETEINPROGRESS RehydrationStatus = "DELETE_IN_PROGRESS"
	RehydrationStatusFAILED           RehydrationStatus = "FAILED"
)

func PossibleValuesForRehydrationStatus() []string {
	return []string{
		string(RehydrationStatusCOMPLETED),
		string(RehydrationStatusCREATEINPROGRESS),
		string(RehydrationStatusDELETED),
		string(RehydrationStatusDELETEINPROGRESS),
		string(RehydrationStatusFAILED),
	}
}
