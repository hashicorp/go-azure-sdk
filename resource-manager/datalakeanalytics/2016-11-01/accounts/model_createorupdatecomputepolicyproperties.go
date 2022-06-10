package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateComputePolicyProperties struct {
	MaxDegreeOfParallelismPerJob *int64        `json:"maxDegreeOfParallelismPerJob,omitempty"`
	MinPriorityPerJob            *int64        `json:"minPriorityPerJob,omitempty"`
	ObjectId                     string        `json:"objectId"`
	ObjectType                   AADObjectType `json:"objectType"`
}
