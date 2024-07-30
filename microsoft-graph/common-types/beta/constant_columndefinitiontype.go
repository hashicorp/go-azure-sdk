package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ColumnDefinitionType string

const (
	ColumnDefinitionType_ApprovalStatus ColumnDefinitionType = "approvalStatus"
	ColumnDefinitionType_Boolean        ColumnDefinitionType = "boolean"
	ColumnDefinitionType_Calculated     ColumnDefinitionType = "calculated"
	ColumnDefinitionType_Choice         ColumnDefinitionType = "choice"
	ColumnDefinitionType_Currency       ColumnDefinitionType = "currency"
	ColumnDefinitionType_DateTime       ColumnDefinitionType = "dateTime"
	ColumnDefinitionType_Geolocation    ColumnDefinitionType = "geolocation"
	ColumnDefinitionType_Location       ColumnDefinitionType = "location"
	ColumnDefinitionType_Lookup         ColumnDefinitionType = "lookup"
	ColumnDefinitionType_Multichoice    ColumnDefinitionType = "multichoice"
	ColumnDefinitionType_Multiterm      ColumnDefinitionType = "multiterm"
	ColumnDefinitionType_Note           ColumnDefinitionType = "note"
	ColumnDefinitionType_Number         ColumnDefinitionType = "number"
	ColumnDefinitionType_Term           ColumnDefinitionType = "term"
	ColumnDefinitionType_Text           ColumnDefinitionType = "text"
	ColumnDefinitionType_Thumbnail      ColumnDefinitionType = "thumbnail"
	ColumnDefinitionType_Url            ColumnDefinitionType = "url"
	ColumnDefinitionType_User           ColumnDefinitionType = "user"
)

func PossibleValuesForColumnDefinitionType() []string {
	return []string{
		string(ColumnDefinitionType_ApprovalStatus),
		string(ColumnDefinitionType_Boolean),
		string(ColumnDefinitionType_Calculated),
		string(ColumnDefinitionType_Choice),
		string(ColumnDefinitionType_Currency),
		string(ColumnDefinitionType_DateTime),
		string(ColumnDefinitionType_Geolocation),
		string(ColumnDefinitionType_Location),
		string(ColumnDefinitionType_Lookup),
		string(ColumnDefinitionType_Multichoice),
		string(ColumnDefinitionType_Multiterm),
		string(ColumnDefinitionType_Note),
		string(ColumnDefinitionType_Number),
		string(ColumnDefinitionType_Term),
		string(ColumnDefinitionType_Text),
		string(ColumnDefinitionType_Thumbnail),
		string(ColumnDefinitionType_Url),
		string(ColumnDefinitionType_User),
	}
}

func (s *ColumnDefinitionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseColumnDefinitionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseColumnDefinitionType(input string) (*ColumnDefinitionType, error) {
	vals := map[string]ColumnDefinitionType{
		"approvalstatus": ColumnDefinitionType_ApprovalStatus,
		"boolean":        ColumnDefinitionType_Boolean,
		"calculated":     ColumnDefinitionType_Calculated,
		"choice":         ColumnDefinitionType_Choice,
		"currency":       ColumnDefinitionType_Currency,
		"datetime":       ColumnDefinitionType_DateTime,
		"geolocation":    ColumnDefinitionType_Geolocation,
		"location":       ColumnDefinitionType_Location,
		"lookup":         ColumnDefinitionType_Lookup,
		"multichoice":    ColumnDefinitionType_Multichoice,
		"multiterm":      ColumnDefinitionType_Multiterm,
		"note":           ColumnDefinitionType_Note,
		"number":         ColumnDefinitionType_Number,
		"term":           ColumnDefinitionType_Term,
		"text":           ColumnDefinitionType_Text,
		"thumbnail":      ColumnDefinitionType_Thumbnail,
		"url":            ColumnDefinitionType_Url,
		"user":           ColumnDefinitionType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ColumnDefinitionType(input)
	return &out, nil
}
