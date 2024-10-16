package checkpolicyrestrictions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FieldRestrictions struct {
	Field        *string             `json:"field,omitempty"`
	Restrictions *[]FieldRestriction `json:"restrictions,omitempty"`
}
