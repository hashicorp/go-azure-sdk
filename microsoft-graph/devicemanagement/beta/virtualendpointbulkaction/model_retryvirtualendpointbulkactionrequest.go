package virtualendpointbulkaction

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetryVirtualEndpointBulkActionRequest struct {
	CloudPCIds *[]string `json:"cloudPcIds,omitempty"`
}
