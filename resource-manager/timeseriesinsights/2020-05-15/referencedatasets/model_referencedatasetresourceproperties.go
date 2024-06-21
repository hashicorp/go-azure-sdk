package referencedatasets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReferenceDataSetResourceProperties struct {
	CreationTime                 *string                       `json:"creationTime,omitempty"`
	DataStringComparisonBehavior *DataStringComparisonBehavior `json:"dataStringComparisonBehavior,omitempty"`
	KeyProperties                []ReferenceDataSetKeyProperty `json:"keyProperties"`
	ProvisioningState            *ProvisioningState            `json:"provisioningState,omitempty"`
}
