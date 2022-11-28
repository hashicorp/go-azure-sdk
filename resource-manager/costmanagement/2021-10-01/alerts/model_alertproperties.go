package alerts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertProperties struct {
	CloseTime                  *string                    `json:"closeTime,omitempty"`
	CostEntityId               *string                    `json:"costEntityId,omitempty"`
	CreationTime               *string                    `json:"creationTime,omitempty"`
	Definition                 *AlertPropertiesDefinition `json:"definition"`
	Description                *string                    `json:"description,omitempty"`
	Details                    *AlertPropertiesDetails    `json:"details"`
	ModificationTime           *string                    `json:"modificationTime,omitempty"`
	Source                     *AlertSource               `json:"source,omitempty"`
	Status                     *AlertStatus               `json:"status,omitempty"`
	StatusModificationTime     *string                    `json:"statusModificationTime,omitempty"`
	StatusModificationUserName *string                    `json:"statusModificationUserName,omitempty"`
}
