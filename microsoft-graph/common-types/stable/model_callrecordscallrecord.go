package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsCallRecord struct {
	EndDateTime          *string                          `json:"endDateTime,omitempty"`
	Id                   *string                          `json:"id,omitempty"`
	JoinWebUrl           *string                          `json:"joinWebUrl,omitempty"`
	LastModifiedDateTime *string                          `json:"lastModifiedDateTime,omitempty"`
	Modalities           *CallRecordsCallRecordModalities `json:"modalities,omitempty"`
	ODataType            *string                          `json:"@odata.type,omitempty"`
	Organizer            *IdentitySet                     `json:"organizer,omitempty"`
	Participants         *[]IdentitySet                   `json:"participants,omitempty"`
	Sessions             *[]CallRecordsSession            `json:"sessions,omitempty"`
	StartDateTime        *string                          `json:"startDateTime,omitempty"`
	Type                 *CallRecordsCallRecordType       `json:"type,omitempty"`
	Version              *int64                           `json:"version,omitempty"`
}
