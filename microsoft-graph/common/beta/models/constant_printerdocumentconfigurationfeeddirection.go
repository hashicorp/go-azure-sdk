package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfigurationFeedDirection string

const (
	PrinterDocumentConfigurationFeedDirectionlongEdgeFirst  PrinterDocumentConfigurationFeedDirection = "LongEdgeFirst"
	PrinterDocumentConfigurationFeedDirectionshortEdgeFirst PrinterDocumentConfigurationFeedDirection = "ShortEdgeFirst"
)

func PossibleValuesForPrinterDocumentConfigurationFeedDirection() []string {
	return []string{
		string(PrinterDocumentConfigurationFeedDirectionlongEdgeFirst),
		string(PrinterDocumentConfigurationFeedDirectionshortEdgeFirst),
	}
}

func parsePrinterDocumentConfigurationFeedDirection(input string) (*PrinterDocumentConfigurationFeedDirection, error) {
	vals := map[string]PrinterDocumentConfigurationFeedDirection{
		"longedgefirst":  PrinterDocumentConfigurationFeedDirectionlongEdgeFirst,
		"shortedgefirst": PrinterDocumentConfigurationFeedDirectionshortEdgeFirst,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterDocumentConfigurationFeedDirection(input)
	return &out, nil
}
