package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessLocalConnectivityConfigurationRegion string

const (
	NetworkaccessLocalConnectivityConfigurationRegion_AustraliaEast      NetworkaccessLocalConnectivityConfigurationRegion = "australiaEast"
	NetworkaccessLocalConnectivityConfigurationRegion_AustraliaSouthEast NetworkaccessLocalConnectivityConfigurationRegion = "australiaSouthEast"
	NetworkaccessLocalConnectivityConfigurationRegion_BrazilSouth        NetworkaccessLocalConnectivityConfigurationRegion = "brazilSouth"
	NetworkaccessLocalConnectivityConfigurationRegion_CanadaCentral      NetworkaccessLocalConnectivityConfigurationRegion = "canadaCentral"
	NetworkaccessLocalConnectivityConfigurationRegion_CanadaEast         NetworkaccessLocalConnectivityConfigurationRegion = "canadaEast"
	NetworkaccessLocalConnectivityConfigurationRegion_CentralIndia       NetworkaccessLocalConnectivityConfigurationRegion = "centralIndia"
	NetworkaccessLocalConnectivityConfigurationRegion_CentralUS          NetworkaccessLocalConnectivityConfigurationRegion = "centralUS"
	NetworkaccessLocalConnectivityConfigurationRegion_EastUS             NetworkaccessLocalConnectivityConfigurationRegion = "eastUS"
	NetworkaccessLocalConnectivityConfigurationRegion_EastUS2            NetworkaccessLocalConnectivityConfigurationRegion = "eastUS2"
	NetworkaccessLocalConnectivityConfigurationRegion_FranceCentral      NetworkaccessLocalConnectivityConfigurationRegion = "franceCentral"
	NetworkaccessLocalConnectivityConfigurationRegion_GermanyWestCentral NetworkaccessLocalConnectivityConfigurationRegion = "germanyWestCentral"
	NetworkaccessLocalConnectivityConfigurationRegion_JapanEast          NetworkaccessLocalConnectivityConfigurationRegion = "japanEast"
	NetworkaccessLocalConnectivityConfigurationRegion_JapanWest          NetworkaccessLocalConnectivityConfigurationRegion = "japanWest"
	NetworkaccessLocalConnectivityConfigurationRegion_KoreaCentral       NetworkaccessLocalConnectivityConfigurationRegion = "koreaCentral"
	NetworkaccessLocalConnectivityConfigurationRegion_NorthCentralUS     NetworkaccessLocalConnectivityConfigurationRegion = "northCentralUS"
	NetworkaccessLocalConnectivityConfigurationRegion_NorthEurope        NetworkaccessLocalConnectivityConfigurationRegion = "northEurope"
	NetworkaccessLocalConnectivityConfigurationRegion_PolandCentral      NetworkaccessLocalConnectivityConfigurationRegion = "polandCentral"
	NetworkaccessLocalConnectivityConfigurationRegion_SouthAfricaNorth   NetworkaccessLocalConnectivityConfigurationRegion = "southAfricaNorth"
	NetworkaccessLocalConnectivityConfigurationRegion_SouthAfricaWest    NetworkaccessLocalConnectivityConfigurationRegion = "southAfricaWest"
	NetworkaccessLocalConnectivityConfigurationRegion_SouthCentralUS     NetworkaccessLocalConnectivityConfigurationRegion = "southCentralUS"
	NetworkaccessLocalConnectivityConfigurationRegion_SouthEastAsia      NetworkaccessLocalConnectivityConfigurationRegion = "southEastAsia"
	NetworkaccessLocalConnectivityConfigurationRegion_SouthIndia         NetworkaccessLocalConnectivityConfigurationRegion = "southIndia"
	NetworkaccessLocalConnectivityConfigurationRegion_SwedenCentral      NetworkaccessLocalConnectivityConfigurationRegion = "swedenCentral"
	NetworkaccessLocalConnectivityConfigurationRegion_SwitzerlandNorth   NetworkaccessLocalConnectivityConfigurationRegion = "switzerlandNorth"
	NetworkaccessLocalConnectivityConfigurationRegion_UaeNorth           NetworkaccessLocalConnectivityConfigurationRegion = "uaeNorth"
	NetworkaccessLocalConnectivityConfigurationRegion_UkSouth            NetworkaccessLocalConnectivityConfigurationRegion = "ukSouth"
	NetworkaccessLocalConnectivityConfigurationRegion_WestCentralUS      NetworkaccessLocalConnectivityConfigurationRegion = "westCentralUS"
	NetworkaccessLocalConnectivityConfigurationRegion_WestEurope         NetworkaccessLocalConnectivityConfigurationRegion = "westEurope"
	NetworkaccessLocalConnectivityConfigurationRegion_WestUS             NetworkaccessLocalConnectivityConfigurationRegion = "westUS"
	NetworkaccessLocalConnectivityConfigurationRegion_WestUS2            NetworkaccessLocalConnectivityConfigurationRegion = "westUS2"
	NetworkaccessLocalConnectivityConfigurationRegion_WestUS3            NetworkaccessLocalConnectivityConfigurationRegion = "westUS3"
)

func PossibleValuesForNetworkaccessLocalConnectivityConfigurationRegion() []string {
	return []string{
		string(NetworkaccessLocalConnectivityConfigurationRegion_AustraliaEast),
		string(NetworkaccessLocalConnectivityConfigurationRegion_AustraliaSouthEast),
		string(NetworkaccessLocalConnectivityConfigurationRegion_BrazilSouth),
		string(NetworkaccessLocalConnectivityConfigurationRegion_CanadaCentral),
		string(NetworkaccessLocalConnectivityConfigurationRegion_CanadaEast),
		string(NetworkaccessLocalConnectivityConfigurationRegion_CentralIndia),
		string(NetworkaccessLocalConnectivityConfigurationRegion_CentralUS),
		string(NetworkaccessLocalConnectivityConfigurationRegion_EastUS),
		string(NetworkaccessLocalConnectivityConfigurationRegion_EastUS2),
		string(NetworkaccessLocalConnectivityConfigurationRegion_FranceCentral),
		string(NetworkaccessLocalConnectivityConfigurationRegion_GermanyWestCentral),
		string(NetworkaccessLocalConnectivityConfigurationRegion_JapanEast),
		string(NetworkaccessLocalConnectivityConfigurationRegion_JapanWest),
		string(NetworkaccessLocalConnectivityConfigurationRegion_KoreaCentral),
		string(NetworkaccessLocalConnectivityConfigurationRegion_NorthCentralUS),
		string(NetworkaccessLocalConnectivityConfigurationRegion_NorthEurope),
		string(NetworkaccessLocalConnectivityConfigurationRegion_PolandCentral),
		string(NetworkaccessLocalConnectivityConfigurationRegion_SouthAfricaNorth),
		string(NetworkaccessLocalConnectivityConfigurationRegion_SouthAfricaWest),
		string(NetworkaccessLocalConnectivityConfigurationRegion_SouthCentralUS),
		string(NetworkaccessLocalConnectivityConfigurationRegion_SouthEastAsia),
		string(NetworkaccessLocalConnectivityConfigurationRegion_SouthIndia),
		string(NetworkaccessLocalConnectivityConfigurationRegion_SwedenCentral),
		string(NetworkaccessLocalConnectivityConfigurationRegion_SwitzerlandNorth),
		string(NetworkaccessLocalConnectivityConfigurationRegion_UaeNorth),
		string(NetworkaccessLocalConnectivityConfigurationRegion_UkSouth),
		string(NetworkaccessLocalConnectivityConfigurationRegion_WestCentralUS),
		string(NetworkaccessLocalConnectivityConfigurationRegion_WestEurope),
		string(NetworkaccessLocalConnectivityConfigurationRegion_WestUS),
		string(NetworkaccessLocalConnectivityConfigurationRegion_WestUS2),
		string(NetworkaccessLocalConnectivityConfigurationRegion_WestUS3),
	}
}

func (s *NetworkaccessLocalConnectivityConfigurationRegion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessLocalConnectivityConfigurationRegion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessLocalConnectivityConfigurationRegion(input string) (*NetworkaccessLocalConnectivityConfigurationRegion, error) {
	vals := map[string]NetworkaccessLocalConnectivityConfigurationRegion{
		"australiaeast":      NetworkaccessLocalConnectivityConfigurationRegion_AustraliaEast,
		"australiasoutheast": NetworkaccessLocalConnectivityConfigurationRegion_AustraliaSouthEast,
		"brazilsouth":        NetworkaccessLocalConnectivityConfigurationRegion_BrazilSouth,
		"canadacentral":      NetworkaccessLocalConnectivityConfigurationRegion_CanadaCentral,
		"canadaeast":         NetworkaccessLocalConnectivityConfigurationRegion_CanadaEast,
		"centralindia":       NetworkaccessLocalConnectivityConfigurationRegion_CentralIndia,
		"centralus":          NetworkaccessLocalConnectivityConfigurationRegion_CentralUS,
		"eastus":             NetworkaccessLocalConnectivityConfigurationRegion_EastUS,
		"eastus2":            NetworkaccessLocalConnectivityConfigurationRegion_EastUS2,
		"francecentral":      NetworkaccessLocalConnectivityConfigurationRegion_FranceCentral,
		"germanywestcentral": NetworkaccessLocalConnectivityConfigurationRegion_GermanyWestCentral,
		"japaneast":          NetworkaccessLocalConnectivityConfigurationRegion_JapanEast,
		"japanwest":          NetworkaccessLocalConnectivityConfigurationRegion_JapanWest,
		"koreacentral":       NetworkaccessLocalConnectivityConfigurationRegion_KoreaCentral,
		"northcentralus":     NetworkaccessLocalConnectivityConfigurationRegion_NorthCentralUS,
		"northeurope":        NetworkaccessLocalConnectivityConfigurationRegion_NorthEurope,
		"polandcentral":      NetworkaccessLocalConnectivityConfigurationRegion_PolandCentral,
		"southafricanorth":   NetworkaccessLocalConnectivityConfigurationRegion_SouthAfricaNorth,
		"southafricawest":    NetworkaccessLocalConnectivityConfigurationRegion_SouthAfricaWest,
		"southcentralus":     NetworkaccessLocalConnectivityConfigurationRegion_SouthCentralUS,
		"southeastasia":      NetworkaccessLocalConnectivityConfigurationRegion_SouthEastAsia,
		"southindia":         NetworkaccessLocalConnectivityConfigurationRegion_SouthIndia,
		"swedencentral":      NetworkaccessLocalConnectivityConfigurationRegion_SwedenCentral,
		"switzerlandnorth":   NetworkaccessLocalConnectivityConfigurationRegion_SwitzerlandNorth,
		"uaenorth":           NetworkaccessLocalConnectivityConfigurationRegion_UaeNorth,
		"uksouth":            NetworkaccessLocalConnectivityConfigurationRegion_UkSouth,
		"westcentralus":      NetworkaccessLocalConnectivityConfigurationRegion_WestCentralUS,
		"westeurope":         NetworkaccessLocalConnectivityConfigurationRegion_WestEurope,
		"westus":             NetworkaccessLocalConnectivityConfigurationRegion_WestUS,
		"westus2":            NetworkaccessLocalConnectivityConfigurationRegion_WestUS2,
		"westus3":            NetworkaccessLocalConnectivityConfigurationRegion_WestUS3,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessLocalConnectivityConfigurationRegion(input)
	return &out, nil
}
