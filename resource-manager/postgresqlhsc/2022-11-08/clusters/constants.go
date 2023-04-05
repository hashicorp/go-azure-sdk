package clusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckNameAvailabilityResourceType string

const (
	CheckNameAvailabilityResourceTypeMicrosoftPointDBforPostgreSQLServerGroupsvTwo CheckNameAvailabilityResourceType = "Microsoft.DBforPostgreSQL/serverGroupsv2"
)

func PossibleValuesForCheckNameAvailabilityResourceType() []string {
	return []string{
		string(CheckNameAvailabilityResourceTypeMicrosoftPointDBforPostgreSQLServerGroupsvTwo),
	}
}

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateEndpointServiceConnectionStatus() []string {
	return []string{
		string(PrivateEndpointServiceConnectionStatusApproved),
		string(PrivateEndpointServiceConnectionStatusPending),
		string(PrivateEndpointServiceConnectionStatusRejected),
	}
}
