package replicationappliances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InMageRcmApplianceDetails struct {
	Dra                                *DraDetails                                          `json:"dra,omitempty"`
	FabricArmId                        *string                                              `json:"fabricArmId,omitempty"`
	Id                                 *string                                              `json:"id,omitempty"`
	MarsAgent                          *MarsAgentDetails                                    `json:"marsAgent,omitempty"`
	Name                               *string                                              `json:"name,omitempty"`
	ProcessServer                      *ProcessServerDetails                                `json:"processServer,omitempty"`
	PushInstaller                      *PushInstallerDetails                                `json:"pushInstaller,omitempty"`
	RcmProxy                           *RcmProxyDetails                                     `json:"rcmProxy,omitempty"`
	ReplicationAgent                   *ReplicationAgentDetails                             `json:"replicationAgent,omitempty"`
	ReprotectAgent                     *ReprotectAgentDetails                               `json:"reprotectAgent,omitempty"`
	SwitchProviderBlockingErrorDetails *[]InMageRcmFabricSwitchProviderBlockingErrorDetails `json:"switchProviderBlockingErrorDetails,omitempty"`
}
