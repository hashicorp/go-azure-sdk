package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemPhoneType string

const (
	ItemPhoneTypeassistant   ItemPhoneType = "Assistant"
	ItemPhoneTypebusiness    ItemPhoneType = "Business"
	ItemPhoneTypebusinessFax ItemPhoneType = "BusinessFax"
	ItemPhoneTypehome        ItemPhoneType = "Home"
	ItemPhoneTypehomeFax     ItemPhoneType = "HomeFax"
	ItemPhoneTypemobile      ItemPhoneType = "Mobile"
	ItemPhoneTypeother       ItemPhoneType = "Other"
	ItemPhoneTypeotherFax    ItemPhoneType = "OtherFax"
	ItemPhoneTypepager       ItemPhoneType = "Pager"
	ItemPhoneTyperadio       ItemPhoneType = "Radio"
)

func PossibleValuesForItemPhoneType() []string {
	return []string{
		string(ItemPhoneTypeassistant),
		string(ItemPhoneTypebusiness),
		string(ItemPhoneTypebusinessFax),
		string(ItemPhoneTypehome),
		string(ItemPhoneTypehomeFax),
		string(ItemPhoneTypemobile),
		string(ItemPhoneTypeother),
		string(ItemPhoneTypeotherFax),
		string(ItemPhoneTypepager),
		string(ItemPhoneTyperadio),
	}
}

func parseItemPhoneType(input string) (*ItemPhoneType, error) {
	vals := map[string]ItemPhoneType{
		"assistant":   ItemPhoneTypeassistant,
		"business":    ItemPhoneTypebusiness,
		"businessfax": ItemPhoneTypebusinessFax,
		"home":        ItemPhoneTypehome,
		"homefax":     ItemPhoneTypehomeFax,
		"mobile":      ItemPhoneTypemobile,
		"other":       ItemPhoneTypeother,
		"otherfax":    ItemPhoneTypeotherFax,
		"pager":       ItemPhoneTypepager,
		"radio":       ItemPhoneTyperadio,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemPhoneType(input)
	return &out, nil
}
