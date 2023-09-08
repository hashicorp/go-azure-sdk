package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StatusDetailsStatus string

const (
	StatusDetailsStatusfailure StatusDetailsStatus = "Failure"
	StatusDetailsStatusskipped StatusDetailsStatus = "Skipped"
	StatusDetailsStatussuccess StatusDetailsStatus = "Success"
	StatusDetailsStatuswarning StatusDetailsStatus = "Warning"
)

func PossibleValuesForStatusDetailsStatus() []string {
	return []string{
		string(StatusDetailsStatusfailure),
		string(StatusDetailsStatusskipped),
		string(StatusDetailsStatussuccess),
		string(StatusDetailsStatuswarning),
	}
}

func parseStatusDetailsStatus(input string) (*StatusDetailsStatus, error) {
	vals := map[string]StatusDetailsStatus{
		"failure": StatusDetailsStatusfailure,
		"skipped": StatusDetailsStatusskipped,
		"success": StatusDetailsStatussuccess,
		"warning": StatusDetailsStatuswarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusDetailsStatus(input)
	return &out, nil
}
