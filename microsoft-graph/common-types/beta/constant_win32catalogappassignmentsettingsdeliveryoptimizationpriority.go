package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32CatalogAppAssignmentSettingsDeliveryOptimizationPriority string

const (
	Win32CatalogAppAssignmentSettingsDeliveryOptimizationPriority_Foreground    Win32CatalogAppAssignmentSettingsDeliveryOptimizationPriority = "foreground"
	Win32CatalogAppAssignmentSettingsDeliveryOptimizationPriority_NotConfigured Win32CatalogAppAssignmentSettingsDeliveryOptimizationPriority = "notConfigured"
)

func PossibleValuesForWin32CatalogAppAssignmentSettingsDeliveryOptimizationPriority() []string {
	return []string{
		string(Win32CatalogAppAssignmentSettingsDeliveryOptimizationPriority_Foreground),
		string(Win32CatalogAppAssignmentSettingsDeliveryOptimizationPriority_NotConfigured),
	}
}

func (s *Win32CatalogAppAssignmentSettingsDeliveryOptimizationPriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32CatalogAppAssignmentSettingsDeliveryOptimizationPriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32CatalogAppAssignmentSettingsDeliveryOptimizationPriority(input string) (*Win32CatalogAppAssignmentSettingsDeliveryOptimizationPriority, error) {
	vals := map[string]Win32CatalogAppAssignmentSettingsDeliveryOptimizationPriority{
		"foreground":    Win32CatalogAppAssignmentSettingsDeliveryOptimizationPriority_Foreground,
		"notconfigured": Win32CatalogAppAssignmentSettingsDeliveryOptimizationPriority_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32CatalogAppAssignmentSettingsDeliveryOptimizationPriority(input)
	return &out, nil
}
