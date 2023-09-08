package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfigurationFinishings string

const (
	PrintJobConfigurationFinishingsbind              PrintJobConfigurationFinishings = "Bind"
	PrintJobConfigurationFinishingscover             PrintJobConfigurationFinishings = "Cover"
	PrintJobConfigurationFinishingsnone              PrintJobConfigurationFinishings = "None"
	PrintJobConfigurationFinishingspunch             PrintJobConfigurationFinishings = "Punch"
	PrintJobConfigurationFinishingssaddleStitch      PrintJobConfigurationFinishings = "SaddleStitch"
	PrintJobConfigurationFinishingsstaple            PrintJobConfigurationFinishings = "Staple"
	PrintJobConfigurationFinishingsstapleBottomLeft  PrintJobConfigurationFinishings = "StapleBottomLeft"
	PrintJobConfigurationFinishingsstapleBottomRight PrintJobConfigurationFinishings = "StapleBottomRight"
	PrintJobConfigurationFinishingsstapleDualBottom  PrintJobConfigurationFinishings = "StapleDualBottom"
	PrintJobConfigurationFinishingsstapleDualLeft    PrintJobConfigurationFinishings = "StapleDualLeft"
	PrintJobConfigurationFinishingsstapleDualRight   PrintJobConfigurationFinishings = "StapleDualRight"
	PrintJobConfigurationFinishingsstapleDualTop     PrintJobConfigurationFinishings = "StapleDualTop"
	PrintJobConfigurationFinishingsstapleTopLeft     PrintJobConfigurationFinishings = "StapleTopLeft"
	PrintJobConfigurationFinishingsstapleTopRight    PrintJobConfigurationFinishings = "StapleTopRight"
	PrintJobConfigurationFinishingsstitchBottomEdge  PrintJobConfigurationFinishings = "StitchBottomEdge"
	PrintJobConfigurationFinishingsstitchEdge        PrintJobConfigurationFinishings = "StitchEdge"
	PrintJobConfigurationFinishingsstitchLeftEdge    PrintJobConfigurationFinishings = "StitchLeftEdge"
	PrintJobConfigurationFinishingsstitchRightEdge   PrintJobConfigurationFinishings = "StitchRightEdge"
	PrintJobConfigurationFinishingsstitchTopEdge     PrintJobConfigurationFinishings = "StitchTopEdge"
)

func PossibleValuesForPrintJobConfigurationFinishings() []string {
	return []string{
		string(PrintJobConfigurationFinishingsbind),
		string(PrintJobConfigurationFinishingscover),
		string(PrintJobConfigurationFinishingsnone),
		string(PrintJobConfigurationFinishingspunch),
		string(PrintJobConfigurationFinishingssaddleStitch),
		string(PrintJobConfigurationFinishingsstaple),
		string(PrintJobConfigurationFinishingsstapleBottomLeft),
		string(PrintJobConfigurationFinishingsstapleBottomRight),
		string(PrintJobConfigurationFinishingsstapleDualBottom),
		string(PrintJobConfigurationFinishingsstapleDualLeft),
		string(PrintJobConfigurationFinishingsstapleDualRight),
		string(PrintJobConfigurationFinishingsstapleDualTop),
		string(PrintJobConfigurationFinishingsstapleTopLeft),
		string(PrintJobConfigurationFinishingsstapleTopRight),
		string(PrintJobConfigurationFinishingsstitchBottomEdge),
		string(PrintJobConfigurationFinishingsstitchEdge),
		string(PrintJobConfigurationFinishingsstitchLeftEdge),
		string(PrintJobConfigurationFinishingsstitchRightEdge),
		string(PrintJobConfigurationFinishingsstitchTopEdge),
	}
}

func parsePrintJobConfigurationFinishings(input string) (*PrintJobConfigurationFinishings, error) {
	vals := map[string]PrintJobConfigurationFinishings{
		"bind":              PrintJobConfigurationFinishingsbind,
		"cover":             PrintJobConfigurationFinishingscover,
		"none":              PrintJobConfigurationFinishingsnone,
		"punch":             PrintJobConfigurationFinishingspunch,
		"saddlestitch":      PrintJobConfigurationFinishingssaddleStitch,
		"staple":            PrintJobConfigurationFinishingsstaple,
		"staplebottomleft":  PrintJobConfigurationFinishingsstapleBottomLeft,
		"staplebottomright": PrintJobConfigurationFinishingsstapleBottomRight,
		"stapledualbottom":  PrintJobConfigurationFinishingsstapleDualBottom,
		"stapledualleft":    PrintJobConfigurationFinishingsstapleDualLeft,
		"stapledualright":   PrintJobConfigurationFinishingsstapleDualRight,
		"stapledualtop":     PrintJobConfigurationFinishingsstapleDualTop,
		"stapletopleft":     PrintJobConfigurationFinishingsstapleTopLeft,
		"stapletopright":    PrintJobConfigurationFinishingsstapleTopRight,
		"stitchbottomedge":  PrintJobConfigurationFinishingsstitchBottomEdge,
		"stitchedge":        PrintJobConfigurationFinishingsstitchEdge,
		"stitchleftedge":    PrintJobConfigurationFinishingsstitchLeftEdge,
		"stitchrightedge":   PrintJobConfigurationFinishingsstitchRightEdge,
		"stitchtopedge":     PrintJobConfigurationFinishingsstitchTopEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationFinishings(input)
	return &out, nil
}
