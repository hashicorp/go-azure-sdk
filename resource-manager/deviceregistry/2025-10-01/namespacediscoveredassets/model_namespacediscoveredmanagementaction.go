package namespacediscoveredassets

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceDiscoveredManagementAction struct {
	ActionConfiguration *string                                  `json:"actionConfiguration,omitempty"`
	ActionType          *NamespaceDiscoveredManagementActionType `json:"actionType,omitempty"`
	LastUpdatedOn       *string                                  `json:"lastUpdatedOn,omitempty"`
	Name                string                                   `json:"name"`
	TargetUri           string                                   `json:"targetUri"`
	TimeoutInSeconds    *int64                                   `json:"timeoutInSeconds,omitempty"`
	Topic               *string                                  `json:"topic,omitempty"`
	TypeRef             *string                                  `json:"typeRef,omitempty"`
}

func (o *NamespaceDiscoveredManagementAction) GetLastUpdatedOnAsTime() (*time.Time, error) {
	if o.LastUpdatedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdatedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *NamespaceDiscoveredManagementAction) SetLastUpdatedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdatedOn = &formatted
}
