package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResponseStatusResponse string

const (
	ResponseStatusResponse_Accepted            ResponseStatusResponse = "accepted"
	ResponseStatusResponse_Declined            ResponseStatusResponse = "declined"
	ResponseStatusResponse_None                ResponseStatusResponse = "none"
	ResponseStatusResponse_NotResponded        ResponseStatusResponse = "notResponded"
	ResponseStatusResponse_Organizer           ResponseStatusResponse = "organizer"
	ResponseStatusResponse_TentativelyAccepted ResponseStatusResponse = "tentativelyAccepted"
)

func PossibleValuesForResponseStatusResponse() []string {
	return []string{
		string(ResponseStatusResponse_Accepted),
		string(ResponseStatusResponse_Declined),
		string(ResponseStatusResponse_None),
		string(ResponseStatusResponse_NotResponded),
		string(ResponseStatusResponse_Organizer),
		string(ResponseStatusResponse_TentativelyAccepted),
	}
}

func (s *ResponseStatusResponse) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResponseStatusResponse(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResponseStatusResponse(input string) (*ResponseStatusResponse, error) {
	vals := map[string]ResponseStatusResponse{
		"accepted":            ResponseStatusResponse_Accepted,
		"declined":            ResponseStatusResponse_Declined,
		"none":                ResponseStatusResponse_None,
		"notresponded":        ResponseStatusResponse_NotResponded,
		"organizer":           ResponseStatusResponse_Organizer,
		"tentativelyaccepted": ResponseStatusResponse_TentativelyAccepted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResponseStatusResponse(input)
	return &out, nil
}
