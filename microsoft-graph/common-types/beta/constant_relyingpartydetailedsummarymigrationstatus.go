package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RelyingPartyDetailedSummaryMigrationStatus string

const (
	RelyingPartyDetailedSummaryMigrationStatus_AdditionalStepsRequired RelyingPartyDetailedSummaryMigrationStatus = "additionalStepsRequired"
	RelyingPartyDetailedSummaryMigrationStatus_NeedsReview             RelyingPartyDetailedSummaryMigrationStatus = "needsReview"
	RelyingPartyDetailedSummaryMigrationStatus_Ready                   RelyingPartyDetailedSummaryMigrationStatus = "ready"
)

func PossibleValuesForRelyingPartyDetailedSummaryMigrationStatus() []string {
	return []string{
		string(RelyingPartyDetailedSummaryMigrationStatus_AdditionalStepsRequired),
		string(RelyingPartyDetailedSummaryMigrationStatus_NeedsReview),
		string(RelyingPartyDetailedSummaryMigrationStatus_Ready),
	}
}

func (s *RelyingPartyDetailedSummaryMigrationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRelyingPartyDetailedSummaryMigrationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRelyingPartyDetailedSummaryMigrationStatus(input string) (*RelyingPartyDetailedSummaryMigrationStatus, error) {
	vals := map[string]RelyingPartyDetailedSummaryMigrationStatus{
		"additionalstepsrequired": RelyingPartyDetailedSummaryMigrationStatus_AdditionalStepsRequired,
		"needsreview":             RelyingPartyDetailedSummaryMigrationStatus_NeedsReview,
		"ready":                   RelyingPartyDetailedSummaryMigrationStatus_Ready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RelyingPartyDetailedSummaryMigrationStatus(input)
	return &out, nil
}
