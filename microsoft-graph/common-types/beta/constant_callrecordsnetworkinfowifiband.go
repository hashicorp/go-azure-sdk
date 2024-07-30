package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsNetworkInfoWifiBand string

const (
	CallRecordsNetworkInfoWifiBand_Frequency24GHz CallRecordsNetworkInfoWifiBand = "frequency24GHz"
	CallRecordsNetworkInfoWifiBand_Frequency50GHz CallRecordsNetworkInfoWifiBand = "frequency50GHz"
	CallRecordsNetworkInfoWifiBand_Frequency60GHz CallRecordsNetworkInfoWifiBand = "frequency60GHz"
	CallRecordsNetworkInfoWifiBand_Unknown        CallRecordsNetworkInfoWifiBand = "unknown"
)

func PossibleValuesForCallRecordsNetworkInfoWifiBand() []string {
	return []string{
		string(CallRecordsNetworkInfoWifiBand_Frequency24GHz),
		string(CallRecordsNetworkInfoWifiBand_Frequency50GHz),
		string(CallRecordsNetworkInfoWifiBand_Frequency60GHz),
		string(CallRecordsNetworkInfoWifiBand_Unknown),
	}
}

func (s *CallRecordsNetworkInfoWifiBand) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsNetworkInfoWifiBand(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsNetworkInfoWifiBand(input string) (*CallRecordsNetworkInfoWifiBand, error) {
	vals := map[string]CallRecordsNetworkInfoWifiBand{
		"frequency24ghz": CallRecordsNetworkInfoWifiBand_Frequency24GHz,
		"frequency50ghz": CallRecordsNetworkInfoWifiBand_Frequency50GHz,
		"frequency60ghz": CallRecordsNetworkInfoWifiBand_Frequency60GHz,
		"unknown":        CallRecordsNetworkInfoWifiBand_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsNetworkInfoWifiBand(input)
	return &out, nil
}
