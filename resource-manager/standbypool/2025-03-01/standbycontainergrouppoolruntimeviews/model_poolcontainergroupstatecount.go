package standbycontainergrouppoolruntimeviews

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PoolContainerGroupStateCount struct {
	Count int64                   `json:"count"`
	State PoolContainerGroupState `json:"state"`
}
