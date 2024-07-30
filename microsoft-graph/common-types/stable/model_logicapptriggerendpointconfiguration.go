package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogicAppTriggerEndpointConfiguration struct {
	LogicAppWorkflowName *string `json:"logicAppWorkflowName,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	ResourceGroupName    *string `json:"resourceGroupName,omitempty"`
	SubscriptionId       *string `json:"subscriptionId,omitempty"`
	Url                  *string `json:"url,omitempty"`
}
