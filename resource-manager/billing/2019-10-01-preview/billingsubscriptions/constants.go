package billingsubscriptions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingSubscriptionStatusType string

const (
	BillingSubscriptionStatusTypeAbandoned BillingSubscriptionStatusType = "Abandoned"
	BillingSubscriptionStatusTypeActive    BillingSubscriptionStatusType = "Active"
	BillingSubscriptionStatusTypeDeleted   BillingSubscriptionStatusType = "Deleted"
	BillingSubscriptionStatusTypeInactive  BillingSubscriptionStatusType = "Inactive"
	BillingSubscriptionStatusTypeWarning   BillingSubscriptionStatusType = "Warning"
)

func PossibleValuesForBillingSubscriptionStatusType() []string {
	return []string{
		string(BillingSubscriptionStatusTypeAbandoned),
		string(BillingSubscriptionStatusTypeActive),
		string(BillingSubscriptionStatusTypeDeleted),
		string(BillingSubscriptionStatusTypeInactive),
		string(BillingSubscriptionStatusTypeWarning),
	}
}

func (s *BillingSubscriptionStatusType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingSubscriptionStatusType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingSubscriptionStatusType(input string) (*BillingSubscriptionStatusType, error) {
	vals := map[string]BillingSubscriptionStatusType{
		"abandoned": BillingSubscriptionStatusTypeAbandoned,
		"active":    BillingSubscriptionStatusTypeActive,
		"deleted":   BillingSubscriptionStatusTypeDeleted,
		"inactive":  BillingSubscriptionStatusTypeInactive,
		"warning":   BillingSubscriptionStatusTypeWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingSubscriptionStatusType(input)
	return &out, nil
}
