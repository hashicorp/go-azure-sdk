package recordsets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecordType string

const (
	RecordTypeA     RecordType = "A"
	RecordTypeAAAA  RecordType = "AAAA"
	RecordTypeCNAME RecordType = "CNAME"
	RecordTypeMX    RecordType = "MX"
	RecordTypePTR   RecordType = "PTR"
	RecordTypeSOA   RecordType = "SOA"
	RecordTypeSRV   RecordType = "SRV"
	RecordTypeTXT   RecordType = "TXT"
)

func PossibleValuesForRecordType() []string {
	return []string{
		string(RecordTypeA),
		string(RecordTypeAAAA),
		string(RecordTypeCNAME),
		string(RecordTypeMX),
		string(RecordTypePTR),
		string(RecordTypeSOA),
		string(RecordTypeSRV),
		string(RecordTypeTXT),
	}
}
