package namespaceassets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementGroup struct {
	Actions                      *[]ManagementAction `json:"actions,omitempty"`
	DataSource                   *string             `json:"dataSource,omitempty"`
	DefaultTimeoutInSeconds      *int64              `json:"defaultTimeoutInSeconds,omitempty"`
	DefaultTopic                 *string             `json:"defaultTopic,omitempty"`
	ManagementGroupConfiguration *string             `json:"managementGroupConfiguration,omitempty"`
	Name                         string              `json:"name"`
	TypeRef                      *string             `json:"typeRef,omitempty"`
}
