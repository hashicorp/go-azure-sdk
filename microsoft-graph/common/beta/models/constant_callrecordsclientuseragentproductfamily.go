package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsClientUserAgentProductFamily string

const (
	CallRecordsClientUserAgentProductFamilyazureCommunicationServices CallRecordsClientUserAgentProductFamily = "AzureCommunicationServices"
	CallRecordsClientUserAgentProductFamilylync                       CallRecordsClientUserAgentProductFamily = "Lync"
	CallRecordsClientUserAgentProductFamilyskypeForBusiness           CallRecordsClientUserAgentProductFamily = "SkypeForBusiness"
	CallRecordsClientUserAgentProductFamilyteams                      CallRecordsClientUserAgentProductFamily = "Teams"
	CallRecordsClientUserAgentProductFamilyunknown                    CallRecordsClientUserAgentProductFamily = "Unknown"
)

func PossibleValuesForCallRecordsClientUserAgentProductFamily() []string {
	return []string{
		string(CallRecordsClientUserAgentProductFamilyazureCommunicationServices),
		string(CallRecordsClientUserAgentProductFamilylync),
		string(CallRecordsClientUserAgentProductFamilyskypeForBusiness),
		string(CallRecordsClientUserAgentProductFamilyteams),
		string(CallRecordsClientUserAgentProductFamilyunknown),
	}
}

func parseCallRecordsClientUserAgentProductFamily(input string) (*CallRecordsClientUserAgentProductFamily, error) {
	vals := map[string]CallRecordsClientUserAgentProductFamily{
		"azurecommunicationservices": CallRecordsClientUserAgentProductFamilyazureCommunicationServices,
		"lync":                       CallRecordsClientUserAgentProductFamilylync,
		"skypeforbusiness":           CallRecordsClientUserAgentProductFamilyskypeForBusiness,
		"teams":                      CallRecordsClientUserAgentProductFamilyteams,
		"unknown":                    CallRecordsClientUserAgentProductFamilyunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsClientUserAgentProductFamily(input)
	return &out, nil
}
