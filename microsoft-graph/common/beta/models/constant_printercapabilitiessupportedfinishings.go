package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesSupportedFinishings string

const (
	PrinterCapabilitiesSupportedFinishingsbind              PrinterCapabilitiesSupportedFinishings = "Bind"
	PrinterCapabilitiesSupportedFinishingscover             PrinterCapabilitiesSupportedFinishings = "Cover"
	PrinterCapabilitiesSupportedFinishingsnone              PrinterCapabilitiesSupportedFinishings = "None"
	PrinterCapabilitiesSupportedFinishingspunch             PrinterCapabilitiesSupportedFinishings = "Punch"
	PrinterCapabilitiesSupportedFinishingssaddleStitch      PrinterCapabilitiesSupportedFinishings = "SaddleStitch"
	PrinterCapabilitiesSupportedFinishingsstaple            PrinterCapabilitiesSupportedFinishings = "Staple"
	PrinterCapabilitiesSupportedFinishingsstapleBottomLeft  PrinterCapabilitiesSupportedFinishings = "StapleBottomLeft"
	PrinterCapabilitiesSupportedFinishingsstapleBottomRight PrinterCapabilitiesSupportedFinishings = "StapleBottomRight"
	PrinterCapabilitiesSupportedFinishingsstapleDualBottom  PrinterCapabilitiesSupportedFinishings = "StapleDualBottom"
	PrinterCapabilitiesSupportedFinishingsstapleDualLeft    PrinterCapabilitiesSupportedFinishings = "StapleDualLeft"
	PrinterCapabilitiesSupportedFinishingsstapleDualRight   PrinterCapabilitiesSupportedFinishings = "StapleDualRight"
	PrinterCapabilitiesSupportedFinishingsstapleDualTop     PrinterCapabilitiesSupportedFinishings = "StapleDualTop"
	PrinterCapabilitiesSupportedFinishingsstapleTopLeft     PrinterCapabilitiesSupportedFinishings = "StapleTopLeft"
	PrinterCapabilitiesSupportedFinishingsstapleTopRight    PrinterCapabilitiesSupportedFinishings = "StapleTopRight"
	PrinterCapabilitiesSupportedFinishingsstitchBottomEdge  PrinterCapabilitiesSupportedFinishings = "StitchBottomEdge"
	PrinterCapabilitiesSupportedFinishingsstitchEdge        PrinterCapabilitiesSupportedFinishings = "StitchEdge"
	PrinterCapabilitiesSupportedFinishingsstitchLeftEdge    PrinterCapabilitiesSupportedFinishings = "StitchLeftEdge"
	PrinterCapabilitiesSupportedFinishingsstitchRightEdge   PrinterCapabilitiesSupportedFinishings = "StitchRightEdge"
	PrinterCapabilitiesSupportedFinishingsstitchTopEdge     PrinterCapabilitiesSupportedFinishings = "StitchTopEdge"
)

func PossibleValuesForPrinterCapabilitiesSupportedFinishings() []string {
	return []string{
		string(PrinterCapabilitiesSupportedFinishingsbind),
		string(PrinterCapabilitiesSupportedFinishingscover),
		string(PrinterCapabilitiesSupportedFinishingsnone),
		string(PrinterCapabilitiesSupportedFinishingspunch),
		string(PrinterCapabilitiesSupportedFinishingssaddleStitch),
		string(PrinterCapabilitiesSupportedFinishingsstaple),
		string(PrinterCapabilitiesSupportedFinishingsstapleBottomLeft),
		string(PrinterCapabilitiesSupportedFinishingsstapleBottomRight),
		string(PrinterCapabilitiesSupportedFinishingsstapleDualBottom),
		string(PrinterCapabilitiesSupportedFinishingsstapleDualLeft),
		string(PrinterCapabilitiesSupportedFinishingsstapleDualRight),
		string(PrinterCapabilitiesSupportedFinishingsstapleDualTop),
		string(PrinterCapabilitiesSupportedFinishingsstapleTopLeft),
		string(PrinterCapabilitiesSupportedFinishingsstapleTopRight),
		string(PrinterCapabilitiesSupportedFinishingsstitchBottomEdge),
		string(PrinterCapabilitiesSupportedFinishingsstitchEdge),
		string(PrinterCapabilitiesSupportedFinishingsstitchLeftEdge),
		string(PrinterCapabilitiesSupportedFinishingsstitchRightEdge),
		string(PrinterCapabilitiesSupportedFinishingsstitchTopEdge),
	}
}

func parsePrinterCapabilitiesSupportedFinishings(input string) (*PrinterCapabilitiesSupportedFinishings, error) {
	vals := map[string]PrinterCapabilitiesSupportedFinishings{
		"bind":              PrinterCapabilitiesSupportedFinishingsbind,
		"cover":             PrinterCapabilitiesSupportedFinishingscover,
		"none":              PrinterCapabilitiesSupportedFinishingsnone,
		"punch":             PrinterCapabilitiesSupportedFinishingspunch,
		"saddlestitch":      PrinterCapabilitiesSupportedFinishingssaddleStitch,
		"staple":            PrinterCapabilitiesSupportedFinishingsstaple,
		"staplebottomleft":  PrinterCapabilitiesSupportedFinishingsstapleBottomLeft,
		"staplebottomright": PrinterCapabilitiesSupportedFinishingsstapleBottomRight,
		"stapledualbottom":  PrinterCapabilitiesSupportedFinishingsstapleDualBottom,
		"stapledualleft":    PrinterCapabilitiesSupportedFinishingsstapleDualLeft,
		"stapledualright":   PrinterCapabilitiesSupportedFinishingsstapleDualRight,
		"stapledualtop":     PrinterCapabilitiesSupportedFinishingsstapleDualTop,
		"stapletopleft":     PrinterCapabilitiesSupportedFinishingsstapleTopLeft,
		"stapletopright":    PrinterCapabilitiesSupportedFinishingsstapleTopRight,
		"stitchbottomedge":  PrinterCapabilitiesSupportedFinishingsstitchBottomEdge,
		"stitchedge":        PrinterCapabilitiesSupportedFinishingsstitchEdge,
		"stitchleftedge":    PrinterCapabilitiesSupportedFinishingsstitchLeftEdge,
		"stitchrightedge":   PrinterCapabilitiesSupportedFinishingsstitchRightEdge,
		"stitchtopedge":     PrinterCapabilitiesSupportedFinishingsstitchTopEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedFinishings(input)
	return &out, nil
}
