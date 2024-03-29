package servicefabrics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicableScheduleProperties struct {
	LabVMsShutdown *Schedule `json:"labVmsShutdown,omitempty"`
	LabVMsStartup  *Schedule `json:"labVmsStartup,omitempty"`
}
