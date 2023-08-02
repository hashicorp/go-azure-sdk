package skus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabServicesSkuRestrictions struct {
	ReasonCode *RestrictionReasonCode `json:"reasonCode,omitempty"`
	Type       *RestrictionType       `json:"type,omitempty"`
	Values     *[]string              `json:"values,omitempty"`
}
