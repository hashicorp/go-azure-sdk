package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesFinishings string

const (
	PrinterCapabilitiesFinishings_Bind              PrinterCapabilitiesFinishings = "bind"
	PrinterCapabilitiesFinishings_Cover             PrinterCapabilitiesFinishings = "cover"
	PrinterCapabilitiesFinishings_None              PrinterCapabilitiesFinishings = "none"
	PrinterCapabilitiesFinishings_Punch             PrinterCapabilitiesFinishings = "punch"
	PrinterCapabilitiesFinishings_SaddleStitch      PrinterCapabilitiesFinishings = "saddleStitch"
	PrinterCapabilitiesFinishings_Staple            PrinterCapabilitiesFinishings = "staple"
	PrinterCapabilitiesFinishings_StapleBottomLeft  PrinterCapabilitiesFinishings = "stapleBottomLeft"
	PrinterCapabilitiesFinishings_StapleBottomRight PrinterCapabilitiesFinishings = "stapleBottomRight"
	PrinterCapabilitiesFinishings_StapleDualBottom  PrinterCapabilitiesFinishings = "stapleDualBottom"
	PrinterCapabilitiesFinishings_StapleDualLeft    PrinterCapabilitiesFinishings = "stapleDualLeft"
	PrinterCapabilitiesFinishings_StapleDualRight   PrinterCapabilitiesFinishings = "stapleDualRight"
	PrinterCapabilitiesFinishings_StapleDualTop     PrinterCapabilitiesFinishings = "stapleDualTop"
	PrinterCapabilitiesFinishings_StapleTopLeft     PrinterCapabilitiesFinishings = "stapleTopLeft"
	PrinterCapabilitiesFinishings_StapleTopRight    PrinterCapabilitiesFinishings = "stapleTopRight"
	PrinterCapabilitiesFinishings_StitchBottomEdge  PrinterCapabilitiesFinishings = "stitchBottomEdge"
	PrinterCapabilitiesFinishings_StitchEdge        PrinterCapabilitiesFinishings = "stitchEdge"
	PrinterCapabilitiesFinishings_StitchLeftEdge    PrinterCapabilitiesFinishings = "stitchLeftEdge"
	PrinterCapabilitiesFinishings_StitchRightEdge   PrinterCapabilitiesFinishings = "stitchRightEdge"
	PrinterCapabilitiesFinishings_StitchTopEdge     PrinterCapabilitiesFinishings = "stitchTopEdge"
)

func PossibleValuesForPrinterCapabilitiesFinishings() []string {
	return []string{
		string(PrinterCapabilitiesFinishings_Bind),
		string(PrinterCapabilitiesFinishings_Cover),
		string(PrinterCapabilitiesFinishings_None),
		string(PrinterCapabilitiesFinishings_Punch),
		string(PrinterCapabilitiesFinishings_SaddleStitch),
		string(PrinterCapabilitiesFinishings_Staple),
		string(PrinterCapabilitiesFinishings_StapleBottomLeft),
		string(PrinterCapabilitiesFinishings_StapleBottomRight),
		string(PrinterCapabilitiesFinishings_StapleDualBottom),
		string(PrinterCapabilitiesFinishings_StapleDualLeft),
		string(PrinterCapabilitiesFinishings_StapleDualRight),
		string(PrinterCapabilitiesFinishings_StapleDualTop),
		string(PrinterCapabilitiesFinishings_StapleTopLeft),
		string(PrinterCapabilitiesFinishings_StapleTopRight),
		string(PrinterCapabilitiesFinishings_StitchBottomEdge),
		string(PrinterCapabilitiesFinishings_StitchEdge),
		string(PrinterCapabilitiesFinishings_StitchLeftEdge),
		string(PrinterCapabilitiesFinishings_StitchRightEdge),
		string(PrinterCapabilitiesFinishings_StitchTopEdge),
	}
}

func (s *PrinterCapabilitiesFinishings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesFinishings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesFinishings(input string) (*PrinterCapabilitiesFinishings, error) {
	vals := map[string]PrinterCapabilitiesFinishings{
		"bind":              PrinterCapabilitiesFinishings_Bind,
		"cover":             PrinterCapabilitiesFinishings_Cover,
		"none":              PrinterCapabilitiesFinishings_None,
		"punch":             PrinterCapabilitiesFinishings_Punch,
		"saddlestitch":      PrinterCapabilitiesFinishings_SaddleStitch,
		"staple":            PrinterCapabilitiesFinishings_Staple,
		"staplebottomleft":  PrinterCapabilitiesFinishings_StapleBottomLeft,
		"staplebottomright": PrinterCapabilitiesFinishings_StapleBottomRight,
		"stapledualbottom":  PrinterCapabilitiesFinishings_StapleDualBottom,
		"stapledualleft":    PrinterCapabilitiesFinishings_StapleDualLeft,
		"stapledualright":   PrinterCapabilitiesFinishings_StapleDualRight,
		"stapledualtop":     PrinterCapabilitiesFinishings_StapleDualTop,
		"stapletopleft":     PrinterCapabilitiesFinishings_StapleTopLeft,
		"stapletopright":    PrinterCapabilitiesFinishings_StapleTopRight,
		"stitchbottomedge":  PrinterCapabilitiesFinishings_StitchBottomEdge,
		"stitchedge":        PrinterCapabilitiesFinishings_StitchEdge,
		"stitchleftedge":    PrinterCapabilitiesFinishings_StitchLeftEdge,
		"stitchrightedge":   PrinterCapabilitiesFinishings_StitchRightEdge,
		"stitchtopedge":     PrinterCapabilitiesFinishings_StitchTopEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesFinishings(input)
	return &out, nil
}
