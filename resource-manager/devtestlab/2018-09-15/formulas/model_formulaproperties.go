package formulas

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FormulaProperties struct {
	Author            *string                             `json:"author,omitempty"`
	CreationDate      *string                             `json:"creationDate,omitempty"`
	Description       *string                             `json:"description,omitempty"`
	FormulaContent    *LabVirtualMachineCreationParameter `json:"formulaContent,omitempty"`
	OsType            *string                             `json:"osType,omitempty"`
	ProvisioningState *string                             `json:"provisioningState,omitempty"`
	UniqueIdentifier  *string                             `json:"uniqueIdentifier,omitempty"`
	VM                *FormulaPropertiesFromVM            `json:"vm,omitempty"`
}

func (o *FormulaProperties) GetCreationDateAsTime() (*time.Time, error) {
	if o.CreationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationDate, "2006-01-02T15:04:05Z07:00")
}

func (o *FormulaProperties) SetCreationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationDate = &formatted
}
