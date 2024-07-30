package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSupportedRegionRegionGroup string

const (
	CloudPCSupportedRegionRegionGroup_Asia            CloudPCSupportedRegionRegionGroup = "asia"
	CloudPCSupportedRegionRegionGroup_Australia       CloudPCSupportedRegionRegionGroup = "australia"
	CloudPCSupportedRegionRegionGroup_Canada          CloudPCSupportedRegionRegionGroup = "canada"
	CloudPCSupportedRegionRegionGroup_Default         CloudPCSupportedRegionRegionGroup = "default"
	CloudPCSupportedRegionRegionGroup_Euap            CloudPCSupportedRegionRegionGroup = "euap"
	CloudPCSupportedRegionRegionGroup_EuropeUnion     CloudPCSupportedRegionRegionGroup = "europeUnion"
	CloudPCSupportedRegionRegionGroup_France          CloudPCSupportedRegionRegionGroup = "france"
	CloudPCSupportedRegionRegionGroup_Germany         CloudPCSupportedRegionRegionGroup = "germany"
	CloudPCSupportedRegionRegionGroup_India           CloudPCSupportedRegionRegionGroup = "india"
	CloudPCSupportedRegionRegionGroup_Japan           CloudPCSupportedRegionRegionGroup = "japan"
	CloudPCSupportedRegionRegionGroup_Norway          CloudPCSupportedRegionRegionGroup = "norway"
	CloudPCSupportedRegionRegionGroup_SouthAmerica    CloudPCSupportedRegionRegionGroup = "southAmerica"
	CloudPCSupportedRegionRegionGroup_SouthKorea      CloudPCSupportedRegionRegionGroup = "southKorea"
	CloudPCSupportedRegionRegionGroup_Switzerland     CloudPCSupportedRegionRegionGroup = "switzerland"
	CloudPCSupportedRegionRegionGroup_UnitedKingdom   CloudPCSupportedRegionRegionGroup = "unitedKingdom"
	CloudPCSupportedRegionRegionGroup_UsCentral       CloudPCSupportedRegionRegionGroup = "usCentral"
	CloudPCSupportedRegionRegionGroup_UsEast          CloudPCSupportedRegionRegionGroup = "usEast"
	CloudPCSupportedRegionRegionGroup_UsGovernment    CloudPCSupportedRegionRegionGroup = "usGovernment"
	CloudPCSupportedRegionRegionGroup_UsGovernmentDOD CloudPCSupportedRegionRegionGroup = "usGovernmentDOD"
	CloudPCSupportedRegionRegionGroup_UsWest          CloudPCSupportedRegionRegionGroup = "usWest"
)

func PossibleValuesForCloudPCSupportedRegionRegionGroup() []string {
	return []string{
		string(CloudPCSupportedRegionRegionGroup_Asia),
		string(CloudPCSupportedRegionRegionGroup_Australia),
		string(CloudPCSupportedRegionRegionGroup_Canada),
		string(CloudPCSupportedRegionRegionGroup_Default),
		string(CloudPCSupportedRegionRegionGroup_Euap),
		string(CloudPCSupportedRegionRegionGroup_EuropeUnion),
		string(CloudPCSupportedRegionRegionGroup_France),
		string(CloudPCSupportedRegionRegionGroup_Germany),
		string(CloudPCSupportedRegionRegionGroup_India),
		string(CloudPCSupportedRegionRegionGroup_Japan),
		string(CloudPCSupportedRegionRegionGroup_Norway),
		string(CloudPCSupportedRegionRegionGroup_SouthAmerica),
		string(CloudPCSupportedRegionRegionGroup_SouthKorea),
		string(CloudPCSupportedRegionRegionGroup_Switzerland),
		string(CloudPCSupportedRegionRegionGroup_UnitedKingdom),
		string(CloudPCSupportedRegionRegionGroup_UsCentral),
		string(CloudPCSupportedRegionRegionGroup_UsEast),
		string(CloudPCSupportedRegionRegionGroup_UsGovernment),
		string(CloudPCSupportedRegionRegionGroup_UsGovernmentDOD),
		string(CloudPCSupportedRegionRegionGroup_UsWest),
	}
}

func (s *CloudPCSupportedRegionRegionGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCSupportedRegionRegionGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCSupportedRegionRegionGroup(input string) (*CloudPCSupportedRegionRegionGroup, error) {
	vals := map[string]CloudPCSupportedRegionRegionGroup{
		"asia":            CloudPCSupportedRegionRegionGroup_Asia,
		"australia":       CloudPCSupportedRegionRegionGroup_Australia,
		"canada":          CloudPCSupportedRegionRegionGroup_Canada,
		"default":         CloudPCSupportedRegionRegionGroup_Default,
		"euap":            CloudPCSupportedRegionRegionGroup_Euap,
		"europeunion":     CloudPCSupportedRegionRegionGroup_EuropeUnion,
		"france":          CloudPCSupportedRegionRegionGroup_France,
		"germany":         CloudPCSupportedRegionRegionGroup_Germany,
		"india":           CloudPCSupportedRegionRegionGroup_India,
		"japan":           CloudPCSupportedRegionRegionGroup_Japan,
		"norway":          CloudPCSupportedRegionRegionGroup_Norway,
		"southamerica":    CloudPCSupportedRegionRegionGroup_SouthAmerica,
		"southkorea":      CloudPCSupportedRegionRegionGroup_SouthKorea,
		"switzerland":     CloudPCSupportedRegionRegionGroup_Switzerland,
		"unitedkingdom":   CloudPCSupportedRegionRegionGroup_UnitedKingdom,
		"uscentral":       CloudPCSupportedRegionRegionGroup_UsCentral,
		"useast":          CloudPCSupportedRegionRegionGroup_UsEast,
		"usgovernment":    CloudPCSupportedRegionRegionGroup_UsGovernment,
		"usgovernmentdod": CloudPCSupportedRegionRegionGroup_UsGovernmentDOD,
		"uswest":          CloudPCSupportedRegionRegionGroup_UsWest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCSupportedRegionRegionGroup(input)
	return &out, nil
}
