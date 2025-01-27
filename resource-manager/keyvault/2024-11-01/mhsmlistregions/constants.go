package mhsmlistregions

import (
	"encoding/json"
	"fmt"
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

func (s *GeoReplicationRegionProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGeoReplicationRegionProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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
