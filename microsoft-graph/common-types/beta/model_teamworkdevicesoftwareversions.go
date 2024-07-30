package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDeviceSoftwareVersions struct {
	AdminAgentSoftwareVersion      *string `json:"adminAgentSoftwareVersion,omitempty"`
	FirmwareSoftwareVersion        *string `json:"firmwareSoftwareVersion,omitempty"`
	ODataType                      *string `json:"@odata.type,omitempty"`
	OperatingSystemSoftwareVersion *string `json:"operatingSystemSoftwareVersion,omitempty"`
	PartnerAgentSoftwareVersion    *string `json:"partnerAgentSoftwareVersion,omitempty"`
	TeamsClientSoftwareVersion     *string `json:"teamsClientSoftwareVersion,omitempty"`
}
