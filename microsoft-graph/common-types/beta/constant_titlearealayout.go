package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TitleAreaLayout string

const (
	TitleAreaLayout_ColorBlock    TitleAreaLayout = "colorBlock"
	TitleAreaLayout_ImageAndTitle TitleAreaLayout = "imageAndTitle"
	TitleAreaLayout_Overlap       TitleAreaLayout = "overlap"
	TitleAreaLayout_Plain         TitleAreaLayout = "plain"
)

func PossibleValuesForTitleAreaLayout() []string {
	return []string{
		string(TitleAreaLayout_ColorBlock),
		string(TitleAreaLayout_ImageAndTitle),
		string(TitleAreaLayout_Overlap),
		string(TitleAreaLayout_Plain),
	}
}

func (s *TitleAreaLayout) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTitleAreaLayout(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTitleAreaLayout(input string) (*TitleAreaLayout, error) {
	vals := map[string]TitleAreaLayout{
		"colorblock":    TitleAreaLayout_ColorBlock,
		"imageandtitle": TitleAreaLayout_ImageAndTitle,
		"overlap":       TitleAreaLayout_Overlap,
		"plain":         TitleAreaLayout_Plain,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TitleAreaLayout(input)
	return &out, nil
}
