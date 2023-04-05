package apidiagnostic

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlwaysLog string

const (
	AlwaysLogAllErrors AlwaysLog = "allErrors"
)

func PossibleValuesForAlwaysLog() []string {
	return []string{
		string(AlwaysLogAllErrors),
	}
}

type DataMaskingMode string

const (
	DataMaskingModeHide DataMaskingMode = "Hide"
	DataMaskingModeMask DataMaskingMode = "Mask"
)

func PossibleValuesForDataMaskingMode() []string {
	return []string{
		string(DataMaskingModeHide),
		string(DataMaskingModeMask),
	}
}

type HTTPCorrelationProtocol string

const (
	HTTPCorrelationProtocolLegacy  HTTPCorrelationProtocol = "Legacy"
	HTTPCorrelationProtocolNone    HTTPCorrelationProtocol = "None"
	HTTPCorrelationProtocolWThreeC HTTPCorrelationProtocol = "W3C"
)

func PossibleValuesForHTTPCorrelationProtocol() []string {
	return []string{
		string(HTTPCorrelationProtocolLegacy),
		string(HTTPCorrelationProtocolNone),
		string(HTTPCorrelationProtocolWThreeC),
	}
}

type OperationNameFormat string

const (
	OperationNameFormatName OperationNameFormat = "Name"
	OperationNameFormatUrl  OperationNameFormat = "Url"
)

func PossibleValuesForOperationNameFormat() []string {
	return []string{
		string(OperationNameFormatName),
		string(OperationNameFormatUrl),
	}
}

type SamplingType string

const (
	SamplingTypeFixed SamplingType = "fixed"
)

func PossibleValuesForSamplingType() []string {
	return []string{
		string(SamplingTypeFixed),
	}
}

type Verbosity string

const (
	VerbosityError       Verbosity = "error"
	VerbosityInformation Verbosity = "information"
	VerbosityVerbose     Verbosity = "verbose"
)

func PossibleValuesForVerbosity() []string {
	return []string{
		string(VerbosityError),
		string(VerbosityInformation),
		string(VerbosityVerbose),
	}
}
