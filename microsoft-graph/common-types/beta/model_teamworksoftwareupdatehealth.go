package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkSoftwareUpdateHealth struct {
	AdminAgentSoftwareUpdateStatus      *TeamworkSoftwareUpdateStatus `json:"adminAgentSoftwareUpdateStatus,omitempty"`
	CompanyPortalSoftwareUpdateStatus   *TeamworkSoftwareUpdateStatus `json:"companyPortalSoftwareUpdateStatus,omitempty"`
	FirmwareSoftwareUpdateStatus        *TeamworkSoftwareUpdateStatus `json:"firmwareSoftwareUpdateStatus,omitempty"`
	ODataType                           *string                       `json:"@odata.type,omitempty"`
	OperatingSystemSoftwareUpdateStatus *TeamworkSoftwareUpdateStatus `json:"operatingSystemSoftwareUpdateStatus,omitempty"`
	PartnerAgentSoftwareUpdateStatus    *TeamworkSoftwareUpdateStatus `json:"partnerAgentSoftwareUpdateStatus,omitempty"`
	TeamsClientSoftwareUpdateStatus     *TeamworkSoftwareUpdateStatus `json:"teamsClientSoftwareUpdateStatus,omitempty"`
}
