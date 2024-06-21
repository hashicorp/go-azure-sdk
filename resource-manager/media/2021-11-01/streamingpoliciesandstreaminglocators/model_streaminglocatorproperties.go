package streamingpoliciesandstreaminglocators

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StreamingLocatorProperties struct {
	AlternativeMediaId          *string                       `json:"alternativeMediaId,omitempty"`
	AssetName                   string                        `json:"assetName"`
	ContentKeys                 *[]StreamingLocatorContentKey `json:"contentKeys,omitempty"`
	Created                     *string                       `json:"created,omitempty"`
	DefaultContentKeyPolicyName *string                       `json:"defaultContentKeyPolicyName,omitempty"`
	EndTime                     *string                       `json:"endTime,omitempty"`
	Filters                     *[]string                     `json:"filters,omitempty"`
	StartTime                   *string                       `json:"startTime,omitempty"`
	StreamingLocatorId          *string                       `json:"streamingLocatorId,omitempty"`
	StreamingPolicyName         string                        `json:"streamingPolicyName"`
}
