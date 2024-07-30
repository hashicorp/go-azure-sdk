package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDeviceLinkBandwidthCapacityInMbps string

const (
	NetworkaccessDeviceLinkBandwidthCapacityInMbps_Mbps1000 NetworkaccessDeviceLinkBandwidthCapacityInMbps = "mbps1000"
	NetworkaccessDeviceLinkBandwidthCapacityInMbps_Mbps250  NetworkaccessDeviceLinkBandwidthCapacityInMbps = "mbps250"
	NetworkaccessDeviceLinkBandwidthCapacityInMbps_Mbps500  NetworkaccessDeviceLinkBandwidthCapacityInMbps = "mbps500"
	NetworkaccessDeviceLinkBandwidthCapacityInMbps_Mbps750  NetworkaccessDeviceLinkBandwidthCapacityInMbps = "mbps750"
)

func PossibleValuesForNetworkaccessDeviceLinkBandwidthCapacityInMbps() []string {
	return []string{
		string(NetworkaccessDeviceLinkBandwidthCapacityInMbps_Mbps1000),
		string(NetworkaccessDeviceLinkBandwidthCapacityInMbps_Mbps250),
		string(NetworkaccessDeviceLinkBandwidthCapacityInMbps_Mbps500),
		string(NetworkaccessDeviceLinkBandwidthCapacityInMbps_Mbps750),
	}
}

func (s *NetworkaccessDeviceLinkBandwidthCapacityInMbps) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessDeviceLinkBandwidthCapacityInMbps(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessDeviceLinkBandwidthCapacityInMbps(input string) (*NetworkaccessDeviceLinkBandwidthCapacityInMbps, error) {
	vals := map[string]NetworkaccessDeviceLinkBandwidthCapacityInMbps{
		"mbps1000": NetworkaccessDeviceLinkBandwidthCapacityInMbps_Mbps1000,
		"mbps250":  NetworkaccessDeviceLinkBandwidthCapacityInMbps_Mbps250,
		"mbps500":  NetworkaccessDeviceLinkBandwidthCapacityInMbps_Mbps500,
		"mbps750":  NetworkaccessDeviceLinkBandwidthCapacityInMbps_Mbps750,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessDeviceLinkBandwidthCapacityInMbps(input)
	return &out, nil
}
