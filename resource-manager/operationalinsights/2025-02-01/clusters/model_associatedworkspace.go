package clusters

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssociatedWorkspace struct {
	AssociateDate *string `json:"associateDate,omitempty"`
	ResourceId    *string `json:"resourceId,omitempty"`
	WorkspaceId   *string `json:"workspaceId,omitempty"`
	WorkspaceName *string `json:"workspaceName,omitempty"`
}

func (o *AssociatedWorkspace) GetAssociateDateAsTime() (*time.Time, error) {
	if o.AssociateDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.AssociateDate, "2006-01-02T15:04:05Z07:00")
}

func (o *AssociatedWorkspace) SetAssociateDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.AssociateDate = &formatted
}
