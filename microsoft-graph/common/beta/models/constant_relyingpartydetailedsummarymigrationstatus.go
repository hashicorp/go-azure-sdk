package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RelyingPartyDetailedSummaryMigrationStatus string

const (
	RelyingPartyDetailedSummaryMigrationStatusadditionalStepsRequired RelyingPartyDetailedSummaryMigrationStatus = "AdditionalStepsRequired"
	RelyingPartyDetailedSummaryMigrationStatusneedsReview             RelyingPartyDetailedSummaryMigrationStatus = "NeedsReview"
	RelyingPartyDetailedSummaryMigrationStatusready                   RelyingPartyDetailedSummaryMigrationStatus = "Ready"
)

func PossibleValuesForRelyingPartyDetailedSummaryMigrationStatus() []string {
	return []string{
		string(RelyingPartyDetailedSummaryMigrationStatusadditionalStepsRequired),
		string(RelyingPartyDetailedSummaryMigrationStatusneedsReview),
		string(RelyingPartyDetailedSummaryMigrationStatusready),
	}
}

func parseRelyingPartyDetailedSummaryMigrationStatus(input string) (*RelyingPartyDetailedSummaryMigrationStatus, error) {
	vals := map[string]RelyingPartyDetailedSummaryMigrationStatus{
		"additionalstepsrequired": RelyingPartyDetailedSummaryMigrationStatusadditionalStepsRequired,
		"needsreview":             RelyingPartyDetailedSummaryMigrationStatusneedsReview,
		"ready":                   RelyingPartyDetailedSummaryMigrationStatusready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RelyingPartyDetailedSummaryMigrationStatus(input)
	return &out, nil
}
