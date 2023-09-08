package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations string

const (
	DefaultManagedAppProtectionAllowedOutboundDataTransferDestinationsallApps     DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations = "AllApps"
	DefaultManagedAppProtectionAllowedOutboundDataTransferDestinationsmanagedApps DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations = "ManagedApps"
	DefaultManagedAppProtectionAllowedOutboundDataTransferDestinationsnone        DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations = "None"
)

func PossibleValuesForDefaultManagedAppProtectionAllowedOutboundDataTransferDestinations() []string {
	return []string{
		string(DefaultManagedAppProtectionAllowedOutboundDataTransferDestinationsallApps),
		string(DefaultManagedAppProtectionAllowedOutboundDataTransferDestinationsmanagedApps),
		string(DefaultManagedAppProtectionAllowedOutboundDataTransferDestinationsnone),
	}
}

func parseDefaultManagedAppProtectionAllowedOutboundDataTransferDestinations(input string) (*DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations, error) {
	vals := map[string]DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations{
		"allapps":     DefaultManagedAppProtectionAllowedOutboundDataTransferDestinationsallApps,
		"managedapps": DefaultManagedAppProtectionAllowedOutboundDataTransferDestinationsmanagedApps,
		"none":        DefaultManagedAppProtectionAllowedOutboundDataTransferDestinationsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations(input)
	return &out, nil
}
