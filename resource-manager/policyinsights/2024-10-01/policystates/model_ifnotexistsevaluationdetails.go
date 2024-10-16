package policystates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IfNotExistsEvaluationDetails struct {
	ResourceId     *string `json:"resourceId,omitempty"`
	TotalResources *int64  `json:"totalResources,omitempty"`
}
