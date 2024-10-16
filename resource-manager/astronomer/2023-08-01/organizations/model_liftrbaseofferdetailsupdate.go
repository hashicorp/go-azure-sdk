package organizations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LiftrBaseOfferDetailsUpdate struct {
	OfferId     *string `json:"offerId,omitempty"`
	PlanId      *string `json:"planId,omitempty"`
	PlanName    *string `json:"planName,omitempty"`
	PublisherId *string `json:"publisherId,omitempty"`
	TermId      *string `json:"termId,omitempty"`
	TermUnit    *string `json:"termUnit,omitempty"`
}
