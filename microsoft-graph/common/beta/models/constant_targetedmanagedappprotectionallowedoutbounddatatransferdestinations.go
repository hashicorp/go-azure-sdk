package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations string

const (
	TargetedManagedAppProtectionAllowedOutboundDataTransferDestinationsallApps     TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations = "AllApps"
	TargetedManagedAppProtectionAllowedOutboundDataTransferDestinationsmanagedApps TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations = "ManagedApps"
	TargetedManagedAppProtectionAllowedOutboundDataTransferDestinationsnone        TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations = "None"
)

func PossibleValuesForTargetedManagedAppProtectionAllowedOutboundDataTransferDestinations() []string {
	return []string{
		string(TargetedManagedAppProtectionAllowedOutboundDataTransferDestinationsallApps),
		string(TargetedManagedAppProtectionAllowedOutboundDataTransferDestinationsmanagedApps),
		string(TargetedManagedAppProtectionAllowedOutboundDataTransferDestinationsnone),
	}
}

func parseTargetedManagedAppProtectionAllowedOutboundDataTransferDestinations(input string) (*TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations, error) {
	vals := map[string]TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations{
		"allapps":     TargetedManagedAppProtectionAllowedOutboundDataTransferDestinationsallApps,
		"managedapps": TargetedManagedAppProtectionAllowedOutboundDataTransferDestinationsmanagedApps,
		"none":        TargetedManagedAppProtectionAllowedOutboundDataTransferDestinationsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations(input)
	return &out, nil
}
