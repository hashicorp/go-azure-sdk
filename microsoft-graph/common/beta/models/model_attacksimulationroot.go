package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttackSimulationRoot struct {
	EndUserNotifications  *[]EndUserNotification       `json:"endUserNotifications,omitempty"`
	Id                    *string                      `json:"id,omitempty"`
	LandingPages          *[]LandingPage               `json:"landingPages,omitempty"`
	LoginPages            *[]LoginPage                 `json:"loginPages,omitempty"`
	ODataType             *string                      `json:"@odata.type,omitempty"`
	Operations            *[]AttackSimulationOperation `json:"operations,omitempty"`
	Payloads              *[]Payload                   `json:"payloads,omitempty"`
	SimulationAutomations *[]SimulationAutomation      `json:"simulationAutomations,omitempty"`
	Simulations           *[]Simulation                `json:"simulations,omitempty"`
	Trainings             *[]Training                  `json:"trainings,omitempty"`
}
