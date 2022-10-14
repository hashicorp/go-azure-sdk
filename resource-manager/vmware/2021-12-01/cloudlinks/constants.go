package cloudlinks

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudLinkStatus string

const (
	CloudLinkStatusActive       CloudLinkStatus = "Active"
	CloudLinkStatusBuilding     CloudLinkStatus = "Building"
	CloudLinkStatusDeleting     CloudLinkStatus = "Deleting"
	CloudLinkStatusDisconnected CloudLinkStatus = "Disconnected"
	CloudLinkStatusFailed       CloudLinkStatus = "Failed"
)

func PossibleValuesForCloudLinkStatus() []string {
	return []string{
		string(CloudLinkStatusActive),
		string(CloudLinkStatusBuilding),
		string(CloudLinkStatusDeleting),
		string(CloudLinkStatusDisconnected),
		string(CloudLinkStatusFailed),
	}
}

func parseCloudLinkStatus(input string) (*CloudLinkStatus, error) {
	vals := map[string]CloudLinkStatus{
		"active":       CloudLinkStatusActive,
		"building":     CloudLinkStatusBuilding,
		"deleting":     CloudLinkStatusDeleting,
		"disconnected": CloudLinkStatusDisconnected,
		"failed":       CloudLinkStatusFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudLinkStatus(input)
	return &out, nil
}
