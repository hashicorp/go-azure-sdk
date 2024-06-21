package connectedregistries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectedRegistryProperties struct {
	Activation        *ActivationProperties     `json:"activation,omitempty"`
	ClientTokenIds    *[]string                 `json:"clientTokenIds,omitempty"`
	ConnectionState   *ConnectionState          `json:"connectionState,omitempty"`
	LastActivityTime  *string                   `json:"lastActivityTime,omitempty"`
	Logging           *LoggingProperties        `json:"logging,omitempty"`
	LoginServer       *LoginServerProperties    `json:"loginServer,omitempty"`
	Mode              ConnectedRegistryMode     `json:"mode"`
	NotificationsList *[]string                 `json:"notificationsList,omitempty"`
	Parent            ParentProperties          `json:"parent"`
	ProvisioningState *ProvisioningState        `json:"provisioningState,omitempty"`
	StatusDetails     *[]StatusDetailProperties `json:"statusDetails,omitempty"`
	Version           *string                   `json:"version,omitempty"`
}
