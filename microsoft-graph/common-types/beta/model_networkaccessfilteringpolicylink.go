package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessFilteringPolicyLink struct {
	CreatedDateTime      *string                                       `json:"createdDateTime,omitempty"`
	Id                   *string                                       `json:"id,omitempty"`
	LastModifiedDateTime *string                                       `json:"lastModifiedDateTime,omitempty"`
	LoggingState         *NetworkaccessFilteringPolicyLinkLoggingState `json:"loggingState,omitempty"`
	ODataType            *string                                       `json:"@odata.type,omitempty"`
	Policy               *NetworkaccessPolicy                          `json:"policy,omitempty"`
	Priority             *int64                                        `json:"priority,omitempty"`
	State                *NetworkaccessFilteringPolicyLinkState        `json:"state,omitempty"`
	Version              *string                                       `json:"version,omitempty"`
}
