package agreements

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OldAgreementProperties struct {
	CancelDate *string `json:"cancelDate,omitempty"`
	Id         *string `json:"id,omitempty"`
	Offer      *string `json:"offer,omitempty"`
	Publisher  *string `json:"publisher,omitempty"`
	SignDate   *string `json:"signDate,omitempty"`
	State      *State  `json:"state,omitempty"`
}
