package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionAllowedOutboundDataTransferDestinations string

const (
	ManagedAppProtectionAllowedOutboundDataTransferDestinationsallApps     ManagedAppProtectionAllowedOutboundDataTransferDestinations = "AllApps"
	ManagedAppProtectionAllowedOutboundDataTransferDestinationsmanagedApps ManagedAppProtectionAllowedOutboundDataTransferDestinations = "ManagedApps"
	ManagedAppProtectionAllowedOutboundDataTransferDestinationsnone        ManagedAppProtectionAllowedOutboundDataTransferDestinations = "None"
)

func PossibleValuesForManagedAppProtectionAllowedOutboundDataTransferDestinations() []string {
	return []string{
		string(ManagedAppProtectionAllowedOutboundDataTransferDestinationsallApps),
		string(ManagedAppProtectionAllowedOutboundDataTransferDestinationsmanagedApps),
		string(ManagedAppProtectionAllowedOutboundDataTransferDestinationsnone),
	}
}

func parseManagedAppProtectionAllowedOutboundDataTransferDestinations(input string) (*ManagedAppProtectionAllowedOutboundDataTransferDestinations, error) {
	vals := map[string]ManagedAppProtectionAllowedOutboundDataTransferDestinations{
		"allapps":     ManagedAppProtectionAllowedOutboundDataTransferDestinationsallApps,
		"managedapps": ManagedAppProtectionAllowedOutboundDataTransferDestinationsmanagedApps,
		"none":        ManagedAppProtectionAllowedOutboundDataTransferDestinationsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionAllowedOutboundDataTransferDestinations(input)
	return &out, nil
}
