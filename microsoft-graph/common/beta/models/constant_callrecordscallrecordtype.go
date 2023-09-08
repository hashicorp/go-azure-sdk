package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsCallRecordType string

const (
	CallRecordsCallRecordTypegroupCall  CallRecordsCallRecordType = "GroupCall"
	CallRecordsCallRecordTypepeerToPeer CallRecordsCallRecordType = "PeerToPeer"
	CallRecordsCallRecordTypeunknown    CallRecordsCallRecordType = "Unknown"
)

func PossibleValuesForCallRecordsCallRecordType() []string {
	return []string{
		string(CallRecordsCallRecordTypegroupCall),
		string(CallRecordsCallRecordTypepeerToPeer),
		string(CallRecordsCallRecordTypeunknown),
	}
}

func parseCallRecordsCallRecordType(input string) (*CallRecordsCallRecordType, error) {
	vals := map[string]CallRecordsCallRecordType{
		"groupcall":  CallRecordsCallRecordTypegroupCall,
		"peertopeer": CallRecordsCallRecordTypepeerToPeer,
		"unknown":    CallRecordsCallRecordTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsCallRecordType(input)
	return &out, nil
}
