package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTunnelServerLogCollectionResponseStatus string

const (
	MicrosoftTunnelServerLogCollectionResponseStatuscompleted MicrosoftTunnelServerLogCollectionResponseStatus = "Completed"
	MicrosoftTunnelServerLogCollectionResponseStatusfailed    MicrosoftTunnelServerLogCollectionResponseStatus = "Failed"
	MicrosoftTunnelServerLogCollectionResponseStatuspending   MicrosoftTunnelServerLogCollectionResponseStatus = "Pending"
)

func PossibleValuesForMicrosoftTunnelServerLogCollectionResponseStatus() []string {
	return []string{
		string(MicrosoftTunnelServerLogCollectionResponseStatuscompleted),
		string(MicrosoftTunnelServerLogCollectionResponseStatusfailed),
		string(MicrosoftTunnelServerLogCollectionResponseStatuspending),
	}
}

func parseMicrosoftTunnelServerLogCollectionResponseStatus(input string) (*MicrosoftTunnelServerLogCollectionResponseStatus, error) {
	vals := map[string]MicrosoftTunnelServerLogCollectionResponseStatus{
		"completed": MicrosoftTunnelServerLogCollectionResponseStatuscompleted,
		"failed":    MicrosoftTunnelServerLogCollectionResponseStatusfailed,
		"pending":   MicrosoftTunnelServerLogCollectionResponseStatuspending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftTunnelServerLogCollectionResponseStatus(input)
	return &out, nil
}
