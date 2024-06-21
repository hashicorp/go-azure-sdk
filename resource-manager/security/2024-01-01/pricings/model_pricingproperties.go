package pricings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PricingProperties struct {
	Deprecated              *bool                    `json:"deprecated,omitempty"`
	EnablementTime          *string                  `json:"enablementTime,omitempty"`
	Enforce                 *Enforce                 `json:"enforce,omitempty"`
	Extensions              *[]Extension             `json:"extensions,omitempty"`
	FreeTrialRemainingTime  *string                  `json:"freeTrialRemainingTime,omitempty"`
	Inherited               *Inherited               `json:"inherited,omitempty"`
	InheritedFrom           *string                  `json:"inheritedFrom,omitempty"`
	PricingTier             PricingTier              `json:"pricingTier"`
	ReplacedBy              *[]string                `json:"replacedBy,omitempty"`
	ResourcesCoverageStatus *ResourcesCoverageStatus `json:"resourcesCoverageStatus,omitempty"`
	SubPlan                 *string                  `json:"subPlan,omitempty"`
}
