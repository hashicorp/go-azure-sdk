package featuresupport

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SupportStatus string

const (
	SupportStatusDefaultOFF   SupportStatus = "DefaultOFF"
	SupportStatusDefaultON    SupportStatus = "DefaultON"
	SupportStatusInvalid      SupportStatus = "Invalid"
	SupportStatusNotSupported SupportStatus = "NotSupported"
	SupportStatusSupported    SupportStatus = "Supported"
)

func PossibleValuesForSupportStatus() []string {
	return []string{
		string(SupportStatusDefaultOFF),
		string(SupportStatusDefaultON),
		string(SupportStatusInvalid),
		string(SupportStatusNotSupported),
		string(SupportStatusSupported),
	}
}
