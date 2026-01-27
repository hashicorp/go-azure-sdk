package computenodes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeNodeEndpointConfiguration struct {
	InboundEndpoints []InboundEndpoint `json:"inboundEndpoints"`
}
