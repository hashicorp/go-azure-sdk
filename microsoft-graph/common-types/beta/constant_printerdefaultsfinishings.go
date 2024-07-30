package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsFinishings string

const (
	PrinterDefaultsFinishings_Bind              PrinterDefaultsFinishings = "bind"
	PrinterDefaultsFinishings_Cover             PrinterDefaultsFinishings = "cover"
	PrinterDefaultsFinishings_None              PrinterDefaultsFinishings = "none"
	PrinterDefaultsFinishings_Punch             PrinterDefaultsFinishings = "punch"
	PrinterDefaultsFinishings_SaddleStitch      PrinterDefaultsFinishings = "saddleStitch"
	PrinterDefaultsFinishings_Staple            PrinterDefaultsFinishings = "staple"
	PrinterDefaultsFinishings_StapleBottomLeft  PrinterDefaultsFinishings = "stapleBottomLeft"
	PrinterDefaultsFinishings_StapleBottomRight PrinterDefaultsFinishings = "stapleBottomRight"
	PrinterDefaultsFinishings_StapleDualBottom  PrinterDefaultsFinishings = "stapleDualBottom"
	PrinterDefaultsFinishings_StapleDualLeft    PrinterDefaultsFinishings = "stapleDualLeft"
	PrinterDefaultsFinishings_StapleDualRight   PrinterDefaultsFinishings = "stapleDualRight"
	PrinterDefaultsFinishings_StapleDualTop     PrinterDefaultsFinishings = "stapleDualTop"
	PrinterDefaultsFinishings_StapleTopLeft     PrinterDefaultsFinishings = "stapleTopLeft"
	PrinterDefaultsFinishings_StapleTopRight    PrinterDefaultsFinishings = "stapleTopRight"
	PrinterDefaultsFinishings_StitchBottomEdge  PrinterDefaultsFinishings = "stitchBottomEdge"
	PrinterDefaultsFinishings_StitchEdge        PrinterDefaultsFinishings = "stitchEdge"
	PrinterDefaultsFinishings_StitchLeftEdge    PrinterDefaultsFinishings = "stitchLeftEdge"
	PrinterDefaultsFinishings_StitchRightEdge   PrinterDefaultsFinishings = "stitchRightEdge"
	PrinterDefaultsFinishings_StitchTopEdge     PrinterDefaultsFinishings = "stitchTopEdge"
)

func PossibleValuesForPrinterDefaultsFinishings() []string {
	return []string{
		string(PrinterDefaultsFinishings_Bind),
		string(PrinterDefaultsFinishings_Cover),
		string(PrinterDefaultsFinishings_None),
		string(PrinterDefaultsFinishings_Punch),
		string(PrinterDefaultsFinishings_SaddleStitch),
		string(PrinterDefaultsFinishings_Staple),
		string(PrinterDefaultsFinishings_StapleBottomLeft),
		string(PrinterDefaultsFinishings_StapleBottomRight),
		string(PrinterDefaultsFinishings_StapleDualBottom),
		string(PrinterDefaultsFinishings_StapleDualLeft),
		string(PrinterDefaultsFinishings_StapleDualRight),
		string(PrinterDefaultsFinishings_StapleDualTop),
		string(PrinterDefaultsFinishings_StapleTopLeft),
		string(PrinterDefaultsFinishings_StapleTopRight),
		string(PrinterDefaultsFinishings_StitchBottomEdge),
		string(PrinterDefaultsFinishings_StitchEdge),
		string(PrinterDefaultsFinishings_StitchLeftEdge),
		string(PrinterDefaultsFinishings_StitchRightEdge),
		string(PrinterDefaultsFinishings_StitchTopEdge),
	}
}

func (s *PrinterDefaultsFinishings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDefaultsFinishings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDefaultsFinishings(input string) (*PrinterDefaultsFinishings, error) {
	vals := map[string]PrinterDefaultsFinishings{
		"bind":              PrinterDefaultsFinishings_Bind,
		"cover":             PrinterDefaultsFinishings_Cover,
		"none":              PrinterDefaultsFinishings_None,
		"punch":             PrinterDefaultsFinishings_Punch,
		"saddlestitch":      PrinterDefaultsFinishings_SaddleStitch,
		"staple":            PrinterDefaultsFinishings_Staple,
		"staplebottomleft":  PrinterDefaultsFinishings_StapleBottomLeft,
		"staplebottomright": PrinterDefaultsFinishings_StapleBottomRight,
		"stapledualbottom":  PrinterDefaultsFinishings_StapleDualBottom,
		"stapledualleft":    PrinterDefaultsFinishings_StapleDualLeft,
		"stapledualright":   PrinterDefaultsFinishings_StapleDualRight,
		"stapledualtop":     PrinterDefaultsFinishings_StapleDualTop,
		"stapletopleft":     PrinterDefaultsFinishings_StapleTopLeft,
		"stapletopright":    PrinterDefaultsFinishings_StapleTopRight,
		"stitchbottomedge":  PrinterDefaultsFinishings_StitchBottomEdge,
		"stitchedge":        PrinterDefaultsFinishings_StitchEdge,
		"stitchleftedge":    PrinterDefaultsFinishings_StitchLeftEdge,
		"stitchrightedge":   PrinterDefaultsFinishings_StitchRightEdge,
		"stitchtopedge":     PrinterDefaultsFinishings_StitchTopEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsFinishings(input)
	return &out, nil
}
