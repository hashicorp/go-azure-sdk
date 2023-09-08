package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewHistoryInstanceStatus string

const (
	AccessReviewHistoryInstanceStatusdone       AccessReviewHistoryInstanceStatus = "Done"
	AccessReviewHistoryInstanceStatuserror      AccessReviewHistoryInstanceStatus = "Error"
	AccessReviewHistoryInstanceStatusinprogress AccessReviewHistoryInstanceStatus = "Inprogress"
	AccessReviewHistoryInstanceStatusrequested  AccessReviewHistoryInstanceStatus = "Requested"
)

func PossibleValuesForAccessReviewHistoryInstanceStatus() []string {
	return []string{
		string(AccessReviewHistoryInstanceStatusdone),
		string(AccessReviewHistoryInstanceStatuserror),
		string(AccessReviewHistoryInstanceStatusinprogress),
		string(AccessReviewHistoryInstanceStatusrequested),
	}
}

func parseAccessReviewHistoryInstanceStatus(input string) (*AccessReviewHistoryInstanceStatus, error) {
	vals := map[string]AccessReviewHistoryInstanceStatus{
		"done":       AccessReviewHistoryInstanceStatusdone,
		"error":      AccessReviewHistoryInstanceStatuserror,
		"inprogress": AccessReviewHistoryInstanceStatusinprogress,
		"requested":  AccessReviewHistoryInstanceStatusrequested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessReviewHistoryInstanceStatus(input)
	return &out, nil
}
