package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDomainJoinConfigurationRegionGroup string

const (
	CloudPCDomainJoinConfigurationRegionGroup_Asia            CloudPCDomainJoinConfigurationRegionGroup = "asia"
	CloudPCDomainJoinConfigurationRegionGroup_Australia       CloudPCDomainJoinConfigurationRegionGroup = "australia"
	CloudPCDomainJoinConfigurationRegionGroup_Canada          CloudPCDomainJoinConfigurationRegionGroup = "canada"
	CloudPCDomainJoinConfigurationRegionGroup_Default         CloudPCDomainJoinConfigurationRegionGroup = "default"
	CloudPCDomainJoinConfigurationRegionGroup_Euap            CloudPCDomainJoinConfigurationRegionGroup = "euap"
	CloudPCDomainJoinConfigurationRegionGroup_EuropeUnion     CloudPCDomainJoinConfigurationRegionGroup = "europeUnion"
	CloudPCDomainJoinConfigurationRegionGroup_France          CloudPCDomainJoinConfigurationRegionGroup = "france"
	CloudPCDomainJoinConfigurationRegionGroup_Germany         CloudPCDomainJoinConfigurationRegionGroup = "germany"
	CloudPCDomainJoinConfigurationRegionGroup_India           CloudPCDomainJoinConfigurationRegionGroup = "india"
	CloudPCDomainJoinConfigurationRegionGroup_Japan           CloudPCDomainJoinConfigurationRegionGroup = "japan"
	CloudPCDomainJoinConfigurationRegionGroup_Norway          CloudPCDomainJoinConfigurationRegionGroup = "norway"
	CloudPCDomainJoinConfigurationRegionGroup_SouthAmerica    CloudPCDomainJoinConfigurationRegionGroup = "southAmerica"
	CloudPCDomainJoinConfigurationRegionGroup_SouthKorea      CloudPCDomainJoinConfigurationRegionGroup = "southKorea"
	CloudPCDomainJoinConfigurationRegionGroup_Switzerland     CloudPCDomainJoinConfigurationRegionGroup = "switzerland"
	CloudPCDomainJoinConfigurationRegionGroup_UnitedKingdom   CloudPCDomainJoinConfigurationRegionGroup = "unitedKingdom"
	CloudPCDomainJoinConfigurationRegionGroup_UsCentral       CloudPCDomainJoinConfigurationRegionGroup = "usCentral"
	CloudPCDomainJoinConfigurationRegionGroup_UsEast          CloudPCDomainJoinConfigurationRegionGroup = "usEast"
	CloudPCDomainJoinConfigurationRegionGroup_UsGovernment    CloudPCDomainJoinConfigurationRegionGroup = "usGovernment"
	CloudPCDomainJoinConfigurationRegionGroup_UsGovernmentDOD CloudPCDomainJoinConfigurationRegionGroup = "usGovernmentDOD"
	CloudPCDomainJoinConfigurationRegionGroup_UsWest          CloudPCDomainJoinConfigurationRegionGroup = "usWest"
)

func PossibleValuesForCloudPCDomainJoinConfigurationRegionGroup() []string {
	return []string{
		string(CloudPCDomainJoinConfigurationRegionGroup_Asia),
		string(CloudPCDomainJoinConfigurationRegionGroup_Australia),
		string(CloudPCDomainJoinConfigurationRegionGroup_Canada),
		string(CloudPCDomainJoinConfigurationRegionGroup_Default),
		string(CloudPCDomainJoinConfigurationRegionGroup_Euap),
		string(CloudPCDomainJoinConfigurationRegionGroup_EuropeUnion),
		string(CloudPCDomainJoinConfigurationRegionGroup_France),
		string(CloudPCDomainJoinConfigurationRegionGroup_Germany),
		string(CloudPCDomainJoinConfigurationRegionGroup_India),
		string(CloudPCDomainJoinConfigurationRegionGroup_Japan),
		string(CloudPCDomainJoinConfigurationRegionGroup_Norway),
		string(CloudPCDomainJoinConfigurationRegionGroup_SouthAmerica),
		string(CloudPCDomainJoinConfigurationRegionGroup_SouthKorea),
		string(CloudPCDomainJoinConfigurationRegionGroup_Switzerland),
		string(CloudPCDomainJoinConfigurationRegionGroup_UnitedKingdom),
		string(CloudPCDomainJoinConfigurationRegionGroup_UsCentral),
		string(CloudPCDomainJoinConfigurationRegionGroup_UsEast),
		string(CloudPCDomainJoinConfigurationRegionGroup_UsGovernment),
		string(CloudPCDomainJoinConfigurationRegionGroup_UsGovernmentDOD),
		string(CloudPCDomainJoinConfigurationRegionGroup_UsWest),
	}
}

func (s *CloudPCDomainJoinConfigurationRegionGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCDomainJoinConfigurationRegionGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCDomainJoinConfigurationRegionGroup(input string) (*CloudPCDomainJoinConfigurationRegionGroup, error) {
	vals := map[string]CloudPCDomainJoinConfigurationRegionGroup{
		"asia":            CloudPCDomainJoinConfigurationRegionGroup_Asia,
		"australia":       CloudPCDomainJoinConfigurationRegionGroup_Australia,
		"canada":          CloudPCDomainJoinConfigurationRegionGroup_Canada,
		"default":         CloudPCDomainJoinConfigurationRegionGroup_Default,
		"euap":            CloudPCDomainJoinConfigurationRegionGroup_Euap,
		"europeunion":     CloudPCDomainJoinConfigurationRegionGroup_EuropeUnion,
		"france":          CloudPCDomainJoinConfigurationRegionGroup_France,
		"germany":         CloudPCDomainJoinConfigurationRegionGroup_Germany,
		"india":           CloudPCDomainJoinConfigurationRegionGroup_India,
		"japan":           CloudPCDomainJoinConfigurationRegionGroup_Japan,
		"norway":          CloudPCDomainJoinConfigurationRegionGroup_Norway,
		"southamerica":    CloudPCDomainJoinConfigurationRegionGroup_SouthAmerica,
		"southkorea":      CloudPCDomainJoinConfigurationRegionGroup_SouthKorea,
		"switzerland":     CloudPCDomainJoinConfigurationRegionGroup_Switzerland,
		"unitedkingdom":   CloudPCDomainJoinConfigurationRegionGroup_UnitedKingdom,
		"uscentral":       CloudPCDomainJoinConfigurationRegionGroup_UsCentral,
		"useast":          CloudPCDomainJoinConfigurationRegionGroup_UsEast,
		"usgovernment":    CloudPCDomainJoinConfigurationRegionGroup_UsGovernment,
		"usgovernmentdod": CloudPCDomainJoinConfigurationRegionGroup_UsGovernmentDOD,
		"uswest":          CloudPCDomainJoinConfigurationRegionGroup_UsWest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDomainJoinConfigurationRegionGroup(input)
	return &out, nil
}
