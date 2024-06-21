package assetsandassetfilters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListContainerSasInput struct {
	ExpiryTime  *string                   `json:"expiryTime,omitempty"`
	Permissions *AssetContainerPermission `json:"permissions,omitempty"`
}
