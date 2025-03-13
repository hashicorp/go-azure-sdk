package bookmark

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailMessageEntityProperties struct {
	AdditionalData         *interface{}           `json:"additionalData,omitempty"`
	AntispamDirection      *AntispamMailDirection `json:"antispamDirection,omitempty"`
	BodyFingerprintBin1    *int64                 `json:"bodyFingerprintBin1,omitempty"`
	BodyFingerprintBin2    *int64                 `json:"bodyFingerprintBin2,omitempty"`
	BodyFingerprintBin3    *int64                 `json:"bodyFingerprintBin3,omitempty"`
	BodyFingerprintBin4    *int64                 `json:"bodyFingerprintBin4,omitempty"`
	BodyFingerprintBin5    *int64                 `json:"bodyFingerprintBin5,omitempty"`
	DeliveryAction         *DeliveryAction        `json:"deliveryAction,omitempty"`
	DeliveryLocation       *DeliveryLocation      `json:"deliveryLocation,omitempty"`
	FileEntityIds          *[]string              `json:"fileEntityIds,omitempty"`
	FriendlyName           *string                `json:"friendlyName,omitempty"`
	InternetMessageId      *string                `json:"internetMessageId,omitempty"`
	Language               *string                `json:"language,omitempty"`
	NetworkMessageId       *string                `json:"networkMessageId,omitempty"`
	P1Sender               *string                `json:"p1Sender,omitempty"`
	P1SenderDisplayName    *string                `json:"p1SenderDisplayName,omitempty"`
	P1SenderDomain         *string                `json:"p1SenderDomain,omitempty"`
	P2Sender               *string                `json:"p2Sender,omitempty"`
	P2SenderDisplayName    *string                `json:"p2SenderDisplayName,omitempty"`
	P2SenderDomain         *string                `json:"p2SenderDomain,omitempty"`
	ReceiveDate            *string                `json:"receiveDate,omitempty"`
	Recipient              *string                `json:"recipient,omitempty"`
	SenderIP               *string                `json:"senderIP,omitempty"`
	Subject                *string                `json:"subject,omitempty"`
	ThreatDetectionMethods *[]string              `json:"threatDetectionMethods,omitempty"`
	Threats                *[]string              `json:"threats,omitempty"`
	Urls                   *[]string              `json:"urls,omitempty"`
}

func (o *MailMessageEntityProperties) GetReceiveDateAsTime() (*time.Time, error) {
	if o.ReceiveDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ReceiveDate, "2006-01-02T15:04:05Z07:00")
}

func (o *MailMessageEntityProperties) SetReceiveDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ReceiveDate = &formatted
}
