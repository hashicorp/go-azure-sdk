package jitnetworkaccesspolicies

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Protocol string

const (
	ProtocolAny Protocol = "*"
	ProtocolTCP Protocol = "TCP"
	ProtocolUDP Protocol = "UDP"
)

func PossibleValuesForProtocol() []string {
	return []string{
		string(ProtocolAny),
		string(ProtocolTCP),
		string(ProtocolUDP),
	}
}

func parseProtocol(input string) (*Protocol, error) {
	vals := map[string]Protocol{
		"*":   ProtocolAny,
		"tcp": ProtocolTCP,
		"udp": ProtocolUDP,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Protocol(input)
	return &out, nil
}

type Status string

const (
	StatusInitiated Status = "Initiated"
	StatusRevoked   Status = "Revoked"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusInitiated),
		string(StatusRevoked),
	}
}

func parseStatus(input string) (*Status, error) {
	vals := map[string]Status{
		"initiated": StatusInitiated,
		"revoked":   StatusRevoked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Status(input)
	return &out, nil
}

type StatusReason string

const (
	StatusReasonExpired               StatusReason = "Expired"
	StatusReasonNewerRequestInitiated StatusReason = "NewerRequestInitiated"
	StatusReasonUserRequested         StatusReason = "UserRequested"
)

func PossibleValuesForStatusReason() []string {
	return []string{
		string(StatusReasonExpired),
		string(StatusReasonNewerRequestInitiated),
		string(StatusReasonUserRequested),
	}
}

func parseStatusReason(input string) (*StatusReason, error) {
	vals := map[string]StatusReason{
		"expired":               StatusReasonExpired,
		"newerrequestinitiated": StatusReasonNewerRequestInitiated,
		"userrequested":         StatusReasonUserRequested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusReason(input)
	return &out, nil
}
