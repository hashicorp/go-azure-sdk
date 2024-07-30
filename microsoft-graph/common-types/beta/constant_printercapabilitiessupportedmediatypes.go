package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesSupportedMediaTypes string

const (
	PrinterCapabilitiesSupportedMediaTypes_Continuous      PrinterCapabilitiesSupportedMediaTypes = "continuous"
	PrinterCapabilitiesSupportedMediaTypes_ContinuousLong  PrinterCapabilitiesSupportedMediaTypes = "continuousLong"
	PrinterCapabilitiesSupportedMediaTypes_ContinuousShort PrinterCapabilitiesSupportedMediaTypes = "continuousShort"
	PrinterCapabilitiesSupportedMediaTypes_Envelope        PrinterCapabilitiesSupportedMediaTypes = "envelope"
	PrinterCapabilitiesSupportedMediaTypes_EnvelopePlain   PrinterCapabilitiesSupportedMediaTypes = "envelopePlain"
	PrinterCapabilitiesSupportedMediaTypes_EnvelopeWindow  PrinterCapabilitiesSupportedMediaTypes = "envelopeWindow"
	PrinterCapabilitiesSupportedMediaTypes_Labels          PrinterCapabilitiesSupportedMediaTypes = "labels"
	PrinterCapabilitiesSupportedMediaTypes_MultiLayer      PrinterCapabilitiesSupportedMediaTypes = "multiLayer"
	PrinterCapabilitiesSupportedMediaTypes_MultiPartForm   PrinterCapabilitiesSupportedMediaTypes = "multiPartForm"
	PrinterCapabilitiesSupportedMediaTypes_Screen          PrinterCapabilitiesSupportedMediaTypes = "screen"
	PrinterCapabilitiesSupportedMediaTypes_ScreenPaged     PrinterCapabilitiesSupportedMediaTypes = "screenPaged"
	PrinterCapabilitiesSupportedMediaTypes_Stationery      PrinterCapabilitiesSupportedMediaTypes = "stationery"
	PrinterCapabilitiesSupportedMediaTypes_Transparency    PrinterCapabilitiesSupportedMediaTypes = "transparency"
)

func PossibleValuesForPrinterCapabilitiesSupportedMediaTypes() []string {
	return []string{
		string(PrinterCapabilitiesSupportedMediaTypes_Continuous),
		string(PrinterCapabilitiesSupportedMediaTypes_ContinuousLong),
		string(PrinterCapabilitiesSupportedMediaTypes_ContinuousShort),
		string(PrinterCapabilitiesSupportedMediaTypes_Envelope),
		string(PrinterCapabilitiesSupportedMediaTypes_EnvelopePlain),
		string(PrinterCapabilitiesSupportedMediaTypes_EnvelopeWindow),
		string(PrinterCapabilitiesSupportedMediaTypes_Labels),
		string(PrinterCapabilitiesSupportedMediaTypes_MultiLayer),
		string(PrinterCapabilitiesSupportedMediaTypes_MultiPartForm),
		string(PrinterCapabilitiesSupportedMediaTypes_Screen),
		string(PrinterCapabilitiesSupportedMediaTypes_ScreenPaged),
		string(PrinterCapabilitiesSupportedMediaTypes_Stationery),
		string(PrinterCapabilitiesSupportedMediaTypes_Transparency),
	}
}

func (s *PrinterCapabilitiesSupportedMediaTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterCapabilitiesSupportedMediaTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterCapabilitiesSupportedMediaTypes(input string) (*PrinterCapabilitiesSupportedMediaTypes, error) {
	vals := map[string]PrinterCapabilitiesSupportedMediaTypes{
		"continuous":      PrinterCapabilitiesSupportedMediaTypes_Continuous,
		"continuouslong":  PrinterCapabilitiesSupportedMediaTypes_ContinuousLong,
		"continuousshort": PrinterCapabilitiesSupportedMediaTypes_ContinuousShort,
		"envelope":        PrinterCapabilitiesSupportedMediaTypes_Envelope,
		"envelopeplain":   PrinterCapabilitiesSupportedMediaTypes_EnvelopePlain,
		"envelopewindow":  PrinterCapabilitiesSupportedMediaTypes_EnvelopeWindow,
		"labels":          PrinterCapabilitiesSupportedMediaTypes_Labels,
		"multilayer":      PrinterCapabilitiesSupportedMediaTypes_MultiLayer,
		"multipartform":   PrinterCapabilitiesSupportedMediaTypes_MultiPartForm,
		"screen":          PrinterCapabilitiesSupportedMediaTypes_Screen,
		"screenpaged":     PrinterCapabilitiesSupportedMediaTypes_ScreenPaged,
		"stationery":      PrinterCapabilitiesSupportedMediaTypes_Stationery,
		"transparency":    PrinterCapabilitiesSupportedMediaTypes_Transparency,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedMediaTypes(input)
	return &out, nil
}
