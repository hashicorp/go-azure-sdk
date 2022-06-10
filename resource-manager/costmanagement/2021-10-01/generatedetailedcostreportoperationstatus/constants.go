package generatedetailedcostreportoperationstatus

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportOperationStatusType string

const (
	ReportOperationStatusTypeCompleted       ReportOperationStatusType = "Completed"
	ReportOperationStatusTypeFailed          ReportOperationStatusType = "Failed"
	ReportOperationStatusTypeInProgress      ReportOperationStatusType = "InProgress"
	ReportOperationStatusTypeNoDataFound     ReportOperationStatusType = "NoDataFound"
	ReportOperationStatusTypeQueued          ReportOperationStatusType = "Queued"
	ReportOperationStatusTypeReadyToDownload ReportOperationStatusType = "ReadyToDownload"
	ReportOperationStatusTypeTimedOut        ReportOperationStatusType = "TimedOut"
)

func PossibleValuesForReportOperationStatusType() []string {
	return []string{
		string(ReportOperationStatusTypeCompleted),
		string(ReportOperationStatusTypeFailed),
		string(ReportOperationStatusTypeInProgress),
		string(ReportOperationStatusTypeNoDataFound),
		string(ReportOperationStatusTypeQueued),
		string(ReportOperationStatusTypeReadyToDownload),
		string(ReportOperationStatusTypeTimedOut),
	}
}

func parseReportOperationStatusType(input string) (*ReportOperationStatusType, error) {
	vals := map[string]ReportOperationStatusType{
		"completed":       ReportOperationStatusTypeCompleted,
		"failed":          ReportOperationStatusTypeFailed,
		"inprogress":      ReportOperationStatusTypeInProgress,
		"nodatafound":     ReportOperationStatusTypeNoDataFound,
		"queued":          ReportOperationStatusTypeQueued,
		"readytodownload": ReportOperationStatusTypeReadyToDownload,
		"timedout":        ReportOperationStatusTypeTimedOut,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReportOperationStatusType(input)
	return &out, nil
}
