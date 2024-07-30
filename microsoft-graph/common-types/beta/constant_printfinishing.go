package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintFinishing string

const (
	PrintFinishing_Bind              PrintFinishing = "bind"
	PrintFinishing_Cover             PrintFinishing = "cover"
	PrintFinishing_None              PrintFinishing = "none"
	PrintFinishing_Punch             PrintFinishing = "punch"
	PrintFinishing_SaddleStitch      PrintFinishing = "saddleStitch"
	PrintFinishing_Staple            PrintFinishing = "staple"
	PrintFinishing_StapleBottomLeft  PrintFinishing = "stapleBottomLeft"
	PrintFinishing_StapleBottomRight PrintFinishing = "stapleBottomRight"
	PrintFinishing_StapleDualBottom  PrintFinishing = "stapleDualBottom"
	PrintFinishing_StapleDualLeft    PrintFinishing = "stapleDualLeft"
	PrintFinishing_StapleDualRight   PrintFinishing = "stapleDualRight"
	PrintFinishing_StapleDualTop     PrintFinishing = "stapleDualTop"
	PrintFinishing_StapleTopLeft     PrintFinishing = "stapleTopLeft"
	PrintFinishing_StapleTopRight    PrintFinishing = "stapleTopRight"
	PrintFinishing_StitchBottomEdge  PrintFinishing = "stitchBottomEdge"
	PrintFinishing_StitchEdge        PrintFinishing = "stitchEdge"
	PrintFinishing_StitchLeftEdge    PrintFinishing = "stitchLeftEdge"
	PrintFinishing_StitchRightEdge   PrintFinishing = "stitchRightEdge"
	PrintFinishing_StitchTopEdge     PrintFinishing = "stitchTopEdge"
)

func PossibleValuesForPrintFinishing() []string {
	return []string{
		string(PrintFinishing_Bind),
		string(PrintFinishing_Cover),
		string(PrintFinishing_None),
		string(PrintFinishing_Punch),
		string(PrintFinishing_SaddleStitch),
		string(PrintFinishing_Staple),
		string(PrintFinishing_StapleBottomLeft),
		string(PrintFinishing_StapleBottomRight),
		string(PrintFinishing_StapleDualBottom),
		string(PrintFinishing_StapleDualLeft),
		string(PrintFinishing_StapleDualRight),
		string(PrintFinishing_StapleDualTop),
		string(PrintFinishing_StapleTopLeft),
		string(PrintFinishing_StapleTopRight),
		string(PrintFinishing_StitchBottomEdge),
		string(PrintFinishing_StitchEdge),
		string(PrintFinishing_StitchLeftEdge),
		string(PrintFinishing_StitchRightEdge),
		string(PrintFinishing_StitchTopEdge),
	}
}

func (s *PrintFinishing) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintFinishing(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintFinishing(input string) (*PrintFinishing, error) {
	vals := map[string]PrintFinishing{
		"bind":              PrintFinishing_Bind,
		"cover":             PrintFinishing_Cover,
		"none":              PrintFinishing_None,
		"punch":             PrintFinishing_Punch,
		"saddlestitch":      PrintFinishing_SaddleStitch,
		"staple":            PrintFinishing_Staple,
		"staplebottomleft":  PrintFinishing_StapleBottomLeft,
		"staplebottomright": PrintFinishing_StapleBottomRight,
		"stapledualbottom":  PrintFinishing_StapleDualBottom,
		"stapledualleft":    PrintFinishing_StapleDualLeft,
		"stapledualright":   PrintFinishing_StapleDualRight,
		"stapledualtop":     PrintFinishing_StapleDualTop,
		"stapletopleft":     PrintFinishing_StapleTopLeft,
		"stapletopright":    PrintFinishing_StapleTopRight,
		"stitchbottomedge":  PrintFinishing_StitchBottomEdge,
		"stitchedge":        PrintFinishing_StitchEdge,
		"stitchleftedge":    PrintFinishing_StitchLeftEdge,
		"stitchrightedge":   PrintFinishing_StitchRightEdge,
		"stitchtopedge":     PrintFinishing_StitchTopEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintFinishing(input)
	return &out, nil
}
