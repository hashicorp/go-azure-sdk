package organizations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfferDetails struct {
	OfferId     string  `json:"offerId"`
	PlanId      string  `json:"planId"`
	PlanName    string  `json:"planName"`
	PublisherId string  `json:"publisherId"`
	TermId      string  `json:"termId"`
	TermUnit    *string `json:"termUnit,omitempty"`
}
