package servicefabrics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceFabricProperties struct {
	ApplicableSchedule      *ApplicableSchedule `json:"applicableSchedule,omitempty"`
	EnvironmentId           *string             `json:"environmentId,omitempty"`
	ExternalServiceFabricId *string             `json:"externalServiceFabricId,omitempty"`
	ProvisioningState       *string             `json:"provisioningState,omitempty"`
	UniqueIdentifier        *string             `json:"uniqueIdentifier,omitempty"`
}
