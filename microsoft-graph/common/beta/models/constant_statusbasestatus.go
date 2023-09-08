package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StatusBaseStatus string

const (
	StatusBaseStatusfailure StatusBaseStatus = "Failure"
	StatusBaseStatusskipped StatusBaseStatus = "Skipped"
	StatusBaseStatussuccess StatusBaseStatus = "Success"
	StatusBaseStatuswarning StatusBaseStatus = "Warning"
)

func PossibleValuesForStatusBaseStatus() []string {
	return []string{
		string(StatusBaseStatusfailure),
		string(StatusBaseStatusskipped),
		string(StatusBaseStatussuccess),
		string(StatusBaseStatuswarning),
	}
}

func parseStatusBaseStatus(input string) (*StatusBaseStatus, error) {
	vals := map[string]StatusBaseStatus{
		"failure": StatusBaseStatusfailure,
		"skipped": StatusBaseStatusskipped,
		"success": StatusBaseStatussuccess,
		"warning": StatusBaseStatuswarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusBaseStatus(input)
	return &out, nil
}
