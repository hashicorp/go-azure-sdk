package servers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterServerListResult struct {
	Value *[]ClusterServer `json:"value,omitempty"`
}
