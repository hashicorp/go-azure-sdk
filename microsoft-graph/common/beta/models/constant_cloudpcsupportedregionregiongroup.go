package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSupportedRegionRegionGroup string

const (
	CloudPCSupportedRegionRegionGroupasia            CloudPCSupportedRegionRegionGroup = "Asia"
	CloudPCSupportedRegionRegionGroupaustralia       CloudPCSupportedRegionRegionGroup = "Australia"
	CloudPCSupportedRegionRegionGroupcanada          CloudPCSupportedRegionRegionGroup = "Canada"
	CloudPCSupportedRegionRegionGroupdefault         CloudPCSupportedRegionRegionGroup = "Default"
	CloudPCSupportedRegionRegionGroupeuap            CloudPCSupportedRegionRegionGroup = "Euap"
	CloudPCSupportedRegionRegionGroupeuropeUnion     CloudPCSupportedRegionRegionGroup = "EuropeUnion"
	CloudPCSupportedRegionRegionGroupfrance          CloudPCSupportedRegionRegionGroup = "France"
	CloudPCSupportedRegionRegionGroupgermany         CloudPCSupportedRegionRegionGroup = "Germany"
	CloudPCSupportedRegionRegionGroupindia           CloudPCSupportedRegionRegionGroup = "India"
	CloudPCSupportedRegionRegionGroupjapan           CloudPCSupportedRegionRegionGroup = "Japan"
	CloudPCSupportedRegionRegionGroupnorway          CloudPCSupportedRegionRegionGroup = "Norway"
	CloudPCSupportedRegionRegionGroupsouthAmerica    CloudPCSupportedRegionRegionGroup = "SouthAmerica"
	CloudPCSupportedRegionRegionGroupsouthKorea      CloudPCSupportedRegionRegionGroup = "SouthKorea"
	CloudPCSupportedRegionRegionGroupswitzerland     CloudPCSupportedRegionRegionGroup = "Switzerland"
	CloudPCSupportedRegionRegionGroupunitedKingdom   CloudPCSupportedRegionRegionGroup = "UnitedKingdom"
	CloudPCSupportedRegionRegionGroupusCentral       CloudPCSupportedRegionRegionGroup = "UsCentral"
	CloudPCSupportedRegionRegionGroupusEast          CloudPCSupportedRegionRegionGroup = "UsEast"
	CloudPCSupportedRegionRegionGroupusGovernment    CloudPCSupportedRegionRegionGroup = "UsGovernment"
	CloudPCSupportedRegionRegionGroupusGovernmentDOD CloudPCSupportedRegionRegionGroup = "UsGovernmentDOD"
	CloudPCSupportedRegionRegionGroupusWest          CloudPCSupportedRegionRegionGroup = "UsWest"
)

func PossibleValuesForCloudPCSupportedRegionRegionGroup() []string {
	return []string{
		string(CloudPCSupportedRegionRegionGroupasia),
		string(CloudPCSupportedRegionRegionGroupaustralia),
		string(CloudPCSupportedRegionRegionGroupcanada),
		string(CloudPCSupportedRegionRegionGroupdefault),
		string(CloudPCSupportedRegionRegionGroupeuap),
		string(CloudPCSupportedRegionRegionGroupeuropeUnion),
		string(CloudPCSupportedRegionRegionGroupfrance),
		string(CloudPCSupportedRegionRegionGroupgermany),
		string(CloudPCSupportedRegionRegionGroupindia),
		string(CloudPCSupportedRegionRegionGroupjapan),
		string(CloudPCSupportedRegionRegionGroupnorway),
		string(CloudPCSupportedRegionRegionGroupsouthAmerica),
		string(CloudPCSupportedRegionRegionGroupsouthKorea),
		string(CloudPCSupportedRegionRegionGroupswitzerland),
		string(CloudPCSupportedRegionRegionGroupunitedKingdom),
		string(CloudPCSupportedRegionRegionGroupusCentral),
		string(CloudPCSupportedRegionRegionGroupusEast),
		string(CloudPCSupportedRegionRegionGroupusGovernment),
		string(CloudPCSupportedRegionRegionGroupusGovernmentDOD),
		string(CloudPCSupportedRegionRegionGroupusWest),
	}
}

func parseCloudPCSupportedRegionRegionGroup(input string) (*CloudPCSupportedRegionRegionGroup, error) {
	vals := map[string]CloudPCSupportedRegionRegionGroup{
		"asia":            CloudPCSupportedRegionRegionGroupasia,
		"australia":       CloudPCSupportedRegionRegionGroupaustralia,
		"canada":          CloudPCSupportedRegionRegionGroupcanada,
		"default":         CloudPCSupportedRegionRegionGroupdefault,
		"euap":            CloudPCSupportedRegionRegionGroupeuap,
		"europeunion":     CloudPCSupportedRegionRegionGroupeuropeUnion,
		"france":          CloudPCSupportedRegionRegionGroupfrance,
		"germany":         CloudPCSupportedRegionRegionGroupgermany,
		"india":           CloudPCSupportedRegionRegionGroupindia,
		"japan":           CloudPCSupportedRegionRegionGroupjapan,
		"norway":          CloudPCSupportedRegionRegionGroupnorway,
		"southamerica":    CloudPCSupportedRegionRegionGroupsouthAmerica,
		"southkorea":      CloudPCSupportedRegionRegionGroupsouthKorea,
		"switzerland":     CloudPCSupportedRegionRegionGroupswitzerland,
		"unitedkingdom":   CloudPCSupportedRegionRegionGroupunitedKingdom,
		"uscentral":       CloudPCSupportedRegionRegionGroupusCentral,
		"useast":          CloudPCSupportedRegionRegionGroupusEast,
		"usgovernment":    CloudPCSupportedRegionRegionGroupusGovernment,
		"usgovernmentdod": CloudPCSupportedRegionRegionGroupusGovernmentDOD,
		"uswest":          CloudPCSupportedRegionRegionGroupusWest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCSupportedRegionRegionGroup(input)
	return &out, nil
}
