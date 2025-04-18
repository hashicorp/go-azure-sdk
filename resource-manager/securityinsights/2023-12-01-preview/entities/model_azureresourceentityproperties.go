package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureResourceEntityProperties struct {
	AdditionalData *interface{} `json:"additionalData,omitempty"`
	FriendlyName   *string      `json:"friendlyName,omitempty"`
	ResourceId     *string      `json:"resourceId,omitempty"`
	SubscriptionId *string      `json:"subscriptionId,omitempty"`
}
