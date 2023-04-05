package checkfrontdoornameavailability

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Availability string

const (
	AvailabilityAvailable   Availability = "Available"
	AvailabilityUnavailable Availability = "Unavailable"
)

func PossibleValuesForAvailability() []string {
	return []string{
		string(AvailabilityAvailable),
		string(AvailabilityUnavailable),
	}
}

type ResourceType string

const (
	ResourceTypeMicrosoftPointNetworkFrontDoors                  ResourceType = "Microsoft.Network/frontDoors"
	ResourceTypeMicrosoftPointNetworkFrontDoorsFrontendEndpoints ResourceType = "Microsoft.Network/frontDoors/frontendEndpoints"
)

func PossibleValuesForResourceType() []string {
	return []string{
		string(ResourceTypeMicrosoftPointNetworkFrontDoors),
		string(ResourceTypeMicrosoftPointNetworkFrontDoorsFrontendEndpoints),
	}
}
