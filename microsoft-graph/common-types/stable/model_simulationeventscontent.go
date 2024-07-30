package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationEventsContent struct {
	CompromisedRate *float64           `json:"compromisedRate,omitempty"`
	Events          *[]SimulationEvent `json:"events,omitempty"`
	ODataType       *string            `json:"@odata.type,omitempty"`
}
