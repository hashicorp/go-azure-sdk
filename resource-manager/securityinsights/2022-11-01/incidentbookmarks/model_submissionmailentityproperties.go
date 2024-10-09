package incidentbookmarks

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubmissionMailEntityProperties struct {
	AdditionalData   *map[string]interface{} `json:"additionalData,omitempty"`
	FriendlyName     *string                 `json:"friendlyName,omitempty"`
	NetworkMessageId *string                 `json:"networkMessageId,omitempty"`
	Recipient        *string                 `json:"recipient,omitempty"`
	ReportType       *string                 `json:"reportType,omitempty"`
	Sender           *string                 `json:"sender,omitempty"`
	SenderIP         *string                 `json:"senderIp,omitempty"`
	Subject          *string                 `json:"subject,omitempty"`
	SubmissionDate   *string                 `json:"submissionDate,omitempty"`
	SubmissionId     *string                 `json:"submissionId,omitempty"`
	Submitter        *string                 `json:"submitter,omitempty"`
	Timestamp        *string                 `json:"timestamp,omitempty"`
}

func (o *SubmissionMailEntityProperties) GetSubmissionDateAsTime() (*time.Time, error) {
	if o.SubmissionDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.SubmissionDate, "2006-01-02T15:04:05Z07:00")
}

func (o *SubmissionMailEntityProperties) SetSubmissionDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.SubmissionDate = &formatted
}

func (o *SubmissionMailEntityProperties) GetTimestampAsTime() (*time.Time, error) {
	if o.Timestamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Timestamp, "2006-01-02T15:04:05Z07:00")
}

func (o *SubmissionMailEntityProperties) SetTimestampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Timestamp = &formatted
}
