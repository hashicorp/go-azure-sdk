package computepolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputePolicyProperties struct {
	MaxDegreeOfParallelismPerJob *int64         `json:"maxDegreeOfParallelismPerJob,omitempty"`
	MinPriorityPerJob            *int64         `json:"minPriorityPerJob,omitempty"`
	ObjectId                     *string        `json:"objectId,omitempty"`
	ObjectType                   *AADObjectType `json:"objectType,omitempty"`
}
