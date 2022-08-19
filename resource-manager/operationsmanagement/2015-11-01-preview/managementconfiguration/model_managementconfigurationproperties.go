package managementconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementConfigurationProperties struct {
	ApplicationId      *string                `json:"applicationId,omitempty"`
	Parameters         []ArmTemplateParameter `json:"parameters"`
	ParentResourceType string                 `json:"parentResourceType"`
	ProvisioningState  *string                `json:"provisioningState,omitempty"`
	Template           interface{}            `json:"template"`
}
