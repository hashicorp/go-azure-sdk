package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessSessionControls struct {
	ApplicationEnforcedRestrictions *ApplicationEnforcedRestrictionsSessionControl `json:"applicationEnforcedRestrictions,omitempty"`
	CloudAppSecurity                *CloudAppSecuritySessionControl                `json:"cloudAppSecurity,omitempty"`
	ContinuousAccessEvaluation      *ContinuousAccessEvaluationSessionControl      `json:"continuousAccessEvaluation,omitempty"`
	DisableResilienceDefaults       *bool                                          `json:"disableResilienceDefaults,omitempty"`
	ODataType                       *string                                        `json:"@odata.type,omitempty"`
	PersistentBrowser               *PersistentBrowserSessionControl               `json:"persistentBrowser,omitempty"`
	SecureSignInSession             *SecureSignInSessionControl                    `json:"secureSignInSession,omitempty"`
	SignInFrequency                 *SignInFrequencySessionControl                 `json:"signInFrequency,omitempty"`
}
