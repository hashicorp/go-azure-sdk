package billingpermission

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckAccessResponse struct {
	AccessDecision *AccessDecision `json:"accessDecision,omitempty"`
	Action         *string         `json:"action,omitempty"`
}
