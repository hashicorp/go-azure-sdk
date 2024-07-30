package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessBranchSiteRegion string

const (
	NetworkaccessBranchSiteRegion_AustraliaEast      NetworkaccessBranchSiteRegion = "australiaEast"
	NetworkaccessBranchSiteRegion_AustraliaSouthEast NetworkaccessBranchSiteRegion = "australiaSouthEast"
	NetworkaccessBranchSiteRegion_BrazilSouth        NetworkaccessBranchSiteRegion = "brazilSouth"
	NetworkaccessBranchSiteRegion_CanadaCentral      NetworkaccessBranchSiteRegion = "canadaCentral"
	NetworkaccessBranchSiteRegion_CanadaEast         NetworkaccessBranchSiteRegion = "canadaEast"
	NetworkaccessBranchSiteRegion_CentralIndia       NetworkaccessBranchSiteRegion = "centralIndia"
	NetworkaccessBranchSiteRegion_CentralUS          NetworkaccessBranchSiteRegion = "centralUS"
	NetworkaccessBranchSiteRegion_EastUS             NetworkaccessBranchSiteRegion = "eastUS"
	NetworkaccessBranchSiteRegion_EastUS2            NetworkaccessBranchSiteRegion = "eastUS2"
	NetworkaccessBranchSiteRegion_FranceCentral      NetworkaccessBranchSiteRegion = "franceCentral"
	NetworkaccessBranchSiteRegion_GermanyWestCentral NetworkaccessBranchSiteRegion = "germanyWestCentral"
	NetworkaccessBranchSiteRegion_JapanEast          NetworkaccessBranchSiteRegion = "japanEast"
	NetworkaccessBranchSiteRegion_JapanWest          NetworkaccessBranchSiteRegion = "japanWest"
	NetworkaccessBranchSiteRegion_KoreaCentral       NetworkaccessBranchSiteRegion = "koreaCentral"
	NetworkaccessBranchSiteRegion_NorthCentralUS     NetworkaccessBranchSiteRegion = "northCentralUS"
	NetworkaccessBranchSiteRegion_NorthEurope        NetworkaccessBranchSiteRegion = "northEurope"
	NetworkaccessBranchSiteRegion_PolandCentral      NetworkaccessBranchSiteRegion = "polandCentral"
	NetworkaccessBranchSiteRegion_SouthAfricaNorth   NetworkaccessBranchSiteRegion = "southAfricaNorth"
	NetworkaccessBranchSiteRegion_SouthAfricaWest    NetworkaccessBranchSiteRegion = "southAfricaWest"
	NetworkaccessBranchSiteRegion_SouthCentralUS     NetworkaccessBranchSiteRegion = "southCentralUS"
	NetworkaccessBranchSiteRegion_SouthEastAsia      NetworkaccessBranchSiteRegion = "southEastAsia"
	NetworkaccessBranchSiteRegion_SouthIndia         NetworkaccessBranchSiteRegion = "southIndia"
	NetworkaccessBranchSiteRegion_SwedenCentral      NetworkaccessBranchSiteRegion = "swedenCentral"
	NetworkaccessBranchSiteRegion_SwitzerlandNorth   NetworkaccessBranchSiteRegion = "switzerlandNorth"
	NetworkaccessBranchSiteRegion_UaeNorth           NetworkaccessBranchSiteRegion = "uaeNorth"
	NetworkaccessBranchSiteRegion_UkSouth            NetworkaccessBranchSiteRegion = "ukSouth"
	NetworkaccessBranchSiteRegion_WestCentralUS      NetworkaccessBranchSiteRegion = "westCentralUS"
	NetworkaccessBranchSiteRegion_WestEurope         NetworkaccessBranchSiteRegion = "westEurope"
	NetworkaccessBranchSiteRegion_WestUS             NetworkaccessBranchSiteRegion = "westUS"
	NetworkaccessBranchSiteRegion_WestUS2            NetworkaccessBranchSiteRegion = "westUS2"
	NetworkaccessBranchSiteRegion_WestUS3            NetworkaccessBranchSiteRegion = "westUS3"
)

func PossibleValuesForNetworkaccessBranchSiteRegion() []string {
	return []string{
		string(NetworkaccessBranchSiteRegion_AustraliaEast),
		string(NetworkaccessBranchSiteRegion_AustraliaSouthEast),
		string(NetworkaccessBranchSiteRegion_BrazilSouth),
		string(NetworkaccessBranchSiteRegion_CanadaCentral),
		string(NetworkaccessBranchSiteRegion_CanadaEast),
		string(NetworkaccessBranchSiteRegion_CentralIndia),
		string(NetworkaccessBranchSiteRegion_CentralUS),
		string(NetworkaccessBranchSiteRegion_EastUS),
		string(NetworkaccessBranchSiteRegion_EastUS2),
		string(NetworkaccessBranchSiteRegion_FranceCentral),
		string(NetworkaccessBranchSiteRegion_GermanyWestCentral),
		string(NetworkaccessBranchSiteRegion_JapanEast),
		string(NetworkaccessBranchSiteRegion_JapanWest),
		string(NetworkaccessBranchSiteRegion_KoreaCentral),
		string(NetworkaccessBranchSiteRegion_NorthCentralUS),
		string(NetworkaccessBranchSiteRegion_NorthEurope),
		string(NetworkaccessBranchSiteRegion_PolandCentral),
		string(NetworkaccessBranchSiteRegion_SouthAfricaNorth),
		string(NetworkaccessBranchSiteRegion_SouthAfricaWest),
		string(NetworkaccessBranchSiteRegion_SouthCentralUS),
		string(NetworkaccessBranchSiteRegion_SouthEastAsia),
		string(NetworkaccessBranchSiteRegion_SouthIndia),
		string(NetworkaccessBranchSiteRegion_SwedenCentral),
		string(NetworkaccessBranchSiteRegion_SwitzerlandNorth),
		string(NetworkaccessBranchSiteRegion_UaeNorth),
		string(NetworkaccessBranchSiteRegion_UkSouth),
		string(NetworkaccessBranchSiteRegion_WestCentralUS),
		string(NetworkaccessBranchSiteRegion_WestEurope),
		string(NetworkaccessBranchSiteRegion_WestUS),
		string(NetworkaccessBranchSiteRegion_WestUS2),
		string(NetworkaccessBranchSiteRegion_WestUS3),
	}
}

func (s *NetworkaccessBranchSiteRegion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessBranchSiteRegion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessBranchSiteRegion(input string) (*NetworkaccessBranchSiteRegion, error) {
	vals := map[string]NetworkaccessBranchSiteRegion{
		"australiaeast":      NetworkaccessBranchSiteRegion_AustraliaEast,
		"australiasoutheast": NetworkaccessBranchSiteRegion_AustraliaSouthEast,
		"brazilsouth":        NetworkaccessBranchSiteRegion_BrazilSouth,
		"canadacentral":      NetworkaccessBranchSiteRegion_CanadaCentral,
		"canadaeast":         NetworkaccessBranchSiteRegion_CanadaEast,
		"centralindia":       NetworkaccessBranchSiteRegion_CentralIndia,
		"centralus":          NetworkaccessBranchSiteRegion_CentralUS,
		"eastus":             NetworkaccessBranchSiteRegion_EastUS,
		"eastus2":            NetworkaccessBranchSiteRegion_EastUS2,
		"francecentral":      NetworkaccessBranchSiteRegion_FranceCentral,
		"germanywestcentral": NetworkaccessBranchSiteRegion_GermanyWestCentral,
		"japaneast":          NetworkaccessBranchSiteRegion_JapanEast,
		"japanwest":          NetworkaccessBranchSiteRegion_JapanWest,
		"koreacentral":       NetworkaccessBranchSiteRegion_KoreaCentral,
		"northcentralus":     NetworkaccessBranchSiteRegion_NorthCentralUS,
		"northeurope":        NetworkaccessBranchSiteRegion_NorthEurope,
		"polandcentral":      NetworkaccessBranchSiteRegion_PolandCentral,
		"southafricanorth":   NetworkaccessBranchSiteRegion_SouthAfricaNorth,
		"southafricawest":    NetworkaccessBranchSiteRegion_SouthAfricaWest,
		"southcentralus":     NetworkaccessBranchSiteRegion_SouthCentralUS,
		"southeastasia":      NetworkaccessBranchSiteRegion_SouthEastAsia,
		"southindia":         NetworkaccessBranchSiteRegion_SouthIndia,
		"swedencentral":      NetworkaccessBranchSiteRegion_SwedenCentral,
		"switzerlandnorth":   NetworkaccessBranchSiteRegion_SwitzerlandNorth,
		"uaenorth":           NetworkaccessBranchSiteRegion_UaeNorth,
		"uksouth":            NetworkaccessBranchSiteRegion_UkSouth,
		"westcentralus":      NetworkaccessBranchSiteRegion_WestCentralUS,
		"westeurope":         NetworkaccessBranchSiteRegion_WestEurope,
		"westus":             NetworkaccessBranchSiteRegion_WestUS,
		"westus2":            NetworkaccessBranchSiteRegion_WestUS2,
		"westus3":            NetworkaccessBranchSiteRegion_WestUS3,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessBranchSiteRegion(input)
	return &out, nil
}
