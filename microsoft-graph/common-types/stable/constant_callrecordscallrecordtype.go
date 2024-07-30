package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsCallRecordType string

const (
	CallRecordsCallRecordType_GroupCall  CallRecordsCallRecordType = "groupCall"
	CallRecordsCallRecordType_PeerToPeer CallRecordsCallRecordType = "peerToPeer"
	CallRecordsCallRecordType_Unknown    CallRecordsCallRecordType = "unknown"
)

func PossibleValuesForCallRecordsCallRecordType() []string {
	return []string{
		string(CallRecordsCallRecordType_GroupCall),
		string(CallRecordsCallRecordType_PeerToPeer),
		string(CallRecordsCallRecordType_Unknown),
	}
}

func (s *CallRecordsCallRecordType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsCallRecordType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsCallRecordType(input string) (*CallRecordsCallRecordType, error) {
	vals := map[string]CallRecordsCallRecordType{
		"groupcall":  CallRecordsCallRecordType_GroupCall,
		"peertopeer": CallRecordsCallRecordType_PeerToPeer,
		"unknown":    CallRecordsCallRecordType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsCallRecordType(input)
	return &out, nil
}
