package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobStatusDetails string

const (
	PrintJobStatusDetails_CompletedSuccessfully PrintJobStatusDetails = "completedSuccessfully"
	PrintJobStatusDetails_CompletedWithErrors   PrintJobStatusDetails = "completedWithErrors"
	PrintJobStatusDetails_CompletedWithWarnings PrintJobStatusDetails = "completedWithWarnings"
	PrintJobStatusDetails_Interpreting          PrintJobStatusDetails = "interpreting"
	PrintJobStatusDetails_ReleaseWait           PrintJobStatusDetails = "releaseWait"
	PrintJobStatusDetails_Transforming          PrintJobStatusDetails = "transforming"
	PrintJobStatusDetails_UploadPending         PrintJobStatusDetails = "uploadPending"
)

func PossibleValuesForPrintJobStatusDetails() []string {
	return []string{
		string(PrintJobStatusDetails_CompletedSuccessfully),
		string(PrintJobStatusDetails_CompletedWithErrors),
		string(PrintJobStatusDetails_CompletedWithWarnings),
		string(PrintJobStatusDetails_Interpreting),
		string(PrintJobStatusDetails_ReleaseWait),
		string(PrintJobStatusDetails_Transforming),
		string(PrintJobStatusDetails_UploadPending),
	}
}

func (s *PrintJobStatusDetails) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobStatusDetails(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobStatusDetails(input string) (*PrintJobStatusDetails, error) {
	vals := map[string]PrintJobStatusDetails{
		"completedsuccessfully": PrintJobStatusDetails_CompletedSuccessfully,
		"completedwitherrors":   PrintJobStatusDetails_CompletedWithErrors,
		"completedwithwarnings": PrintJobStatusDetails_CompletedWithWarnings,
		"interpreting":          PrintJobStatusDetails_Interpreting,
		"releasewait":           PrintJobStatusDetails_ReleaseWait,
		"transforming":          PrintJobStatusDetails_Transforming,
		"uploadpending":         PrintJobStatusDetails_UploadPending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobStatusDetails(input)
	return &out, nil
}
