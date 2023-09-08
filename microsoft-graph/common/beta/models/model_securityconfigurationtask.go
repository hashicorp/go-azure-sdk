package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityConfigurationTask struct {
	ApplicablePlatform            *SecurityConfigurationTaskApplicablePlatform            `json:"applicablePlatform,omitempty"`
	AssignedTo                    *string                                                 `json:"assignedTo,omitempty"`
	Category                      *SecurityConfigurationTaskCategory                      `json:"category,omitempty"`
	CreatedDateTime               *string                                                 `json:"createdDateTime,omitempty"`
	Creator                       *string                                                 `json:"creator,omitempty"`
	CreatorNotes                  *string                                                 `json:"creatorNotes,omitempty"`
	Description                   *string                                                 `json:"description,omitempty"`
	DisplayName                   *string                                                 `json:"displayName,omitempty"`
	DueDateTime                   *string                                                 `json:"dueDateTime,omitempty"`
	EndpointSecurityPolicy        *SecurityConfigurationTaskEndpointSecurityPolicy        `json:"endpointSecurityPolicy,omitempty"`
	EndpointSecurityPolicyProfile *SecurityConfigurationTaskEndpointSecurityPolicyProfile `json:"endpointSecurityPolicyProfile,omitempty"`
	Id                            *string                                                 `json:"id,omitempty"`
	Insights                      *string                                                 `json:"insights,omitempty"`
	IntendedSettings              *[]KeyValuePair                                         `json:"intendedSettings,omitempty"`
	ManagedDeviceCount            *int64                                                  `json:"managedDeviceCount,omitempty"`
	ManagedDevices                *[]VulnerableManagedDevice                              `json:"managedDevices,omitempty"`
	ODataType                     *string                                                 `json:"@odata.type,omitempty"`
	Priority                      *SecurityConfigurationTaskPriority                      `json:"priority,omitempty"`
	Status                        *SecurityConfigurationTaskStatus                        `json:"status,omitempty"`
}
