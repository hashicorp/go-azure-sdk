package certificates

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Certificate struct {
	DeleteCertificateError      *DeleteCertificateError `json:"deleteCertificateError,omitempty"`
	PreviousState               *CertificateState       `json:"previousState,omitempty"`
	PreviousStateTransitionTime *string                 `json:"previousStateTransitionTime,omitempty"`
	PublicData                  *string                 `json:"publicData,omitempty"`
	State                       *CertificateState       `json:"state,omitempty"`
	StateTransitionTime         *string                 `json:"stateTransitionTime,omitempty"`
	Thumbprint                  *string                 `json:"thumbprint,omitempty"`
	ThumbprintAlgorithm         *string                 `json:"thumbprintAlgorithm,omitempty"`
	Url                         *string                 `json:"url,omitempty"`
}

func (o *Certificate) GetPreviousStateTransitionTimeAsTime() (*time.Time, error) {
	if o.PreviousStateTransitionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.PreviousStateTransitionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *Certificate) SetPreviousStateTransitionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.PreviousStateTransitionTime = &formatted
}

func (o *Certificate) GetStateTransitionTimeAsTime() (*time.Time, error) {
	if o.StateTransitionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StateTransitionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *Certificate) SetStateTransitionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StateTransitionTime = &formatted
}
