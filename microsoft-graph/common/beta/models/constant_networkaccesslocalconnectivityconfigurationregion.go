package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessLocalConnectivityConfigurationRegion string

const (
	NetworkaccessLocalConnectivityConfigurationRegioncanadaCentral      NetworkaccessLocalConnectivityConfigurationRegion = "CanadaCentral"
	NetworkaccessLocalConnectivityConfigurationRegioncanadaEast         NetworkaccessLocalConnectivityConfigurationRegion = "CanadaEast"
	NetworkaccessLocalConnectivityConfigurationRegioncentralUS          NetworkaccessLocalConnectivityConfigurationRegion = "CentralUS"
	NetworkaccessLocalConnectivityConfigurationRegioneastUS             NetworkaccessLocalConnectivityConfigurationRegion = "EastUS"
	NetworkaccessLocalConnectivityConfigurationRegioneastUS2            NetworkaccessLocalConnectivityConfigurationRegion = "EastUS2"
	NetworkaccessLocalConnectivityConfigurationRegionfranceCentral      NetworkaccessLocalConnectivityConfigurationRegion = "FranceCentral"
	NetworkaccessLocalConnectivityConfigurationRegiongermanyWestCentral NetworkaccessLocalConnectivityConfigurationRegion = "GermanyWestCentral"
	NetworkaccessLocalConnectivityConfigurationRegionnorthCentralUS     NetworkaccessLocalConnectivityConfigurationRegion = "NorthCentralUS"
	NetworkaccessLocalConnectivityConfigurationRegionnorthEurope        NetworkaccessLocalConnectivityConfigurationRegion = "NorthEurope"
	NetworkaccessLocalConnectivityConfigurationRegionsouthAfricaNorth   NetworkaccessLocalConnectivityConfigurationRegion = "SouthAfricaNorth"
	NetworkaccessLocalConnectivityConfigurationRegionsouthAfricaWest    NetworkaccessLocalConnectivityConfigurationRegion = "SouthAfricaWest"
	NetworkaccessLocalConnectivityConfigurationRegionsouthCentralUS     NetworkaccessLocalConnectivityConfigurationRegion = "SouthCentralUS"
	NetworkaccessLocalConnectivityConfigurationRegionswitzerlandNorth   NetworkaccessLocalConnectivityConfigurationRegion = "SwitzerlandNorth"
	NetworkaccessLocalConnectivityConfigurationRegionuaeNorth           NetworkaccessLocalConnectivityConfigurationRegion = "UaeNorth"
	NetworkaccessLocalConnectivityConfigurationRegionukSouth            NetworkaccessLocalConnectivityConfigurationRegion = "UkSouth"
	NetworkaccessLocalConnectivityConfigurationRegionwestEurope         NetworkaccessLocalConnectivityConfigurationRegion = "WestEurope"
	NetworkaccessLocalConnectivityConfigurationRegionwestUS             NetworkaccessLocalConnectivityConfigurationRegion = "WestUS"
	NetworkaccessLocalConnectivityConfigurationRegionwestUS2            NetworkaccessLocalConnectivityConfigurationRegion = "WestUS2"
	NetworkaccessLocalConnectivityConfigurationRegionwestUS3            NetworkaccessLocalConnectivityConfigurationRegion = "WestUS3"
)

func PossibleValuesForNetworkaccessLocalConnectivityConfigurationRegion() []string {
	return []string{
		string(NetworkaccessLocalConnectivityConfigurationRegioncanadaCentral),
		string(NetworkaccessLocalConnectivityConfigurationRegioncanadaEast),
		string(NetworkaccessLocalConnectivityConfigurationRegioncentralUS),
		string(NetworkaccessLocalConnectivityConfigurationRegioneastUS),
		string(NetworkaccessLocalConnectivityConfigurationRegioneastUS2),
		string(NetworkaccessLocalConnectivityConfigurationRegionfranceCentral),
		string(NetworkaccessLocalConnectivityConfigurationRegiongermanyWestCentral),
		string(NetworkaccessLocalConnectivityConfigurationRegionnorthCentralUS),
		string(NetworkaccessLocalConnectivityConfigurationRegionnorthEurope),
		string(NetworkaccessLocalConnectivityConfigurationRegionsouthAfricaNorth),
		string(NetworkaccessLocalConnectivityConfigurationRegionsouthAfricaWest),
		string(NetworkaccessLocalConnectivityConfigurationRegionsouthCentralUS),
		string(NetworkaccessLocalConnectivityConfigurationRegionswitzerlandNorth),
		string(NetworkaccessLocalConnectivityConfigurationRegionuaeNorth),
		string(NetworkaccessLocalConnectivityConfigurationRegionukSouth),
		string(NetworkaccessLocalConnectivityConfigurationRegionwestEurope),
		string(NetworkaccessLocalConnectivityConfigurationRegionwestUS),
		string(NetworkaccessLocalConnectivityConfigurationRegionwestUS2),
		string(NetworkaccessLocalConnectivityConfigurationRegionwestUS3),
	}
}

func parseNetworkaccessLocalConnectivityConfigurationRegion(input string) (*NetworkaccessLocalConnectivityConfigurationRegion, error) {
	vals := map[string]NetworkaccessLocalConnectivityConfigurationRegion{
		"canadacentral":      NetworkaccessLocalConnectivityConfigurationRegioncanadaCentral,
		"canadaeast":         NetworkaccessLocalConnectivityConfigurationRegioncanadaEast,
		"centralus":          NetworkaccessLocalConnectivityConfigurationRegioncentralUS,
		"eastus":             NetworkaccessLocalConnectivityConfigurationRegioneastUS,
		"eastus2":            NetworkaccessLocalConnectivityConfigurationRegioneastUS2,
		"francecentral":      NetworkaccessLocalConnectivityConfigurationRegionfranceCentral,
		"germanywestcentral": NetworkaccessLocalConnectivityConfigurationRegiongermanyWestCentral,
		"northcentralus":     NetworkaccessLocalConnectivityConfigurationRegionnorthCentralUS,
		"northeurope":        NetworkaccessLocalConnectivityConfigurationRegionnorthEurope,
		"southafricanorth":   NetworkaccessLocalConnectivityConfigurationRegionsouthAfricaNorth,
		"southafricawest":    NetworkaccessLocalConnectivityConfigurationRegionsouthAfricaWest,
		"southcentralus":     NetworkaccessLocalConnectivityConfigurationRegionsouthCentralUS,
		"switzerlandnorth":   NetworkaccessLocalConnectivityConfigurationRegionswitzerlandNorth,
		"uaenorth":           NetworkaccessLocalConnectivityConfigurationRegionuaeNorth,
		"uksouth":            NetworkaccessLocalConnectivityConfigurationRegionukSouth,
		"westeurope":         NetworkaccessLocalConnectivityConfigurationRegionwestEurope,
		"westus":             NetworkaccessLocalConnectivityConfigurationRegionwestUS,
		"westus2":            NetworkaccessLocalConnectivityConfigurationRegionwestUS2,
		"westus3":            NetworkaccessLocalConnectivityConfigurationRegionwestUS3,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessLocalConnectivityConfigurationRegion(input)
	return &out, nil
}
