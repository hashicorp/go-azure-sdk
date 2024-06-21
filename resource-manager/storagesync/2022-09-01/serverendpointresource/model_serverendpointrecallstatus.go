package serverendpointresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerEndpointRecallStatus struct {
	LastUpdatedTimestamp   *string                      `json:"lastUpdatedTimestamp,omitempty"`
	RecallErrors           *[]ServerEndpointRecallError `json:"recallErrors,omitempty"`
	TotalRecallErrorsCount *int64                       `json:"totalRecallErrorsCount,omitempty"`
}
