package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDetectionAction struct {
	AlertTemplate       *AlertTemplate               `json:"alertTemplate,omitempty"`
	ODataType           *string                      `json:"@odata.type,omitempty"`
	OrganizationalScope *SecurityOrganizationalScope `json:"organizationalScope,omitempty"`
	ResponseActions     *[]SecurityResponseAction    `json:"responseActions,omitempty"`
}
