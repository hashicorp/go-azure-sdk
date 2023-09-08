package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobStatusDetails string

const (
	PrintJobStatusDetailscompletedSuccessfully PrintJobStatusDetails = "CompletedSuccessfully"
	PrintJobStatusDetailscompletedWithErrors   PrintJobStatusDetails = "CompletedWithErrors"
	PrintJobStatusDetailscompletedWithWarnings PrintJobStatusDetails = "CompletedWithWarnings"
	PrintJobStatusDetailsinterpreting          PrintJobStatusDetails = "Interpreting"
	PrintJobStatusDetailsreleaseWait           PrintJobStatusDetails = "ReleaseWait"
	PrintJobStatusDetailstransforming          PrintJobStatusDetails = "Transforming"
	PrintJobStatusDetailsuploadPending         PrintJobStatusDetails = "UploadPending"
)

func PossibleValuesForPrintJobStatusDetails() []string {
	return []string{
		string(PrintJobStatusDetailscompletedSuccessfully),
		string(PrintJobStatusDetailscompletedWithErrors),
		string(PrintJobStatusDetailscompletedWithWarnings),
		string(PrintJobStatusDetailsinterpreting),
		string(PrintJobStatusDetailsreleaseWait),
		string(PrintJobStatusDetailstransforming),
		string(PrintJobStatusDetailsuploadPending),
	}
}

func parsePrintJobStatusDetails(input string) (*PrintJobStatusDetails, error) {
	vals := map[string]PrintJobStatusDetails{
		"completedsuccessfully": PrintJobStatusDetailscompletedSuccessfully,
		"completedwitherrors":   PrintJobStatusDetailscompletedWithErrors,
		"completedwithwarnings": PrintJobStatusDetailscompletedWithWarnings,
		"interpreting":          PrintJobStatusDetailsinterpreting,
		"releasewait":           PrintJobStatusDetailsreleaseWait,
		"transforming":          PrintJobStatusDetailstransforming,
		"uploadpending":         PrintJobStatusDetailsuploadPending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobStatusDetails(input)
	return &out, nil
}
