package standbyvirtualmachinepoolruntimeviews

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PoolStatus struct {
	Code    HealthStateCode `json:"code"`
	Message *string         `json:"message,omitempty"`
}
