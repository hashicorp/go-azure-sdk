package machinepools

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachinePoolUpdate struct {
	Properties *MachinePoolProperties `json:"properties,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
}
