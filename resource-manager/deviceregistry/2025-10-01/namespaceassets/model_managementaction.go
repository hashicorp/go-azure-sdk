package namespaceassets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementAction struct {
	ActionConfiguration *string               `json:"actionConfiguration,omitempty"`
	ActionType          *ManagementActionType `json:"actionType,omitempty"`
	Name                string                `json:"name"`
	TargetUri           string                `json:"targetUri"`
	TimeoutInSeconds    *int64                `json:"timeoutInSeconds,omitempty"`
	Topic               *string               `json:"topic,omitempty"`
	TypeRef             *string               `json:"typeRef,omitempty"`
}
