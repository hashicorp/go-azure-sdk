package models

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Model struct {
	Description *string       `json:"description,omitempty"`
	Kind        *string       `json:"kind,omitempty"`
	Model       *AccountModel `json:"model,omitempty"`
	SkuName     *string       `json:"skuName,omitempty"`
}
