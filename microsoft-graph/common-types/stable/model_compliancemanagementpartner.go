package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceManagementPartner struct {
	AndroidEnrollmentAssignments *[]ComplianceManagementPartnerAssignment `json:"androidEnrollmentAssignments,omitempty"`
	AndroidOnboarded             *bool                                    `json:"androidOnboarded,omitempty"`
	DisplayName                  *string                                  `json:"displayName,omitempty"`
	Id                           *string                                  `json:"id,omitempty"`
	IosEnrollmentAssignments     *[]ComplianceManagementPartnerAssignment `json:"iosEnrollmentAssignments,omitempty"`
	IosOnboarded                 *bool                                    `json:"iosOnboarded,omitempty"`
	LastHeartbeatDateTime        *string                                  `json:"lastHeartbeatDateTime,omitempty"`
	MacOsEnrollmentAssignments   *[]ComplianceManagementPartnerAssignment `json:"macOsEnrollmentAssignments,omitempty"`
	MacOsOnboarded               *bool                                    `json:"macOsOnboarded,omitempty"`
	ODataType                    *string                                  `json:"@odata.type,omitempty"`
	PartnerState                 *ComplianceManagementPartnerPartnerState `json:"partnerState,omitempty"`
}
