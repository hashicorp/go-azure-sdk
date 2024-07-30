package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsClientUserAgentProductFamily string

const (
	CallRecordsClientUserAgentProductFamily_AzureCommunicationServices CallRecordsClientUserAgentProductFamily = "azureCommunicationServices"
	CallRecordsClientUserAgentProductFamily_Lync                       CallRecordsClientUserAgentProductFamily = "lync"
	CallRecordsClientUserAgentProductFamily_SkypeForBusiness           CallRecordsClientUserAgentProductFamily = "skypeForBusiness"
	CallRecordsClientUserAgentProductFamily_Teams                      CallRecordsClientUserAgentProductFamily = "teams"
	CallRecordsClientUserAgentProductFamily_Unknown                    CallRecordsClientUserAgentProductFamily = "unknown"
)

func PossibleValuesForCallRecordsClientUserAgentProductFamily() []string {
	return []string{
		string(CallRecordsClientUserAgentProductFamily_AzureCommunicationServices),
		string(CallRecordsClientUserAgentProductFamily_Lync),
		string(CallRecordsClientUserAgentProductFamily_SkypeForBusiness),
		string(CallRecordsClientUserAgentProductFamily_Teams),
		string(CallRecordsClientUserAgentProductFamily_Unknown),
	}
}

func (s *CallRecordsClientUserAgentProductFamily) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsClientUserAgentProductFamily(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsClientUserAgentProductFamily(input string) (*CallRecordsClientUserAgentProductFamily, error) {
	vals := map[string]CallRecordsClientUserAgentProductFamily{
		"azurecommunicationservices": CallRecordsClientUserAgentProductFamily_AzureCommunicationServices,
		"lync":                       CallRecordsClientUserAgentProductFamily_Lync,
		"skypeforbusiness":           CallRecordsClientUserAgentProductFamily_SkypeForBusiness,
		"teams":                      CallRecordsClientUserAgentProductFamily_Teams,
		"unknown":                    CallRecordsClientUserAgentProductFamily_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsClientUserAgentProductFamily(input)
	return &out, nil
}
