package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TitleAreaLayout string

const (
	TitleAreaLayoutcolorBlock    TitleAreaLayout = "ColorBlock"
	TitleAreaLayoutimageAndTitle TitleAreaLayout = "ImageAndTitle"
	TitleAreaLayoutoverlap       TitleAreaLayout = "Overlap"
	TitleAreaLayoutplain         TitleAreaLayout = "Plain"
)

func PossibleValuesForTitleAreaLayout() []string {
	return []string{
		string(TitleAreaLayoutcolorBlock),
		string(TitleAreaLayoutimageAndTitle),
		string(TitleAreaLayoutoverlap),
		string(TitleAreaLayoutplain),
	}
}

func parseTitleAreaLayout(input string) (*TitleAreaLayout, error) {
	vals := map[string]TitleAreaLayout{
		"colorblock":    TitleAreaLayoutcolorBlock,
		"imageandtitle": TitleAreaLayoutimageAndTitle,
		"overlap":       TitleAreaLayoutoverlap,
		"plain":         TitleAreaLayoutplain,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TitleAreaLayout(input)
	return &out, nil
}
