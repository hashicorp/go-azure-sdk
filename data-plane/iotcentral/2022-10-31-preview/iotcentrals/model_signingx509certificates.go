package iotcentrals

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SigningX509Certificates struct {
	Primary   *SigningX509Certificate `json:"primary,omitempty"`
	Secondary *SigningX509Certificate `json:"secondary,omitempty"`
}
