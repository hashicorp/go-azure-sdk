package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfigurationFinishings string

const (
	PrinterDocumentConfigurationFinishingsbind              PrinterDocumentConfigurationFinishings = "Bind"
	PrinterDocumentConfigurationFinishingscover             PrinterDocumentConfigurationFinishings = "Cover"
	PrinterDocumentConfigurationFinishingsnone              PrinterDocumentConfigurationFinishings = "None"
	PrinterDocumentConfigurationFinishingspunch             PrinterDocumentConfigurationFinishings = "Punch"
	PrinterDocumentConfigurationFinishingssaddleStitch      PrinterDocumentConfigurationFinishings = "SaddleStitch"
	PrinterDocumentConfigurationFinishingsstaple            PrinterDocumentConfigurationFinishings = "Staple"
	PrinterDocumentConfigurationFinishingsstapleBottomLeft  PrinterDocumentConfigurationFinishings = "StapleBottomLeft"
	PrinterDocumentConfigurationFinishingsstapleBottomRight PrinterDocumentConfigurationFinishings = "StapleBottomRight"
	PrinterDocumentConfigurationFinishingsstapleDualBottom  PrinterDocumentConfigurationFinishings = "StapleDualBottom"
	PrinterDocumentConfigurationFinishingsstapleDualLeft    PrinterDocumentConfigurationFinishings = "StapleDualLeft"
	PrinterDocumentConfigurationFinishingsstapleDualRight   PrinterDocumentConfigurationFinishings = "StapleDualRight"
	PrinterDocumentConfigurationFinishingsstapleDualTop     PrinterDocumentConfigurationFinishings = "StapleDualTop"
	PrinterDocumentConfigurationFinishingsstapleTopLeft     PrinterDocumentConfigurationFinishings = "StapleTopLeft"
	PrinterDocumentConfigurationFinishingsstapleTopRight    PrinterDocumentConfigurationFinishings = "StapleTopRight"
	PrinterDocumentConfigurationFinishingsstitchBottomEdge  PrinterDocumentConfigurationFinishings = "StitchBottomEdge"
	PrinterDocumentConfigurationFinishingsstitchEdge        PrinterDocumentConfigurationFinishings = "StitchEdge"
	PrinterDocumentConfigurationFinishingsstitchLeftEdge    PrinterDocumentConfigurationFinishings = "StitchLeftEdge"
	PrinterDocumentConfigurationFinishingsstitchRightEdge   PrinterDocumentConfigurationFinishings = "StitchRightEdge"
	PrinterDocumentConfigurationFinishingsstitchTopEdge     PrinterDocumentConfigurationFinishings = "StitchTopEdge"
)

func PossibleValuesForPrinterDocumentConfigurationFinishings() []string {
	return []string{
		string(PrinterDocumentConfigurationFinishingsbind),
		string(PrinterDocumentConfigurationFinishingscover),
		string(PrinterDocumentConfigurationFinishingsnone),
		string(PrinterDocumentConfigurationFinishingspunch),
		string(PrinterDocumentConfigurationFinishingssaddleStitch),
		string(PrinterDocumentConfigurationFinishingsstaple),
		string(PrinterDocumentConfigurationFinishingsstapleBottomLeft),
		string(PrinterDocumentConfigurationFinishingsstapleBottomRight),
		string(PrinterDocumentConfigurationFinishingsstapleDualBottom),
		string(PrinterDocumentConfigurationFinishingsstapleDualLeft),
		string(PrinterDocumentConfigurationFinishingsstapleDualRight),
		string(PrinterDocumentConfigurationFinishingsstapleDualTop),
		string(PrinterDocumentConfigurationFinishingsstapleTopLeft),
		string(PrinterDocumentConfigurationFinishingsstapleTopRight),
		string(PrinterDocumentConfigurationFinishingsstitchBottomEdge),
		string(PrinterDocumentConfigurationFinishingsstitchEdge),
		string(PrinterDocumentConfigurationFinishingsstitchLeftEdge),
		string(PrinterDocumentConfigurationFinishingsstitchRightEdge),
		string(PrinterDocumentConfigurationFinishingsstitchTopEdge),
	}
}

func parsePrinterDocumentConfigurationFinishings(input string) (*PrinterDocumentConfigurationFinishings, error) {
	vals := map[string]PrinterDocumentConfigurationFinishings{
		"bind":              PrinterDocumentConfigurationFinishingsbind,
		"cover":             PrinterDocumentConfigurationFinishingscover,
		"none":              PrinterDocumentConfigurationFinishingsnone,
		"punch":             PrinterDocumentConfigurationFinishingspunch,
		"saddlestitch":      PrinterDocumentConfigurationFinishingssaddleStitch,
		"staple":            PrinterDocumentConfigurationFinishingsstaple,
		"staplebottomleft":  PrinterDocumentConfigurationFinishingsstapleBottomLeft,
		"staplebottomright": PrinterDocumentConfigurationFinishingsstapleBottomRight,
		"stapledualbottom":  PrinterDocumentConfigurationFinishingsstapleDualBottom,
		"stapledualleft":    PrinterDocumentConfigurationFinishingsstapleDualLeft,
		"stapledualright":   PrinterDocumentConfigurationFinishingsstapleDualRight,
		"stapledualtop":     PrinterDocumentConfigurationFinishingsstapleDualTop,
		"stapletopleft":     PrinterDocumentConfigurationFinishingsstapleTopLeft,
		"stapletopright":    PrinterDocumentConfigurationFinishingsstapleTopRight,
		"stitchbottomedge":  PrinterDocumentConfigurationFinishingsstitchBottomEdge,
		"stitchedge":        PrinterDocumentConfigurationFinishingsstitchEdge,
		"stitchleftedge":    PrinterDocumentConfigurationFinishingsstitchLeftEdge,
		"stitchrightedge":   PrinterDocumentConfigurationFinishingsstitchRightEdge,
		"stitchtopedge":     PrinterDocumentConfigurationFinishingsstitchTopEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationFinishings(input)
	return &out, nil
}
