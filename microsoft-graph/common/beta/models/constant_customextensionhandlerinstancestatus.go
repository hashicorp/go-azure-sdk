package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionHandlerInstanceStatus string

const (
	CustomExtensionHandlerInstanceStatusrequestReceived CustomExtensionHandlerInstanceStatus = "RequestReceived"
	CustomExtensionHandlerInstanceStatusrequestSent     CustomExtensionHandlerInstanceStatus = "RequestSent"
)

func PossibleValuesForCustomExtensionHandlerInstanceStatus() []string {
	return []string{
		string(CustomExtensionHandlerInstanceStatusrequestReceived),
		string(CustomExtensionHandlerInstanceStatusrequestSent),
	}
}

func parseCustomExtensionHandlerInstanceStatus(input string) (*CustomExtensionHandlerInstanceStatus, error) {
	vals := map[string]CustomExtensionHandlerInstanceStatus{
		"requestreceived": CustomExtensionHandlerInstanceStatusrequestReceived,
		"requestsent":     CustomExtensionHandlerInstanceStatusrequestSent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomExtensionHandlerInstanceStatus(input)
	return &out, nil
}
