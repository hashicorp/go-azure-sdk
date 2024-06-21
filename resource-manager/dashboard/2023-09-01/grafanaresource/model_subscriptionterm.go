package grafanaresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionTerm struct {
	EndDate   *string `json:"endDate,omitempty"`
	StartDate *string `json:"startDate,omitempty"`
	TermUnit  *string `json:"termUnit,omitempty"`
}
