package channels

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceMoveChangeHistory struct {
	AzureSubscriptionId *string `json:"azureSubscriptionId,omitempty"`
	ChangedTimeUtc      *string `json:"changedTimeUtc,omitempty"`
	ResourceGroupName   *string `json:"resourceGroupName,omitempty"`
}
