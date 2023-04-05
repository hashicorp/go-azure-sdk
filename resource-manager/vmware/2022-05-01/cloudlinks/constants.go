package cloudlinks

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
