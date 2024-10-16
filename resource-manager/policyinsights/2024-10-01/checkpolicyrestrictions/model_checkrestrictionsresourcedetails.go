package checkpolicyrestrictions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckRestrictionsResourceDetails struct {
	ApiVersion      *string     `json:"apiVersion,omitempty"`
	ResourceContent interface{} `json:"resourceContent"`
	Scope           *string     `json:"scope,omitempty"`
}
