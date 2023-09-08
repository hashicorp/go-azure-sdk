package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCPartnerAgentInstallResult struct {
	InstallStatus       *CloudPCPartnerAgentInstallResultInstallStatus    `json:"installStatus,omitempty"`
	IsThirdPartyPartner *bool                                             `json:"isThirdPartyPartner,omitempty"`
	ODataType           *string                                           `json:"@odata.type,omitempty"`
	PartnerAgentName    *CloudPCPartnerAgentInstallResultPartnerAgentName `json:"partnerAgentName,omitempty"`
	Retriable           *bool                                             `json:"retriable,omitempty"`
}
