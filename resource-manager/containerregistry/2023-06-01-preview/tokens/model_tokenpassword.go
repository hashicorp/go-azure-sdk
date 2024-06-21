package tokens

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TokenPassword struct {
	CreationTime *string            `json:"creationTime,omitempty"`
	Expiry       *string            `json:"expiry,omitempty"`
	Name         *TokenPasswordName `json:"name,omitempty"`
	Value        *string            `json:"value,omitempty"`
}
