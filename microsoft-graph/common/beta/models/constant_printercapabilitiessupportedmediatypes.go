package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesSupportedMediaTypes string

const (
	PrinterCapabilitiesSupportedMediaTypescontinuous      PrinterCapabilitiesSupportedMediaTypes = "Continuous"
	PrinterCapabilitiesSupportedMediaTypescontinuousLong  PrinterCapabilitiesSupportedMediaTypes = "ContinuousLong"
	PrinterCapabilitiesSupportedMediaTypescontinuousShort PrinterCapabilitiesSupportedMediaTypes = "ContinuousShort"
	PrinterCapabilitiesSupportedMediaTypesenvelope        PrinterCapabilitiesSupportedMediaTypes = "Envelope"
	PrinterCapabilitiesSupportedMediaTypesenvelopePlain   PrinterCapabilitiesSupportedMediaTypes = "EnvelopePlain"
	PrinterCapabilitiesSupportedMediaTypesenvelopeWindow  PrinterCapabilitiesSupportedMediaTypes = "EnvelopeWindow"
	PrinterCapabilitiesSupportedMediaTypeslabels          PrinterCapabilitiesSupportedMediaTypes = "Labels"
	PrinterCapabilitiesSupportedMediaTypesmultiLayer      PrinterCapabilitiesSupportedMediaTypes = "MultiLayer"
	PrinterCapabilitiesSupportedMediaTypesmultiPartForm   PrinterCapabilitiesSupportedMediaTypes = "MultiPartForm"
	PrinterCapabilitiesSupportedMediaTypesscreen          PrinterCapabilitiesSupportedMediaTypes = "Screen"
	PrinterCapabilitiesSupportedMediaTypesscreenPaged     PrinterCapabilitiesSupportedMediaTypes = "ScreenPaged"
	PrinterCapabilitiesSupportedMediaTypesstationery      PrinterCapabilitiesSupportedMediaTypes = "Stationery"
	PrinterCapabilitiesSupportedMediaTypestransparency    PrinterCapabilitiesSupportedMediaTypes = "Transparency"
)

func PossibleValuesForPrinterCapabilitiesSupportedMediaTypes() []string {
	return []string{
		string(PrinterCapabilitiesSupportedMediaTypescontinuous),
		string(PrinterCapabilitiesSupportedMediaTypescontinuousLong),
		string(PrinterCapabilitiesSupportedMediaTypescontinuousShort),
		string(PrinterCapabilitiesSupportedMediaTypesenvelope),
		string(PrinterCapabilitiesSupportedMediaTypesenvelopePlain),
		string(PrinterCapabilitiesSupportedMediaTypesenvelopeWindow),
		string(PrinterCapabilitiesSupportedMediaTypeslabels),
		string(PrinterCapabilitiesSupportedMediaTypesmultiLayer),
		string(PrinterCapabilitiesSupportedMediaTypesmultiPartForm),
		string(PrinterCapabilitiesSupportedMediaTypesscreen),
		string(PrinterCapabilitiesSupportedMediaTypesscreenPaged),
		string(PrinterCapabilitiesSupportedMediaTypesstationery),
		string(PrinterCapabilitiesSupportedMediaTypestransparency),
	}
}

func parsePrinterCapabilitiesSupportedMediaTypes(input string) (*PrinterCapabilitiesSupportedMediaTypes, error) {
	vals := map[string]PrinterCapabilitiesSupportedMediaTypes{
		"continuous":      PrinterCapabilitiesSupportedMediaTypescontinuous,
		"continuouslong":  PrinterCapabilitiesSupportedMediaTypescontinuousLong,
		"continuousshort": PrinterCapabilitiesSupportedMediaTypescontinuousShort,
		"envelope":        PrinterCapabilitiesSupportedMediaTypesenvelope,
		"envelopeplain":   PrinterCapabilitiesSupportedMediaTypesenvelopePlain,
		"envelopewindow":  PrinterCapabilitiesSupportedMediaTypesenvelopeWindow,
		"labels":          PrinterCapabilitiesSupportedMediaTypeslabels,
		"multilayer":      PrinterCapabilitiesSupportedMediaTypesmultiLayer,
		"multipartform":   PrinterCapabilitiesSupportedMediaTypesmultiPartForm,
		"screen":          PrinterCapabilitiesSupportedMediaTypesscreen,
		"screenpaged":     PrinterCapabilitiesSupportedMediaTypesscreenPaged,
		"stationery":      PrinterCapabilitiesSupportedMediaTypesstationery,
		"transparency":    PrinterCapabilitiesSupportedMediaTypestransparency,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterCapabilitiesSupportedMediaTypes(input)
	return &out, nil
}
