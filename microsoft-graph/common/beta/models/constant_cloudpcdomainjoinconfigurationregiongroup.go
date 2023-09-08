package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDomainJoinConfigurationRegionGroup string

const (
	CloudPCDomainJoinConfigurationRegionGroupasia            CloudPCDomainJoinConfigurationRegionGroup = "Asia"
	CloudPCDomainJoinConfigurationRegionGroupaustralia       CloudPCDomainJoinConfigurationRegionGroup = "Australia"
	CloudPCDomainJoinConfigurationRegionGroupcanada          CloudPCDomainJoinConfigurationRegionGroup = "Canada"
	CloudPCDomainJoinConfigurationRegionGroupdefault         CloudPCDomainJoinConfigurationRegionGroup = "Default"
	CloudPCDomainJoinConfigurationRegionGroupeuap            CloudPCDomainJoinConfigurationRegionGroup = "Euap"
	CloudPCDomainJoinConfigurationRegionGroupeuropeUnion     CloudPCDomainJoinConfigurationRegionGroup = "EuropeUnion"
	CloudPCDomainJoinConfigurationRegionGroupfrance          CloudPCDomainJoinConfigurationRegionGroup = "France"
	CloudPCDomainJoinConfigurationRegionGroupgermany         CloudPCDomainJoinConfigurationRegionGroup = "Germany"
	CloudPCDomainJoinConfigurationRegionGroupindia           CloudPCDomainJoinConfigurationRegionGroup = "India"
	CloudPCDomainJoinConfigurationRegionGroupjapan           CloudPCDomainJoinConfigurationRegionGroup = "Japan"
	CloudPCDomainJoinConfigurationRegionGroupnorway          CloudPCDomainJoinConfigurationRegionGroup = "Norway"
	CloudPCDomainJoinConfigurationRegionGroupsouthAmerica    CloudPCDomainJoinConfigurationRegionGroup = "SouthAmerica"
	CloudPCDomainJoinConfigurationRegionGroupsouthKorea      CloudPCDomainJoinConfigurationRegionGroup = "SouthKorea"
	CloudPCDomainJoinConfigurationRegionGroupswitzerland     CloudPCDomainJoinConfigurationRegionGroup = "Switzerland"
	CloudPCDomainJoinConfigurationRegionGroupunitedKingdom   CloudPCDomainJoinConfigurationRegionGroup = "UnitedKingdom"
	CloudPCDomainJoinConfigurationRegionGroupusCentral       CloudPCDomainJoinConfigurationRegionGroup = "UsCentral"
	CloudPCDomainJoinConfigurationRegionGroupusEast          CloudPCDomainJoinConfigurationRegionGroup = "UsEast"
	CloudPCDomainJoinConfigurationRegionGroupusGovernment    CloudPCDomainJoinConfigurationRegionGroup = "UsGovernment"
	CloudPCDomainJoinConfigurationRegionGroupusGovernmentDOD CloudPCDomainJoinConfigurationRegionGroup = "UsGovernmentDOD"
	CloudPCDomainJoinConfigurationRegionGroupusWest          CloudPCDomainJoinConfigurationRegionGroup = "UsWest"
)

func PossibleValuesForCloudPCDomainJoinConfigurationRegionGroup() []string {
	return []string{
		string(CloudPCDomainJoinConfigurationRegionGroupasia),
		string(CloudPCDomainJoinConfigurationRegionGroupaustralia),
		string(CloudPCDomainJoinConfigurationRegionGroupcanada),
		string(CloudPCDomainJoinConfigurationRegionGroupdefault),
		string(CloudPCDomainJoinConfigurationRegionGroupeuap),
		string(CloudPCDomainJoinConfigurationRegionGroupeuropeUnion),
		string(CloudPCDomainJoinConfigurationRegionGroupfrance),
		string(CloudPCDomainJoinConfigurationRegionGroupgermany),
		string(CloudPCDomainJoinConfigurationRegionGroupindia),
		string(CloudPCDomainJoinConfigurationRegionGroupjapan),
		string(CloudPCDomainJoinConfigurationRegionGroupnorway),
		string(CloudPCDomainJoinConfigurationRegionGroupsouthAmerica),
		string(CloudPCDomainJoinConfigurationRegionGroupsouthKorea),
		string(CloudPCDomainJoinConfigurationRegionGroupswitzerland),
		string(CloudPCDomainJoinConfigurationRegionGroupunitedKingdom),
		string(CloudPCDomainJoinConfigurationRegionGroupusCentral),
		string(CloudPCDomainJoinConfigurationRegionGroupusEast),
		string(CloudPCDomainJoinConfigurationRegionGroupusGovernment),
		string(CloudPCDomainJoinConfigurationRegionGroupusGovernmentDOD),
		string(CloudPCDomainJoinConfigurationRegionGroupusWest),
	}
}

func parseCloudPCDomainJoinConfigurationRegionGroup(input string) (*CloudPCDomainJoinConfigurationRegionGroup, error) {
	vals := map[string]CloudPCDomainJoinConfigurationRegionGroup{
		"asia":            CloudPCDomainJoinConfigurationRegionGroupasia,
		"australia":       CloudPCDomainJoinConfigurationRegionGroupaustralia,
		"canada":          CloudPCDomainJoinConfigurationRegionGroupcanada,
		"default":         CloudPCDomainJoinConfigurationRegionGroupdefault,
		"euap":            CloudPCDomainJoinConfigurationRegionGroupeuap,
		"europeunion":     CloudPCDomainJoinConfigurationRegionGroupeuropeUnion,
		"france":          CloudPCDomainJoinConfigurationRegionGroupfrance,
		"germany":         CloudPCDomainJoinConfigurationRegionGroupgermany,
		"india":           CloudPCDomainJoinConfigurationRegionGroupindia,
		"japan":           CloudPCDomainJoinConfigurationRegionGroupjapan,
		"norway":          CloudPCDomainJoinConfigurationRegionGroupnorway,
		"southamerica":    CloudPCDomainJoinConfigurationRegionGroupsouthAmerica,
		"southkorea":      CloudPCDomainJoinConfigurationRegionGroupsouthKorea,
		"switzerland":     CloudPCDomainJoinConfigurationRegionGroupswitzerland,
		"unitedkingdom":   CloudPCDomainJoinConfigurationRegionGroupunitedKingdom,
		"uscentral":       CloudPCDomainJoinConfigurationRegionGroupusCentral,
		"useast":          CloudPCDomainJoinConfigurationRegionGroupusEast,
		"usgovernment":    CloudPCDomainJoinConfigurationRegionGroupusGovernment,
		"usgovernmentdod": CloudPCDomainJoinConfigurationRegionGroupusGovernmentDOD,
		"uswest":          CloudPCDomainJoinConfigurationRegionGroupusWest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDomainJoinConfigurationRegionGroup(input)
	return &out, nil
}
