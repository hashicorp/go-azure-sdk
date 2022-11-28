package replicationappliances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InMageRcmApplianceDetails struct {
	Dra                                *DraDetails                                          `json:"dra"`
	FabricArmId                        *string                                              `json:"fabricArmId,omitempty"`
	Id                                 *string                                              `json:"id,omitempty"`
	MarsAgent                          *MarsAgentDetails                                    `json:"marsAgent"`
	Name                               *string                                              `json:"name,omitempty"`
	ProcessServer                      *ProcessServerDetails                                `json:"processServer"`
	PushInstaller                      *PushInstallerDetails                                `json:"pushInstaller"`
	RcmProxy                           *RcmProxyDetails                                     `json:"rcmProxy"`
	ReplicationAgent                   *ReplicationAgentDetails                             `json:"replicationAgent"`
	ReprotectAgent                     *ReprotectAgentDetails                               `json:"reprotectAgent"`
	SwitchProviderBlockingErrorDetails *[]InMageRcmFabricSwitchProviderBlockingErrorDetails `json:"switchProviderBlockingErrorDetails,omitempty"`
}
