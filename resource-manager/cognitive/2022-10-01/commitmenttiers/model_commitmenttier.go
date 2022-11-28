package commitmenttiers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommitmentTier struct {
	Cost         *CommitmentCost  `json:"cost"`
	HostingModel *HostingModel    `json:"hostingModel,omitempty"`
	Kind         *string          `json:"kind,omitempty"`
	MaxCount     *int64           `json:"maxCount,omitempty"`
	PlanType     *string          `json:"planType,omitempty"`
	Quota        *CommitmentQuota `json:"quota"`
	SkuName      *string          `json:"skuName,omitempty"`
	Tier         *string          `json:"tier,omitempty"`
}
