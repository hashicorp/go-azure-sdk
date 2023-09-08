package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAllowedOutboundDataTransferDestinations string

const (
	IosManagedAppProtectionAllowedOutboundDataTransferDestinationsallApps     IosManagedAppProtectionAllowedOutboundDataTransferDestinations = "AllApps"
	IosManagedAppProtectionAllowedOutboundDataTransferDestinationsmanagedApps IosManagedAppProtectionAllowedOutboundDataTransferDestinations = "ManagedApps"
	IosManagedAppProtectionAllowedOutboundDataTransferDestinationsnone        IosManagedAppProtectionAllowedOutboundDataTransferDestinations = "None"
)

func PossibleValuesForIosManagedAppProtectionAllowedOutboundDataTransferDestinations() []string {
	return []string{
		string(IosManagedAppProtectionAllowedOutboundDataTransferDestinationsallApps),
		string(IosManagedAppProtectionAllowedOutboundDataTransferDestinationsmanagedApps),
		string(IosManagedAppProtectionAllowedOutboundDataTransferDestinationsnone),
	}
}

func parseIosManagedAppProtectionAllowedOutboundDataTransferDestinations(input string) (*IosManagedAppProtectionAllowedOutboundDataTransferDestinations, error) {
	vals := map[string]IosManagedAppProtectionAllowedOutboundDataTransferDestinations{
		"allapps":     IosManagedAppProtectionAllowedOutboundDataTransferDestinationsallApps,
		"managedapps": IosManagedAppProtectionAllowedOutboundDataTransferDestinationsmanagedApps,
		"none":        IosManagedAppProtectionAllowedOutboundDataTransferDestinationsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAllowedOutboundDataTransferDestinations(input)
	return &out, nil
}
