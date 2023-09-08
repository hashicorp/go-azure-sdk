package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PublishedResourcePublishingType string

const (
	PublishedResourcePublishingTypeapplicationProxy PublishedResourcePublishingType = "ApplicationProxy"
	PublishedResourcePublishingTypeauthentication   PublishedResourcePublishingType = "Authentication"
	PublishedResourcePublishingTypeexchangeOnline   PublishedResourcePublishingType = "ExchangeOnline"
	PublishedResourcePublishingTypeintunePfx        PublishedResourcePublishingType = "IntunePfx"
	PublishedResourcePublishingTypeoflineDomainJoin PublishedResourcePublishingType = "OflineDomainJoin"
	PublishedResourcePublishingTypeprovisioning     PublishedResourcePublishingType = "Provisioning"
)

func PossibleValuesForPublishedResourcePublishingType() []string {
	return []string{
		string(PublishedResourcePublishingTypeapplicationProxy),
		string(PublishedResourcePublishingTypeauthentication),
		string(PublishedResourcePublishingTypeexchangeOnline),
		string(PublishedResourcePublishingTypeintunePfx),
		string(PublishedResourcePublishingTypeoflineDomainJoin),
		string(PublishedResourcePublishingTypeprovisioning),
	}
}

func parsePublishedResourcePublishingType(input string) (*PublishedResourcePublishingType, error) {
	vals := map[string]PublishedResourcePublishingType{
		"applicationproxy": PublishedResourcePublishingTypeapplicationProxy,
		"authentication":   PublishedResourcePublishingTypeauthentication,
		"exchangeonline":   PublishedResourcePublishingTypeexchangeOnline,
		"intunepfx":        PublishedResourcePublishingTypeintunePfx,
		"oflinedomainjoin": PublishedResourcePublishingTypeoflineDomainJoin,
		"provisioning":     PublishedResourcePublishingTypeprovisioning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PublishedResourcePublishingType(input)
	return &out, nil
}
