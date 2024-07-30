package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerCollectionAssignmentTarget struct {
	CollectionId                               *string                                                                                   `json:"collectionId,omitempty"`
	DeviceAndAppManagementAssignmentFilterId   *string                                                                                   `json:"deviceAndAppManagementAssignmentFilterId,omitempty"`
	DeviceAndAppManagementAssignmentFilterType *ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType `json:"deviceAndAppManagementAssignmentFilterType,omitempty"`
	ODataType                                  *string                                                                                   `json:"@odata.type,omitempty"`
}
