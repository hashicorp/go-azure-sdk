package dedicatedhsms

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkInterface struct {
	PrivateIPAddress *string `json:"privateIpAddress,omitempty"`
	ResourceId       *string `json:"resourceId,omitempty"`
}
