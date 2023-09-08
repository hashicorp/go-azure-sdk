package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsFinishings string

const (
	PrinterDefaultsFinishingsbind              PrinterDefaultsFinishings = "Bind"
	PrinterDefaultsFinishingscover             PrinterDefaultsFinishings = "Cover"
	PrinterDefaultsFinishingsnone              PrinterDefaultsFinishings = "None"
	PrinterDefaultsFinishingspunch             PrinterDefaultsFinishings = "Punch"
	PrinterDefaultsFinishingssaddleStitch      PrinterDefaultsFinishings = "SaddleStitch"
	PrinterDefaultsFinishingsstaple            PrinterDefaultsFinishings = "Staple"
	PrinterDefaultsFinishingsstapleBottomLeft  PrinterDefaultsFinishings = "StapleBottomLeft"
	PrinterDefaultsFinishingsstapleBottomRight PrinterDefaultsFinishings = "StapleBottomRight"
	PrinterDefaultsFinishingsstapleDualBottom  PrinterDefaultsFinishings = "StapleDualBottom"
	PrinterDefaultsFinishingsstapleDualLeft    PrinterDefaultsFinishings = "StapleDualLeft"
	PrinterDefaultsFinishingsstapleDualRight   PrinterDefaultsFinishings = "StapleDualRight"
	PrinterDefaultsFinishingsstapleDualTop     PrinterDefaultsFinishings = "StapleDualTop"
	PrinterDefaultsFinishingsstapleTopLeft     PrinterDefaultsFinishings = "StapleTopLeft"
	PrinterDefaultsFinishingsstapleTopRight    PrinterDefaultsFinishings = "StapleTopRight"
	PrinterDefaultsFinishingsstitchBottomEdge  PrinterDefaultsFinishings = "StitchBottomEdge"
	PrinterDefaultsFinishingsstitchEdge        PrinterDefaultsFinishings = "StitchEdge"
	PrinterDefaultsFinishingsstitchLeftEdge    PrinterDefaultsFinishings = "StitchLeftEdge"
	PrinterDefaultsFinishingsstitchRightEdge   PrinterDefaultsFinishings = "StitchRightEdge"
	PrinterDefaultsFinishingsstitchTopEdge     PrinterDefaultsFinishings = "StitchTopEdge"
)

func PossibleValuesForPrinterDefaultsFinishings() []string {
	return []string{
		string(PrinterDefaultsFinishingsbind),
		string(PrinterDefaultsFinishingscover),
		string(PrinterDefaultsFinishingsnone),
		string(PrinterDefaultsFinishingspunch),
		string(PrinterDefaultsFinishingssaddleStitch),
		string(PrinterDefaultsFinishingsstaple),
		string(PrinterDefaultsFinishingsstapleBottomLeft),
		string(PrinterDefaultsFinishingsstapleBottomRight),
		string(PrinterDefaultsFinishingsstapleDualBottom),
		string(PrinterDefaultsFinishingsstapleDualLeft),
		string(PrinterDefaultsFinishingsstapleDualRight),
		string(PrinterDefaultsFinishingsstapleDualTop),
		string(PrinterDefaultsFinishingsstapleTopLeft),
		string(PrinterDefaultsFinishingsstapleTopRight),
		string(PrinterDefaultsFinishingsstitchBottomEdge),
		string(PrinterDefaultsFinishingsstitchEdge),
		string(PrinterDefaultsFinishingsstitchLeftEdge),
		string(PrinterDefaultsFinishingsstitchRightEdge),
		string(PrinterDefaultsFinishingsstitchTopEdge),
	}
}

func parsePrinterDefaultsFinishings(input string) (*PrinterDefaultsFinishings, error) {
	vals := map[string]PrinterDefaultsFinishings{
		"bind":              PrinterDefaultsFinishingsbind,
		"cover":             PrinterDefaultsFinishingscover,
		"none":              PrinterDefaultsFinishingsnone,
		"punch":             PrinterDefaultsFinishingspunch,
		"saddlestitch":      PrinterDefaultsFinishingssaddleStitch,
		"staple":            PrinterDefaultsFinishingsstaple,
		"staplebottomleft":  PrinterDefaultsFinishingsstapleBottomLeft,
		"staplebottomright": PrinterDefaultsFinishingsstapleBottomRight,
		"stapledualbottom":  PrinterDefaultsFinishingsstapleDualBottom,
		"stapledualleft":    PrinterDefaultsFinishingsstapleDualLeft,
		"stapledualright":   PrinterDefaultsFinishingsstapleDualRight,
		"stapledualtop":     PrinterDefaultsFinishingsstapleDualTop,
		"stapletopleft":     PrinterDefaultsFinishingsstapleTopLeft,
		"stapletopright":    PrinterDefaultsFinishingsstapleTopRight,
		"stitchbottomedge":  PrinterDefaultsFinishingsstitchBottomEdge,
		"stitchedge":        PrinterDefaultsFinishingsstitchEdge,
		"stitchleftedge":    PrinterDefaultsFinishingsstitchLeftEdge,
		"stitchrightedge":   PrinterDefaultsFinishingsstitchRightEdge,
		"stitchtopedge":     PrinterDefaultsFinishingsstitchTopEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDefaultsFinishings(input)
	return &out, nil
}
