package checkpolicyrestrictions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FieldRestriction struct {
	DefaultValue *string                 `json:"defaultValue,omitempty"`
	Policy       *PolicyReference        `json:"policy,omitempty"`
	PolicyEffect *string                 `json:"policyEffect,omitempty"`
	Reason       *string                 `json:"reason,omitempty"`
	Result       *FieldRestrictionResult `json:"result,omitempty"`
	Values       *[]string               `json:"values,omitempty"`
}
