package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IpApplicationSegmentProtocol string

const (
	IpApplicationSegmentProtocol_Tcp IpApplicationSegmentProtocol = "tcp"
	IpApplicationSegmentProtocol_Udp IpApplicationSegmentProtocol = "udp"
)

func PossibleValuesForIpApplicationSegmentProtocol() []string {
	return []string{
		string(IpApplicationSegmentProtocol_Tcp),
		string(IpApplicationSegmentProtocol_Udp),
	}
}

func (s *IpApplicationSegmentProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIpApplicationSegmentProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIpApplicationSegmentProtocol(input string) (*IpApplicationSegmentProtocol, error) {
	vals := map[string]IpApplicationSegmentProtocol{
		"tcp": IpApplicationSegmentProtocol_Tcp,
		"udp": IpApplicationSegmentProtocol_Udp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IpApplicationSegmentProtocol(input)
	return &out, nil
}
