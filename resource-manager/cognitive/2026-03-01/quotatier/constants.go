package quotatier

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TierUpgradePolicy string

const (
	TierUpgradePolicyNoAutoUpgrade          TierUpgradePolicy = "NoAutoUpgrade"
	TierUpgradePolicyOnceUpgradeIsAvailable TierUpgradePolicy = "OnceUpgradeIsAvailable"
)

func PossibleValuesForTierUpgradePolicy() []string {
	return []string{
		string(TierUpgradePolicyNoAutoUpgrade),
		string(TierUpgradePolicyOnceUpgradeIsAvailable),
	}
}

func (s *TierUpgradePolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTierUpgradePolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTierUpgradePolicy(input string) (*TierUpgradePolicy, error) {
	vals := map[string]TierUpgradePolicy{
		"noautoupgrade":          TierUpgradePolicyNoAutoUpgrade,
		"onceupgradeisavailable": TierUpgradePolicyOnceUpgradeIsAvailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TierUpgradePolicy(input)
	return &out, nil
}

type UpgradeAvailabilityStatus string

const (
	UpgradeAvailabilityStatusAvailable    UpgradeAvailabilityStatus = "Available"
	UpgradeAvailabilityStatusNotAvailable UpgradeAvailabilityStatus = "NotAvailable"
)

func PossibleValuesForUpgradeAvailabilityStatus() []string {
	return []string{
		string(UpgradeAvailabilityStatusAvailable),
		string(UpgradeAvailabilityStatusNotAvailable),
	}
}

func (s *UpgradeAvailabilityStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUpgradeAvailabilityStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUpgradeAvailabilityStatus(input string) (*UpgradeAvailabilityStatus, error) {
	vals := map[string]UpgradeAvailabilityStatus{
		"available":    UpgradeAvailabilityStatusAvailable,
		"notavailable": UpgradeAvailabilityStatusNotAvailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpgradeAvailabilityStatus(input)
	return &out, nil
}
