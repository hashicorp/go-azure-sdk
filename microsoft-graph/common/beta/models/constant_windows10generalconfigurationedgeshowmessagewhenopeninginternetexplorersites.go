package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites string

const (
	Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSitesdisabled      Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites = "Disabled"
	Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSitesenabled       Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites = "Enabled"
	Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSiteskeepGoing     Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites = "KeepGoing"
	Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSitesnotConfigured Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites = "NotConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites() []string {
	return []string{
		string(Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSitesdisabled),
		string(Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSitesenabled),
		string(Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSiteskeepGoing),
		string(Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSitesnotConfigured),
	}
}

func parseWindows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites(input string) (*Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites, error) {
	vals := map[string]Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites{
		"disabled":      Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSitesdisabled,
		"enabled":       Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSitesenabled,
		"keepgoing":     Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSiteskeepGoing,
		"notconfigured": Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSitesnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites(input)
	return &out, nil
}
