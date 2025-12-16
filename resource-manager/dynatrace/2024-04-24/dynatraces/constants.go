package dynatraces

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarketplaceSubscriptionStatus string

const (
	MarketplaceSubscriptionStatusActive       MarketplaceSubscriptionStatus = "Active"
	MarketplaceSubscriptionStatusSuspended    MarketplaceSubscriptionStatus = "Suspended"
	MarketplaceSubscriptionStatusUnsubscribed MarketplaceSubscriptionStatus = "Unsubscribed"
)

func PossibleValuesForMarketplaceSubscriptionStatus() []string {
	return []string{
		string(MarketplaceSubscriptionStatusActive),
		string(MarketplaceSubscriptionStatusSuspended),
		string(MarketplaceSubscriptionStatusUnsubscribed),
	}
}

func (s *MarketplaceSubscriptionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMarketplaceSubscriptionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMarketplaceSubscriptionStatus(input string) (*MarketplaceSubscriptionStatus, error) {
	vals := map[string]MarketplaceSubscriptionStatus{
		"active":       MarketplaceSubscriptionStatusActive,
		"suspended":    MarketplaceSubscriptionStatusSuspended,
		"unsubscribed": MarketplaceSubscriptionStatusUnsubscribed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MarketplaceSubscriptionStatus(input)
	return &out, nil
}
