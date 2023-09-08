package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsFailureInfoStage string

const (
	CallRecordsFailureInfoStagecallSetup CallRecordsFailureInfoStage = "CallSetup"
	CallRecordsFailureInfoStagemidcall   CallRecordsFailureInfoStage = "Midcall"
	CallRecordsFailureInfoStageunknown   CallRecordsFailureInfoStage = "Unknown"
)

func PossibleValuesForCallRecordsFailureInfoStage() []string {
	return []string{
		string(CallRecordsFailureInfoStagecallSetup),
		string(CallRecordsFailureInfoStagemidcall),
		string(CallRecordsFailureInfoStageunknown),
	}
}

func parseCallRecordsFailureInfoStage(input string) (*CallRecordsFailureInfoStage, error) {
	vals := map[string]CallRecordsFailureInfoStage{
		"callsetup": CallRecordsFailureInfoStagecallSetup,
		"midcall":   CallRecordsFailureInfoStagemidcall,
		"unknown":   CallRecordsFailureInfoStageunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsFailureInfoStage(input)
	return &out, nil
}
