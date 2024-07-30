package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiedPublisher struct {
	AddedDateTime       *string `json:"addedDateTime,omitempty"`
	DisplayName         *string `json:"displayName,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
	VerifiedPublisherId *string `json:"verifiedPublisherId,omitempty"`
}
