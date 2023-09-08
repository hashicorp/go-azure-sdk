package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations string

const (
	AndroidManagedAppProtectionAllowedOutboundDataTransferDestinationsallApps     AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations = "AllApps"
	AndroidManagedAppProtectionAllowedOutboundDataTransferDestinationsmanagedApps AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations = "ManagedApps"
	AndroidManagedAppProtectionAllowedOutboundDataTransferDestinationsnone        AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations = "None"
)

func PossibleValuesForAndroidManagedAppProtectionAllowedOutboundDataTransferDestinations() []string {
	return []string{
		string(AndroidManagedAppProtectionAllowedOutboundDataTransferDestinationsallApps),
		string(AndroidManagedAppProtectionAllowedOutboundDataTransferDestinationsmanagedApps),
		string(AndroidManagedAppProtectionAllowedOutboundDataTransferDestinationsnone),
	}
}

func parseAndroidManagedAppProtectionAllowedOutboundDataTransferDestinations(input string) (*AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations, error) {
	vals := map[string]AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations{
		"allapps":     AndroidManagedAppProtectionAllowedOutboundDataTransferDestinationsallApps,
		"managedapps": AndroidManagedAppProtectionAllowedOutboundDataTransferDestinationsmanagedApps,
		"none":        AndroidManagedAppProtectionAllowedOutboundDataTransferDestinationsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations(input)
	return &out, nil
}
