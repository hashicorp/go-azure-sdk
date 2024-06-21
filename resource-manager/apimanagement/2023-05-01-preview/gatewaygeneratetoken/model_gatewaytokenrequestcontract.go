package gatewaygeneratetoken

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayTokenRequestContract struct {
	Expiry  string  `json:"expiry"`
	KeyType KeyType `json:"keyType"`
}
