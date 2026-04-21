package syncmembers

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncMember struct {
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	Properties *SyncMemberProperties  `json:"properties,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}
