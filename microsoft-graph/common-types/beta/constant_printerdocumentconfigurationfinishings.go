package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfigurationFinishings string

const (
	PrinterDocumentConfigurationFinishings_Bind              PrinterDocumentConfigurationFinishings = "bind"
	PrinterDocumentConfigurationFinishings_Cover             PrinterDocumentConfigurationFinishings = "cover"
	PrinterDocumentConfigurationFinishings_None              PrinterDocumentConfigurationFinishings = "none"
	PrinterDocumentConfigurationFinishings_Punch             PrinterDocumentConfigurationFinishings = "punch"
	PrinterDocumentConfigurationFinishings_SaddleStitch      PrinterDocumentConfigurationFinishings = "saddleStitch"
	PrinterDocumentConfigurationFinishings_Staple            PrinterDocumentConfigurationFinishings = "staple"
	PrinterDocumentConfigurationFinishings_StapleBottomLeft  PrinterDocumentConfigurationFinishings = "stapleBottomLeft"
	PrinterDocumentConfigurationFinishings_StapleBottomRight PrinterDocumentConfigurationFinishings = "stapleBottomRight"
	PrinterDocumentConfigurationFinishings_StapleDualBottom  PrinterDocumentConfigurationFinishings = "stapleDualBottom"
	PrinterDocumentConfigurationFinishings_StapleDualLeft    PrinterDocumentConfigurationFinishings = "stapleDualLeft"
	PrinterDocumentConfigurationFinishings_StapleDualRight   PrinterDocumentConfigurationFinishings = "stapleDualRight"
	PrinterDocumentConfigurationFinishings_StapleDualTop     PrinterDocumentConfigurationFinishings = "stapleDualTop"
	PrinterDocumentConfigurationFinishings_StapleTopLeft     PrinterDocumentConfigurationFinishings = "stapleTopLeft"
	PrinterDocumentConfigurationFinishings_StapleTopRight    PrinterDocumentConfigurationFinishings = "stapleTopRight"
	PrinterDocumentConfigurationFinishings_StitchBottomEdge  PrinterDocumentConfigurationFinishings = "stitchBottomEdge"
	PrinterDocumentConfigurationFinishings_StitchEdge        PrinterDocumentConfigurationFinishings = "stitchEdge"
	PrinterDocumentConfigurationFinishings_StitchLeftEdge    PrinterDocumentConfigurationFinishings = "stitchLeftEdge"
	PrinterDocumentConfigurationFinishings_StitchRightEdge   PrinterDocumentConfigurationFinishings = "stitchRightEdge"
	PrinterDocumentConfigurationFinishings_StitchTopEdge     PrinterDocumentConfigurationFinishings = "stitchTopEdge"
)

func PossibleValuesForPrinterDocumentConfigurationFinishings() []string {
	return []string{
		string(PrinterDocumentConfigurationFinishings_Bind),
		string(PrinterDocumentConfigurationFinishings_Cover),
		string(PrinterDocumentConfigurationFinishings_None),
		string(PrinterDocumentConfigurationFinishings_Punch),
		string(PrinterDocumentConfigurationFinishings_SaddleStitch),
		string(PrinterDocumentConfigurationFinishings_Staple),
		string(PrinterDocumentConfigurationFinishings_StapleBottomLeft),
		string(PrinterDocumentConfigurationFinishings_StapleBottomRight),
		string(PrinterDocumentConfigurationFinishings_StapleDualBottom),
		string(PrinterDocumentConfigurationFinishings_StapleDualLeft),
		string(PrinterDocumentConfigurationFinishings_StapleDualRight),
		string(PrinterDocumentConfigurationFinishings_StapleDualTop),
		string(PrinterDocumentConfigurationFinishings_StapleTopLeft),
		string(PrinterDocumentConfigurationFinishings_StapleTopRight),
		string(PrinterDocumentConfigurationFinishings_StitchBottomEdge),
		string(PrinterDocumentConfigurationFinishings_StitchEdge),
		string(PrinterDocumentConfigurationFinishings_StitchLeftEdge),
		string(PrinterDocumentConfigurationFinishings_StitchRightEdge),
		string(PrinterDocumentConfigurationFinishings_StitchTopEdge),
	}
}

func (s *PrinterDocumentConfigurationFinishings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterDocumentConfigurationFinishings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterDocumentConfigurationFinishings(input string) (*PrinterDocumentConfigurationFinishings, error) {
	vals := map[string]PrinterDocumentConfigurationFinishings{
		"bind":              PrinterDocumentConfigurationFinishings_Bind,
		"cover":             PrinterDocumentConfigurationFinishings_Cover,
		"none":              PrinterDocumentConfigurationFinishings_None,
		"punch":             PrinterDocumentConfigurationFinishings_Punch,
		"saddlestitch":      PrinterDocumentConfigurationFinishings_SaddleStitch,
		"staple":            PrinterDocumentConfigurationFinishings_Staple,
		"staplebottomleft":  PrinterDocumentConfigurationFinishings_StapleBottomLeft,
		"staplebottomright": PrinterDocumentConfigurationFinishings_StapleBottomRight,
		"stapledualbottom":  PrinterDocumentConfigurationFinishings_StapleDualBottom,
		"stapledualleft":    PrinterDocumentConfigurationFinishings_StapleDualLeft,
		"stapledualright":   PrinterDocumentConfigurationFinishings_StapleDualRight,
		"stapledualtop":     PrinterDocumentConfigurationFinishings_StapleDualTop,
		"stapletopleft":     PrinterDocumentConfigurationFinishings_StapleTopLeft,
		"stapletopright":    PrinterDocumentConfigurationFinishings_StapleTopRight,
		"stitchbottomedge":  PrinterDocumentConfigurationFinishings_StitchBottomEdge,
		"stitchedge":        PrinterDocumentConfigurationFinishings_StitchEdge,
		"stitchleftedge":    PrinterDocumentConfigurationFinishings_StitchLeftEdge,
		"stitchrightedge":   PrinterDocumentConfigurationFinishings_StitchRightEdge,
		"stitchtopedge":     PrinterDocumentConfigurationFinishings_StitchTopEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationFinishings(input)
	return &out, nil
}
