package changedatacaptureconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeDataCaptureTopologyBindingStatus struct {
	ActiveClusterResourceId            string                                `json:"activeClusterResourceId"`
	Generation                         int64                                 `json:"generation"`
	ReplacementConfigurationResourceId *string                               `json:"replacementConfigurationResourceId,omitempty"`
	State                              ChangeDataCaptureTopologyBindingState `json:"state"`
}
