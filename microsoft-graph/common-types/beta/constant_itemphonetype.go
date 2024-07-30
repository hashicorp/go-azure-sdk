package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemPhoneType string

const (
	ItemPhoneType_Assistant   ItemPhoneType = "assistant"
	ItemPhoneType_Business    ItemPhoneType = "business"
	ItemPhoneType_BusinessFax ItemPhoneType = "businessFax"
	ItemPhoneType_Home        ItemPhoneType = "home"
	ItemPhoneType_HomeFax     ItemPhoneType = "homeFax"
	ItemPhoneType_Mobile      ItemPhoneType = "mobile"
	ItemPhoneType_Other       ItemPhoneType = "other"
	ItemPhoneType_OtherFax    ItemPhoneType = "otherFax"
	ItemPhoneType_Pager       ItemPhoneType = "pager"
	ItemPhoneType_Radio       ItemPhoneType = "radio"
)

func PossibleValuesForItemPhoneType() []string {
	return []string{
		string(ItemPhoneType_Assistant),
		string(ItemPhoneType_Business),
		string(ItemPhoneType_BusinessFax),
		string(ItemPhoneType_Home),
		string(ItemPhoneType_HomeFax),
		string(ItemPhoneType_Mobile),
		string(ItemPhoneType_Other),
		string(ItemPhoneType_OtherFax),
		string(ItemPhoneType_Pager),
		string(ItemPhoneType_Radio),
	}
}

func (s *ItemPhoneType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemPhoneType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemPhoneType(input string) (*ItemPhoneType, error) {
	vals := map[string]ItemPhoneType{
		"assistant":   ItemPhoneType_Assistant,
		"business":    ItemPhoneType_Business,
		"businessfax": ItemPhoneType_BusinessFax,
		"home":        ItemPhoneType_Home,
		"homefax":     ItemPhoneType_HomeFax,
		"mobile":      ItemPhoneType_Mobile,
		"other":       ItemPhoneType_Other,
		"otherfax":    ItemPhoneType_OtherFax,
		"pager":       ItemPhoneType_Pager,
		"radio":       ItemPhoneType_Radio,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemPhoneType(input)
	return &out, nil
}
