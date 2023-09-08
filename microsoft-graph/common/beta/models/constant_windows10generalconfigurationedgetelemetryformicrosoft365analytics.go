package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics string

const (
	Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analyticsinternet            Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics = "Internet"
	Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analyticsintranet            Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics = "Intranet"
	Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365AnalyticsintranetAndInternet Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics = "IntranetAndInternet"
	Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365AnalyticsnotConfigured       Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics = "NotConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics() []string {
	return []string{
		string(Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analyticsinternet),
		string(Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analyticsintranet),
		string(Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365AnalyticsintranetAndInternet),
		string(Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365AnalyticsnotConfigured),
	}
}

func parseWindows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics(input string) (*Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics, error) {
	vals := map[string]Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics{
		"internet":            Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analyticsinternet,
		"intranet":            Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analyticsintranet,
		"intranetandinternet": Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365AnalyticsintranetAndInternet,
		"notconfigured":       Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365AnalyticsnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationEdgeTelemetryForMicrosoft365Analytics(input)
	return &out, nil
}
