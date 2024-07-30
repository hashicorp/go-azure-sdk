package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySecurity struct {
	Id                    *string                        `json:"id,omitempty"`
	InformationProtection *SecurityInformationProtection `json:"informationProtection,omitempty"`
	ODataType             *string                        `json:"@odata.type,omitempty"`
}
