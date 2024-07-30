package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedMobileApp struct {
	Id                  *string              `json:"id,omitempty"`
	MobileAppIdentifier *MobileAppIdentifier `json:"mobileAppIdentifier,omitempty"`
	ODataType           *string              `json:"@odata.type,omitempty"`
	Version             *string              `json:"version,omitempty"`
}
