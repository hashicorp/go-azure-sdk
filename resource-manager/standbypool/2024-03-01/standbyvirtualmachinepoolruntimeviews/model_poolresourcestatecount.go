package standbyvirtualmachinepoolruntimeviews

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PoolResourceStateCount struct {
	Count int64  `json:"count"`
	State string `json:"state"`
}
