package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfigurationFinishings string

const (
	PrintJobConfigurationFinishings_Bind              PrintJobConfigurationFinishings = "bind"
	PrintJobConfigurationFinishings_Cover             PrintJobConfigurationFinishings = "cover"
	PrintJobConfigurationFinishings_None              PrintJobConfigurationFinishings = "none"
	PrintJobConfigurationFinishings_Punch             PrintJobConfigurationFinishings = "punch"
	PrintJobConfigurationFinishings_SaddleStitch      PrintJobConfigurationFinishings = "saddleStitch"
	PrintJobConfigurationFinishings_Staple            PrintJobConfigurationFinishings = "staple"
	PrintJobConfigurationFinishings_StapleBottomLeft  PrintJobConfigurationFinishings = "stapleBottomLeft"
	PrintJobConfigurationFinishings_StapleBottomRight PrintJobConfigurationFinishings = "stapleBottomRight"
	PrintJobConfigurationFinishings_StapleDualBottom  PrintJobConfigurationFinishings = "stapleDualBottom"
	PrintJobConfigurationFinishings_StapleDualLeft    PrintJobConfigurationFinishings = "stapleDualLeft"
	PrintJobConfigurationFinishings_StapleDualRight   PrintJobConfigurationFinishings = "stapleDualRight"
	PrintJobConfigurationFinishings_StapleDualTop     PrintJobConfigurationFinishings = "stapleDualTop"
	PrintJobConfigurationFinishings_StapleTopLeft     PrintJobConfigurationFinishings = "stapleTopLeft"
	PrintJobConfigurationFinishings_StapleTopRight    PrintJobConfigurationFinishings = "stapleTopRight"
	PrintJobConfigurationFinishings_StitchBottomEdge  PrintJobConfigurationFinishings = "stitchBottomEdge"
	PrintJobConfigurationFinishings_StitchEdge        PrintJobConfigurationFinishings = "stitchEdge"
	PrintJobConfigurationFinishings_StitchLeftEdge    PrintJobConfigurationFinishings = "stitchLeftEdge"
	PrintJobConfigurationFinishings_StitchRightEdge   PrintJobConfigurationFinishings = "stitchRightEdge"
	PrintJobConfigurationFinishings_StitchTopEdge     PrintJobConfigurationFinishings = "stitchTopEdge"
)

func PossibleValuesForPrintJobConfigurationFinishings() []string {
	return []string{
		string(PrintJobConfigurationFinishings_Bind),
		string(PrintJobConfigurationFinishings_Cover),
		string(PrintJobConfigurationFinishings_None),
		string(PrintJobConfigurationFinishings_Punch),
		string(PrintJobConfigurationFinishings_SaddleStitch),
		string(PrintJobConfigurationFinishings_Staple),
		string(PrintJobConfigurationFinishings_StapleBottomLeft),
		string(PrintJobConfigurationFinishings_StapleBottomRight),
		string(PrintJobConfigurationFinishings_StapleDualBottom),
		string(PrintJobConfigurationFinishings_StapleDualLeft),
		string(PrintJobConfigurationFinishings_StapleDualRight),
		string(PrintJobConfigurationFinishings_StapleDualTop),
		string(PrintJobConfigurationFinishings_StapleTopLeft),
		string(PrintJobConfigurationFinishings_StapleTopRight),
		string(PrintJobConfigurationFinishings_StitchBottomEdge),
		string(PrintJobConfigurationFinishings_StitchEdge),
		string(PrintJobConfigurationFinishings_StitchLeftEdge),
		string(PrintJobConfigurationFinishings_StitchRightEdge),
		string(PrintJobConfigurationFinishings_StitchTopEdge),
	}
}

func (s *PrintJobConfigurationFinishings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobConfigurationFinishings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobConfigurationFinishings(input string) (*PrintJobConfigurationFinishings, error) {
	vals := map[string]PrintJobConfigurationFinishings{
		"bind":              PrintJobConfigurationFinishings_Bind,
		"cover":             PrintJobConfigurationFinishings_Cover,
		"none":              PrintJobConfigurationFinishings_None,
		"punch":             PrintJobConfigurationFinishings_Punch,
		"saddlestitch":      PrintJobConfigurationFinishings_SaddleStitch,
		"staple":            PrintJobConfigurationFinishings_Staple,
		"staplebottomleft":  PrintJobConfigurationFinishings_StapleBottomLeft,
		"staplebottomright": PrintJobConfigurationFinishings_StapleBottomRight,
		"stapledualbottom":  PrintJobConfigurationFinishings_StapleDualBottom,
		"stapledualleft":    PrintJobConfigurationFinishings_StapleDualLeft,
		"stapledualright":   PrintJobConfigurationFinishings_StapleDualRight,
		"stapledualtop":     PrintJobConfigurationFinishings_StapleDualTop,
		"stapletopleft":     PrintJobConfigurationFinishings_StapleTopLeft,
		"stapletopright":    PrintJobConfigurationFinishings_StapleTopRight,
		"stitchbottomedge":  PrintJobConfigurationFinishings_StitchBottomEdge,
		"stitchedge":        PrintJobConfigurationFinishings_StitchEdge,
		"stitchleftedge":    PrintJobConfigurationFinishings_StitchLeftEdge,
		"stitchrightedge":   PrintJobConfigurationFinishings_StitchRightEdge,
		"stitchtopedge":     PrintJobConfigurationFinishings_StitchTopEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationFinishings(input)
	return &out, nil
}
