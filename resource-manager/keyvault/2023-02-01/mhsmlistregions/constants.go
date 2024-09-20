package mhsmlistregions

import (
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GeoReplicationRegionProvisioningState string

const (
	GeoReplicationRegionProvisioningStateCleanup         GeoReplicationRegionProvisioningState = "Cleanup"
	GeoReplicationRegionProvisioningStateDeleting        GeoReplicationRegionProvisioningState = "Deleting"
	GeoReplicationRegionProvisioningStateFailed          GeoReplicationRegionProvisioningState = "Failed"
	GeoReplicationRegionProvisioningStatePreprovisioning GeoReplicationRegionProvisioningState = "Preprovisioning"
	GeoReplicationRegionProvisioningStateProvisioning    GeoReplicationRegionProvisioningState = "Provisioning"
	GeoReplicationRegionProvisioningStateSucceeded       GeoReplicationRegionProvisioningState = "Succeeded"
)

func PossibleValuesForGeoReplicationRegionProvisioningState() []string {
	return []string{
		string(GeoReplicationRegionProvisioningStateCleanup),
		string(GeoReplicationRegionProvisioningStateDeleting),
		string(GeoReplicationRegionProvisioningStateFailed),
		string(GeoReplicationRegionProvisioningStatePreprovisioning),
		string(GeoReplicationRegionProvisioningStateProvisioning),
		string(GeoReplicationRegionProvisioningStateSucceeded),
	}
}

func parseGeoReplicationRegionProvisioningState(input string) (*GeoReplicationRegionProvisioningState, error) {
	vals := map[string]GeoReplicationRegionProvisioningState{
		"cleanup":         GeoReplicationRegionProvisioningStateCleanup,
		"deleting":        GeoReplicationRegionProvisioningStateDeleting,
		"failed":          GeoReplicationRegionProvisioningStateFailed,
		"preprovisioning": GeoReplicationRegionProvisioningStatePreprovisioning,
		"provisioning":    GeoReplicationRegionProvisioningStateProvisioning,
		"succeeded":       GeoReplicationRegionProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GeoReplicationRegionProvisioningState(input)
	return &out, nil
}
