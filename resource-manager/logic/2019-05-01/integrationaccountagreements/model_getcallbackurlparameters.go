package integrationaccountagreements

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetCallbackUrlParameters struct {
	KeyType  *KeyType `json:"keyType,omitempty"`
	NotAfter *string  `json:"notAfter,omitempty"`
}
