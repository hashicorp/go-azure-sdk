package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IpApplicationSegmentDestinationType string

const (
	IpApplicationSegmentDestinationType_DnsSuffix   IpApplicationSegmentDestinationType = "dnsSuffix"
	IpApplicationSegmentDestinationType_Fqdn        IpApplicationSegmentDestinationType = "fqdn"
	IpApplicationSegmentDestinationType_IpAddress   IpApplicationSegmentDestinationType = "ipAddress"
	IpApplicationSegmentDestinationType_IpRange     IpApplicationSegmentDestinationType = "ipRange"
	IpApplicationSegmentDestinationType_IpRangeCidr IpApplicationSegmentDestinationType = "ipRangeCidr"
)

func PossibleValuesForIpApplicationSegmentDestinationType() []string {
	return []string{
		string(IpApplicationSegmentDestinationType_DnsSuffix),
		string(IpApplicationSegmentDestinationType_Fqdn),
		string(IpApplicationSegmentDestinationType_IpAddress),
		string(IpApplicationSegmentDestinationType_IpRange),
		string(IpApplicationSegmentDestinationType_IpRangeCidr),
	}
}

func (s *IpApplicationSegmentDestinationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIpApplicationSegmentDestinationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIpApplicationSegmentDestinationType(input string) (*IpApplicationSegmentDestinationType, error) {
	vals := map[string]IpApplicationSegmentDestinationType{
		"dnssuffix":   IpApplicationSegmentDestinationType_DnsSuffix,
		"fqdn":        IpApplicationSegmentDestinationType_Fqdn,
		"ipaddress":   IpApplicationSegmentDestinationType_IpAddress,
		"iprange":     IpApplicationSegmentDestinationType_IpRange,
		"iprangecidr": IpApplicationSegmentDestinationType_IpRangeCidr,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IpApplicationSegmentDestinationType(input)
	return &out, nil
}
