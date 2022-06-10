package namespacesprivatelinkresources

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkResourcesListResult struct {
	NextLink *string                `json:"nextLink,omitempty"`
	Value    *[]PrivateLinkResource `json:"value,omitempty"`
}
