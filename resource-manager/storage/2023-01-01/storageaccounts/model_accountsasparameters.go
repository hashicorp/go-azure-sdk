package storageaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountSasParameters struct {
	KeyToSign           *string             `json:"keyToSign,omitempty"`
	SignedExpiry        string              `json:"signedExpiry"`
	SignedIP            *string             `json:"signedIp,omitempty"`
	SignedPermission    Permissions         `json:"signedPermission"`
	SignedProtocol      *HTTPProtocol       `json:"signedProtocol,omitempty"`
	SignedResourceTypes SignedResourceTypes `json:"signedResourceTypes"`
	SignedServices      Services            `json:"signedServices"`
	SignedStart         *string             `json:"signedStart,omitempty"`
}
