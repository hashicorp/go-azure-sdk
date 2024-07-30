package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Simulation struct {
	AttackTechnique            *SimulationAttackTechnique         `json:"attackTechnique,omitempty"`
	AttackType                 *SimulationAttackType              `json:"attackType,omitempty"`
	AutomationId               *string                            `json:"automationId,omitempty"`
	CompletionDateTime         *string                            `json:"completionDateTime,omitempty"`
	CreatedBy                  *EmailIdentity                     `json:"createdBy,omitempty"`
	CreatedDateTime            *string                            `json:"createdDateTime,omitempty"`
	Description                *string                            `json:"description,omitempty"`
	DisplayName                *string                            `json:"displayName,omitempty"`
	DurationInDays             *int64                             `json:"durationInDays,omitempty"`
	EndUserNotificationSetting *EndUserNotificationSetting        `json:"endUserNotificationSetting,omitempty"`
	ExcludedAccountTarget      *AccountTargetContent              `json:"excludedAccountTarget,omitempty"`
	Id                         *string                            `json:"id,omitempty"`
	IncludedAccountTarget      *AccountTargetContent              `json:"includedAccountTarget,omitempty"`
	IsAutomated                *bool                              `json:"isAutomated,omitempty"`
	LandingPage                *LandingPage                       `json:"landingPage,omitempty"`
	LastModifiedBy             *EmailIdentity                     `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime       *string                            `json:"lastModifiedDateTime,omitempty"`
	LaunchDateTime             *string                            `json:"launchDateTime,omitempty"`
	LoginPage                  *LoginPage                         `json:"loginPage,omitempty"`
	OAuthConsentAppDetail      *OAuthConsentAppDetail             `json:"oAuthConsentAppDetail,omitempty"`
	ODataType                  *string                            `json:"@odata.type,omitempty"`
	Payload                    *Payload                           `json:"payload,omitempty"`
	PayloadDeliveryPlatform    *SimulationPayloadDeliveryPlatform `json:"payloadDeliveryPlatform,omitempty"`
	Report                     *SimulationReport                  `json:"report,omitempty"`
	Status                     *SimulationStatus                  `json:"status,omitempty"`
	TrainingSetting            *TrainingSetting                   `json:"trainingSetting,omitempty"`
}
