package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResponseStatusResponse string

const (
	ResponseStatusResponseaccepted            ResponseStatusResponse = "Accepted"
	ResponseStatusResponsedeclined            ResponseStatusResponse = "Declined"
	ResponseStatusResponsenone                ResponseStatusResponse = "None"
	ResponseStatusResponsenotResponded        ResponseStatusResponse = "NotResponded"
	ResponseStatusResponseorganizer           ResponseStatusResponse = "Organizer"
	ResponseStatusResponsetentativelyAccepted ResponseStatusResponse = "TentativelyAccepted"
)

func PossibleValuesForResponseStatusResponse() []string {
	return []string{
		string(ResponseStatusResponseaccepted),
		string(ResponseStatusResponsedeclined),
		string(ResponseStatusResponsenone),
		string(ResponseStatusResponsenotResponded),
		string(ResponseStatusResponseorganizer),
		string(ResponseStatusResponsetentativelyAccepted),
	}
}

func parseResponseStatusResponse(input string) (*ResponseStatusResponse, error) {
	vals := map[string]ResponseStatusResponse{
		"accepted":            ResponseStatusResponseaccepted,
		"declined":            ResponseStatusResponsedeclined,
		"none":                ResponseStatusResponsenone,
		"notresponded":        ResponseStatusResponsenotResponded,
		"organizer":           ResponseStatusResponseorganizer,
		"tentativelyaccepted": ResponseStatusResponsetentativelyAccepted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResponseStatusResponse(input)
	return &out, nil
}
