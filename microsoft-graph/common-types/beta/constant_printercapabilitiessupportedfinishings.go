package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesSupportedFinishings string

const (
	PrinterCapabilitiesSupportedFinishings_Bind              PrinterCapabilitiesSupportedFinishings = "bind"
	PrinterCapabilitiesSupportedFinishings_Cover             PrinterCapabilitiesSupportedFinishings = "cover"
	PrinterCapabilitiesSupportedFinishings_None              PrinterCapabilitiesSupportedFinishings = "none"
	PrinterCapabilitiesSupportedFinishings_Punch             PrinterCapabilitiesSupportedFinishings = "punch"
	PrinterCapabilitiesSupportedFinishings_SaddleStitch      PrinterCapabilitiesSupportedFinishings = "saddleStitch"
	PrinterCapabilitiesSupportedFinishings_Staple            PrinterCapabilitiesSupportedFinishings = "staple"
	PrinterCapabilitiesSupportedFinishings_StapleBottomLeft  PrinterCapabilitiesSupportedFinishings = "stapleBottomLeft"
	PrinterCapabilitiesSupportedFinishings_StapleBottomRight PrinterCapabilitiesSupportedFinishings = "stapleBottomRight"
	PrinterCapabilitiesSupportedFinishings_StapleDualBottom  PrinterCapabilitiesSupportedFinishings = "stapleDualBottom"
	PrinterCapabilitiesSupportedFinishings_StapleDualLeft    PrinterCapabilitiesSupportedFinishings = "stapleDualLeft"
	PrinterCapabilitiesSupportedFinishings_StapleDualRight   PrinterCapabilitiesSupportedFinishings = "stapleDualRight"
	PrinterCapabilitiesSupportedFinishings_StapleDualTop     PrinterCapabilitiesSupportedFinishings = "stapleDualTop"
	PrinterCapabilitiesSupportedFinishings_StapleTopLeft     PrinterCapabilitiesSupportedFinishings = "stapleTopLeft"
	PrinterCapabilitiesSupportedFinishings_StapleTopRight    PrinterCapabilitiesSupportedFinishings = "stapleTopRight"
	PrinterCapabilitiesSupportedFinishings_StitchBottomEdge  PrinterCapabilitiesSupportedFinishings = "stitchBottomEdge"
	PrinterCapabilitiesSupportedFinishings_StitchEdge        PrinterCapabilitiesSupportedFinishings = "stitchEdge"
	PrinterCapabilitiesSupportedFinishings_StitchLeftEdge    PrinterCapabilitiesSupportedFinishings = "stitchLeftEdge"
	PrinterCapabilitiesSupportedFinishings_StitchRightEdge   PrinterCapabilitiesSupportedFinishings = "stitchRightEdge"
	PrinterCapabilitiesSupportedFinishings_StitchTopEdge     PrinterCapabilitiesSupportedFinishings = "stitchTopEdge"
)

func PossibleValuesForPrinterCapabilitiesSupportedFinishings() []string {
	return []string{
		string(PrinterCapabilitiesSupportedFinishings_Bind),
		string(PrinterCapabilitiesSupportedFinishings_Cover),
		string(PrinterCapabilitiesSupportedFinishings_None),
		string(PrinterCapabilitiesSupportedFinishings_Punch),
		string(PrinterCapabilitiesSupportedFinishings_SaddleStitch),
		string(PrinterCapabilitiesSupportedFinishings_Staple),
		string(PrinterCapabilitiesSupportedFinishings_StapleBottomLeft),
		string(PrinterCapabilitiesSupportedFinishings_StapleBottomRight),
		string(PrinterCapabilitiesSupportedFinishings_StapleDualBottom),
		string(PrinterCapabilitiesSupportedFinishings_StapleDualLeft),
		string(PrinterCapabilitiesSupportedFinishings_StapleDualRight),
		string(PrinterCapabilitiesSupportedFinishings_StapleDualTop),
		string(PrinterCapabilitiesSupportedFinishings_StapleTopLeft),
		string(PrinterCapabilitiesSupportedFinishings_StapleTopRight),
		string(PrinterCapabilitiesSupportedFinishings_StitchBottomEdge),
		string(PrinterCapabilitiesSupportedFinishings_StitchEdge),
		string(PrinterCapabilitiesSupportedFinishings_StitchLeftEdge),
		string(PrinterCapabilitiesSupportedFinishings_StitchRightEdge),
		string(PrinterCapabilitiesSupportedFinishings_StitchTopEdge),
	}
}

func (s *PrinterCapabilitiesSupportedFinishings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesSupportedFinishings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesSupportedFinishings(input string) (*PrinterCapabilitiesSupportedFinishings, error) {
	vals := map[string]PrinterCapabilitiesSupportedFinishings{
		"bind":              PrinterCapabilitiesSupportedFinishings_Bind,
		"cover":             PrinterCapabilitiesSupportedFinishings_Cover,
		"none":              PrinterCapabilitiesSupportedFinishings_None,
		"punch":             PrinterCapabilitiesSupportedFinishings_Punch,
		"saddlestitch":      PrinterCapabilitiesSupportedFinishings_SaddleStitch,
		"staple":            PrinterCapabilitiesSupportedFinishings_Staple,
		"staplebottomleft":  PrinterCapabilitiesSupportedFinishings_StapleBottomLeft,
		"staplebottomright": PrinterCapabilitiesSupportedFinishings_StapleBottomRight,
		"stapledualbottom":  PrinterCapabilitiesSupportedFinishings_StapleDualBottom,
		"stapledualleft":    PrinterCapabilitiesSupportedFinishings_StapleDualLeft,
		"stapledualright":   PrinterCapabilitiesSupportedFinishings_StapleDualRight,
		"stapledualtop":     PrinterCapabilitiesSupportedFinishings_StapleDualTop,
		"stapletopleft":     PrinterCapabilitiesSupportedFinishings_StapleTopLeft,
		"stapletopright":    PrinterCapabilitiesSupportedFinishings_StapleTopRight,
		"stitchbottomedge":  PrinterCapabilitiesSupportedFinishings_StitchBottomEdge,
		"stitchedge":        PrinterCapabilitiesSupportedFinishings_StitchEdge,
		"stitchleftedge":    PrinterCapabilitiesSupportedFinishings_StitchLeftEdge,
		"stitchrightedge":   PrinterCapabilitiesSupportedFinishings_StitchRightEdge,
		"stitchtopedge":     PrinterCapabilitiesSupportedFinishings_StitchTopEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedFinishings(input)
	return &out, nil
}
