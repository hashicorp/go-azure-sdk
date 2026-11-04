package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FreeLimitExhaustionBehaviorCapability struct {
	ExhaustionBehaviorType *FreeLimitExhaustionBehavior `json:"exhaustionBehaviorType,omitempty"`
	Status                 *CapabilityStatus            `json:"status,omitempty"`
}
