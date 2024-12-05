package marketplacesubscription

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarketplaceSubscriptionProvisioningState string

const (
	MarketplaceSubscriptionProvisioningStateCanceled  MarketplaceSubscriptionProvisioningState = "Canceled"
	MarketplaceSubscriptionProvisioningStateCreating  MarketplaceSubscriptionProvisioningState = "Creating"
	MarketplaceSubscriptionProvisioningStateDeleting  MarketplaceSubscriptionProvisioningState = "Deleting"
	MarketplaceSubscriptionProvisioningStateFailed    MarketplaceSubscriptionProvisioningState = "Failed"
	MarketplaceSubscriptionProvisioningStateSucceeded MarketplaceSubscriptionProvisioningState = "Succeeded"
	MarketplaceSubscriptionProvisioningStateUpdating  MarketplaceSubscriptionProvisioningState = "Updating"
)

func PossibleValuesForMarketplaceSubscriptionProvisioningState() []string {
	return []string{
		string(MarketplaceSubscriptionProvisioningStateCanceled),
		string(MarketplaceSubscriptionProvisioningStateCreating),
		string(MarketplaceSubscriptionProvisioningStateDeleting),
		string(MarketplaceSubscriptionProvisioningStateFailed),
		string(MarketplaceSubscriptionProvisioningStateSucceeded),
		string(MarketplaceSubscriptionProvisioningStateUpdating),
	}
}

func (s *MarketplaceSubscriptionProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMarketplaceSubscriptionProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMarketplaceSubscriptionProvisioningState(input string) (*MarketplaceSubscriptionProvisioningState, error) {
	vals := map[string]MarketplaceSubscriptionProvisioningState{
		"canceled":  MarketplaceSubscriptionProvisioningStateCanceled,
		"creating":  MarketplaceSubscriptionProvisioningStateCreating,
		"deleting":  MarketplaceSubscriptionProvisioningStateDeleting,
		"failed":    MarketplaceSubscriptionProvisioningStateFailed,
		"succeeded": MarketplaceSubscriptionProvisioningStateSucceeded,
		"updating":  MarketplaceSubscriptionProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MarketplaceSubscriptionProvisioningState(input)
	return &out, nil
}

type MarketplaceSubscriptionStatus string

const (
	MarketplaceSubscriptionStatusSubscribed   MarketplaceSubscriptionStatus = "Subscribed"
	MarketplaceSubscriptionStatusSuspended    MarketplaceSubscriptionStatus = "Suspended"
	MarketplaceSubscriptionStatusUnsubscribed MarketplaceSubscriptionStatus = "Unsubscribed"
)

func PossibleValuesForMarketplaceSubscriptionStatus() []string {
	return []string{
		string(MarketplaceSubscriptionStatusSubscribed),
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
		"subscribed":   MarketplaceSubscriptionStatusSubscribed,
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
