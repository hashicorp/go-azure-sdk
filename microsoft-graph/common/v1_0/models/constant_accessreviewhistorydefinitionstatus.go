package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewHistoryDefinitionStatus string

const (
	AccessReviewHistoryDefinitionStatusdone       AccessReviewHistoryDefinitionStatus = "Done"
	AccessReviewHistoryDefinitionStatuserror      AccessReviewHistoryDefinitionStatus = "Error"
	AccessReviewHistoryDefinitionStatusinprogress AccessReviewHistoryDefinitionStatus = "Inprogress"
	AccessReviewHistoryDefinitionStatusrequested  AccessReviewHistoryDefinitionStatus = "Requested"
)

func PossibleValuesForAccessReviewHistoryDefinitionStatus() []string {
	return []string{
		string(AccessReviewHistoryDefinitionStatusdone),
		string(AccessReviewHistoryDefinitionStatuserror),
		string(AccessReviewHistoryDefinitionStatusinprogress),
		string(AccessReviewHistoryDefinitionStatusrequested),
	}
}

func parseAccessReviewHistoryDefinitionStatus(input string) (*AccessReviewHistoryDefinitionStatus, error) {
	vals := map[string]AccessReviewHistoryDefinitionStatus{
		"done":       AccessReviewHistoryDefinitionStatusdone,
		"error":      AccessReviewHistoryDefinitionStatuserror,
		"inprogress": AccessReviewHistoryDefinitionStatusinprogress,
		"requested":  AccessReviewHistoryDefinitionStatusrequested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessReviewHistoryDefinitionStatus(input)
	return &out, nil
}
