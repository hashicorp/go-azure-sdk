package akriconnectortemplate

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AkriConnectorTemplateResource struct {
	ExtendedLocation *ExtendedLocation                `json:"extendedLocation,omitempty"`
	Id               *string                          `json:"id,omitempty"`
	Name             *string                          `json:"name,omitempty"`
	Properties       *AkriConnectorTemplateProperties `json:"properties,omitempty"`
	SystemData       *systemdata.SystemData           `json:"systemData,omitempty"`
	Type             *string                          `json:"type,omitempty"`
}
