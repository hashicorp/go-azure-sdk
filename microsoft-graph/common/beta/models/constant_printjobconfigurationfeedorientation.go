package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfigurationFeedOrientation string

const (
	PrintJobConfigurationFeedOrientationlongEdgeFirst  PrintJobConfigurationFeedOrientation = "LongEdgeFirst"
	PrintJobConfigurationFeedOrientationshortEdgeFirst PrintJobConfigurationFeedOrientation = "ShortEdgeFirst"
)

func PossibleValuesForPrintJobConfigurationFeedOrientation() []string {
	return []string{
		string(PrintJobConfigurationFeedOrientationlongEdgeFirst),
		string(PrintJobConfigurationFeedOrientationshortEdgeFirst),
	}
}

func parsePrintJobConfigurationFeedOrientation(input string) (*PrintJobConfigurationFeedOrientation, error) {
	vals := map[string]PrintJobConfigurationFeedOrientation{
		"longedgefirst":  PrintJobConfigurationFeedOrientationlongEdgeFirst,
		"shortedgefirst": PrintJobConfigurationFeedOrientationshortEdgeFirst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobConfigurationFeedOrientation(input)
	return &out, nil
}
