package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesFinishings string

const (
	PrinterCapabilitiesFinishingsbind              PrinterCapabilitiesFinishings = "Bind"
	PrinterCapabilitiesFinishingscover             PrinterCapabilitiesFinishings = "Cover"
	PrinterCapabilitiesFinishingsnone              PrinterCapabilitiesFinishings = "None"
	PrinterCapabilitiesFinishingspunch             PrinterCapabilitiesFinishings = "Punch"
	PrinterCapabilitiesFinishingssaddleStitch      PrinterCapabilitiesFinishings = "SaddleStitch"
	PrinterCapabilitiesFinishingsstaple            PrinterCapabilitiesFinishings = "Staple"
	PrinterCapabilitiesFinishingsstapleBottomLeft  PrinterCapabilitiesFinishings = "StapleBottomLeft"
	PrinterCapabilitiesFinishingsstapleBottomRight PrinterCapabilitiesFinishings = "StapleBottomRight"
	PrinterCapabilitiesFinishingsstapleDualBottom  PrinterCapabilitiesFinishings = "StapleDualBottom"
	PrinterCapabilitiesFinishingsstapleDualLeft    PrinterCapabilitiesFinishings = "StapleDualLeft"
	PrinterCapabilitiesFinishingsstapleDualRight   PrinterCapabilitiesFinishings = "StapleDualRight"
	PrinterCapabilitiesFinishingsstapleDualTop     PrinterCapabilitiesFinishings = "StapleDualTop"
	PrinterCapabilitiesFinishingsstapleTopLeft     PrinterCapabilitiesFinishings = "StapleTopLeft"
	PrinterCapabilitiesFinishingsstapleTopRight    PrinterCapabilitiesFinishings = "StapleTopRight"
	PrinterCapabilitiesFinishingsstitchBottomEdge  PrinterCapabilitiesFinishings = "StitchBottomEdge"
	PrinterCapabilitiesFinishingsstitchEdge        PrinterCapabilitiesFinishings = "StitchEdge"
	PrinterCapabilitiesFinishingsstitchLeftEdge    PrinterCapabilitiesFinishings = "StitchLeftEdge"
	PrinterCapabilitiesFinishingsstitchRightEdge   PrinterCapabilitiesFinishings = "StitchRightEdge"
	PrinterCapabilitiesFinishingsstitchTopEdge     PrinterCapabilitiesFinishings = "StitchTopEdge"
)

func PossibleValuesForPrinterCapabilitiesFinishings() []string {
	return []string{
		string(PrinterCapabilitiesFinishingsbind),
		string(PrinterCapabilitiesFinishingscover),
		string(PrinterCapabilitiesFinishingsnone),
		string(PrinterCapabilitiesFinishingspunch),
		string(PrinterCapabilitiesFinishingssaddleStitch),
		string(PrinterCapabilitiesFinishingsstaple),
		string(PrinterCapabilitiesFinishingsstapleBottomLeft),
		string(PrinterCapabilitiesFinishingsstapleBottomRight),
		string(PrinterCapabilitiesFinishingsstapleDualBottom),
		string(PrinterCapabilitiesFinishingsstapleDualLeft),
		string(PrinterCapabilitiesFinishingsstapleDualRight),
		string(PrinterCapabilitiesFinishingsstapleDualTop),
		string(PrinterCapabilitiesFinishingsstapleTopLeft),
		string(PrinterCapabilitiesFinishingsstapleTopRight),
		string(PrinterCapabilitiesFinishingsstitchBottomEdge),
		string(PrinterCapabilitiesFinishingsstitchEdge),
		string(PrinterCapabilitiesFinishingsstitchLeftEdge),
		string(PrinterCapabilitiesFinishingsstitchRightEdge),
		string(PrinterCapabilitiesFinishingsstitchTopEdge),
	}
}

func parsePrinterCapabilitiesFinishings(input string) (*PrinterCapabilitiesFinishings, error) {
	vals := map[string]PrinterCapabilitiesFinishings{
		"bind":              PrinterCapabilitiesFinishingsbind,
		"cover":             PrinterCapabilitiesFinishingscover,
		"none":              PrinterCapabilitiesFinishingsnone,
		"punch":             PrinterCapabilitiesFinishingspunch,
		"saddlestitch":      PrinterCapabilitiesFinishingssaddleStitch,
		"staple":            PrinterCapabilitiesFinishingsstaple,
		"staplebottomleft":  PrinterCapabilitiesFinishingsstapleBottomLeft,
		"staplebottomright": PrinterCapabilitiesFinishingsstapleBottomRight,
		"stapledualbottom":  PrinterCapabilitiesFinishingsstapleDualBottom,
		"stapledualleft":    PrinterCapabilitiesFinishingsstapleDualLeft,
		"stapledualright":   PrinterCapabilitiesFinishingsstapleDualRight,
		"stapledualtop":     PrinterCapabilitiesFinishingsstapleDualTop,
		"stapletopleft":     PrinterCapabilitiesFinishingsstapleTopLeft,
		"stapletopright":    PrinterCapabilitiesFinishingsstapleTopRight,
		"stitchbottomedge":  PrinterCapabilitiesFinishingsstitchBottomEdge,
		"stitchedge":        PrinterCapabilitiesFinishingsstitchEdge,
		"stitchleftedge":    PrinterCapabilitiesFinishingsstitchLeftEdge,
		"stitchrightedge":   PrinterCapabilitiesFinishingsstitchRightEdge,
		"stitchtopedge":     PrinterCapabilitiesFinishingsstitchTopEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesFinishings(input)
	return &out, nil
}
