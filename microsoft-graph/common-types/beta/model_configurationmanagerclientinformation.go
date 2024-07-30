package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerClientInformation struct {
	ClientIdentifier *string `json:"clientIdentifier,omitempty"`
	ClientVersion    *string `json:"clientVersion,omitempty"`
	IsBlocked        *bool   `json:"isBlocked,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
}
