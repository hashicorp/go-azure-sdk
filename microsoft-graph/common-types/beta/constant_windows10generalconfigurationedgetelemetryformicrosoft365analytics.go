package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics string

const (
	Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics_Internet            Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics = "internet"
	Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics_Intranet            Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics = "intranet"
	Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics_IntranetAndInternet Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics = "intranetAndInternet"
	Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics_NotConfigured       Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics = "notConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics() []string {
	return []string{
		string(Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics_Internet),
		string(Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics_Intranet),
		string(Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics_IntranetAndInternet),
		string(Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics_NotConfigured),
	}
}

func (s *Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics(input string) (*Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics, error) {
	vals := map[string]Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics{
		"internet":            Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics_Internet,
		"intranet":            Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics_Intranet,
		"intranetandinternet": Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics_IntranetAndInternet,
		"notconfigured":       Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics(input)
	return &out, nil
}
