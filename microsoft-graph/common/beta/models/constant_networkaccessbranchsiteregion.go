package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessBranchSiteRegion string

const (
	NetworkaccessBranchSiteRegioncanadaCentral      NetworkaccessBranchSiteRegion = "CanadaCentral"
	NetworkaccessBranchSiteRegioncanadaEast         NetworkaccessBranchSiteRegion = "CanadaEast"
	NetworkaccessBranchSiteRegioncentralUS          NetworkaccessBranchSiteRegion = "CentralUS"
	NetworkaccessBranchSiteRegioneastUS             NetworkaccessBranchSiteRegion = "EastUS"
	NetworkaccessBranchSiteRegioneastUS2            NetworkaccessBranchSiteRegion = "EastUS2"
	NetworkaccessBranchSiteRegionfranceCentral      NetworkaccessBranchSiteRegion = "FranceCentral"
	NetworkaccessBranchSiteRegiongermanyWestCentral NetworkaccessBranchSiteRegion = "GermanyWestCentral"
	NetworkaccessBranchSiteRegionnorthCentralUS     NetworkaccessBranchSiteRegion = "NorthCentralUS"
	NetworkaccessBranchSiteRegionnorthEurope        NetworkaccessBranchSiteRegion = "NorthEurope"
	NetworkaccessBranchSiteRegionsouthAfricaNorth   NetworkaccessBranchSiteRegion = "SouthAfricaNorth"
	NetworkaccessBranchSiteRegionsouthAfricaWest    NetworkaccessBranchSiteRegion = "SouthAfricaWest"
	NetworkaccessBranchSiteRegionsouthCentralUS     NetworkaccessBranchSiteRegion = "SouthCentralUS"
	NetworkaccessBranchSiteRegionswitzerlandNorth   NetworkaccessBranchSiteRegion = "SwitzerlandNorth"
	NetworkaccessBranchSiteRegionuaeNorth           NetworkaccessBranchSiteRegion = "UaeNorth"
	NetworkaccessBranchSiteRegionukSouth            NetworkaccessBranchSiteRegion = "UkSouth"
	NetworkaccessBranchSiteRegionwestEurope         NetworkaccessBranchSiteRegion = "WestEurope"
	NetworkaccessBranchSiteRegionwestUS             NetworkaccessBranchSiteRegion = "WestUS"
	NetworkaccessBranchSiteRegionwestUS2            NetworkaccessBranchSiteRegion = "WestUS2"
	NetworkaccessBranchSiteRegionwestUS3            NetworkaccessBranchSiteRegion = "WestUS3"
)

func PossibleValuesForNetworkaccessBranchSiteRegion() []string {
	return []string{
		string(NetworkaccessBranchSiteRegioncanadaCentral),
		string(NetworkaccessBranchSiteRegioncanadaEast),
		string(NetworkaccessBranchSiteRegioncentralUS),
		string(NetworkaccessBranchSiteRegioneastUS),
		string(NetworkaccessBranchSiteRegioneastUS2),
		string(NetworkaccessBranchSiteRegionfranceCentral),
		string(NetworkaccessBranchSiteRegiongermanyWestCentral),
		string(NetworkaccessBranchSiteRegionnorthCentralUS),
		string(NetworkaccessBranchSiteRegionnorthEurope),
		string(NetworkaccessBranchSiteRegionsouthAfricaNorth),
		string(NetworkaccessBranchSiteRegionsouthAfricaWest),
		string(NetworkaccessBranchSiteRegionsouthCentralUS),
		string(NetworkaccessBranchSiteRegionswitzerlandNorth),
		string(NetworkaccessBranchSiteRegionuaeNorth),
		string(NetworkaccessBranchSiteRegionukSouth),
		string(NetworkaccessBranchSiteRegionwestEurope),
		string(NetworkaccessBranchSiteRegionwestUS),
		string(NetworkaccessBranchSiteRegionwestUS2),
		string(NetworkaccessBranchSiteRegionwestUS3),
	}
}

func parseNetworkaccessBranchSiteRegion(input string) (*NetworkaccessBranchSiteRegion, error) {
	vals := map[string]NetworkaccessBranchSiteRegion{
		"canadacentral":      NetworkaccessBranchSiteRegioncanadaCentral,
		"canadaeast":         NetworkaccessBranchSiteRegioncanadaEast,
		"centralus":          NetworkaccessBranchSiteRegioncentralUS,
		"eastus":             NetworkaccessBranchSiteRegioneastUS,
		"eastus2":            NetworkaccessBranchSiteRegioneastUS2,
		"francecentral":      NetworkaccessBranchSiteRegionfranceCentral,
		"germanywestcentral": NetworkaccessBranchSiteRegiongermanyWestCentral,
		"northcentralus":     NetworkaccessBranchSiteRegionnorthCentralUS,
		"northeurope":        NetworkaccessBranchSiteRegionnorthEurope,
		"southafricanorth":   NetworkaccessBranchSiteRegionsouthAfricaNorth,
		"southafricawest":    NetworkaccessBranchSiteRegionsouthAfricaWest,
		"southcentralus":     NetworkaccessBranchSiteRegionsouthCentralUS,
		"switzerlandnorth":   NetworkaccessBranchSiteRegionswitzerlandNorth,
		"uaenorth":           NetworkaccessBranchSiteRegionuaeNorth,
		"uksouth":            NetworkaccessBranchSiteRegionukSouth,
		"westeurope":         NetworkaccessBranchSiteRegionwestEurope,
		"westus":             NetworkaccessBranchSiteRegionwestUS,
		"westus2":            NetworkaccessBranchSiteRegionwestUS2,
		"westus3":            NetworkaccessBranchSiteRegionwestUS3,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessBranchSiteRegion(input)
	return &out, nil
}
