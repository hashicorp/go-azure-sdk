package assetsandassetfilters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssetStreamingLocator struct {
	AssetName                   *string `json:"assetName,omitempty"`
	Created                     *string `json:"created,omitempty"`
	DefaultContentKeyPolicyName *string `json:"defaultContentKeyPolicyName,omitempty"`
	EndTime                     *string `json:"endTime,omitempty"`
	Name                        *string `json:"name,omitempty"`
	StartTime                   *string `json:"startTime,omitempty"`
	StreamingLocatorId          *string `json:"streamingLocatorId,omitempty"`
	StreamingPolicyName         *string `json:"streamingPolicyName,omitempty"`
}
