package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RejectJoinResponseReason string

const (
	RejectJoinResponseReasonbusy      RejectJoinResponseReason = "Busy"
	RejectJoinResponseReasonforbidden RejectJoinResponseReason = "Forbidden"
	RejectJoinResponseReasonnone      RejectJoinResponseReason = "None"
)

func PossibleValuesForRejectJoinResponseReason() []string {
	return []string{
		string(RejectJoinResponseReasonbusy),
		string(RejectJoinResponseReasonforbidden),
		string(RejectJoinResponseReasonnone),
	}
}

func parseRejectJoinResponseReason(input string) (*RejectJoinResponseReason, error) {
	vals := map[string]RejectJoinResponseReason{
		"busy":      RejectJoinResponseReasonbusy,
		"forbidden": RejectJoinResponseReasonforbidden,
		"none":      RejectJoinResponseReasonnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RejectJoinResponseReason(input)
	return &out, nil
}
