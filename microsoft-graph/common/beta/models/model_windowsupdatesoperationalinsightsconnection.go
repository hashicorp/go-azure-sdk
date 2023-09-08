package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesOperationalInsightsConnection struct {
	AzureResourceGroupName *string                                           `json:"azureResourceGroupName,omitempty"`
	AzureSubscriptionId    *string                                           `json:"azureSubscriptionId,omitempty"`
	Id                     *string                                           `json:"id,omitempty"`
	ODataType              *string                                           `json:"@odata.type,omitempty"`
	State                  *WindowsUpdatesOperationalInsightsConnectionState `json:"state,omitempty"`
	WorkspaceName          *string                                           `json:"workspaceName,omitempty"`
}
