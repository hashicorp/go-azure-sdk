package subscriptionfeatureregistrations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionFeatureRegistrationProperties struct {
	ApprovalType                 *SubscriptionFeatureRegistrationApprovalType `json:"approvalType,omitempty"`
	AuthorizationProfile         *AuthorizationProfile                        `json:"authorizationProfile,omitempty"`
	Description                  *string                                      `json:"description,omitempty"`
	DisplayName                  *string                                      `json:"displayName,omitempty"`
	DocumentationLink            *string                                      `json:"documentationLink,omitempty"`
	FeatureName                  *string                                      `json:"featureName,omitempty"`
	Metadata                     *map[string]string                           `json:"metadata,omitempty"`
	ProviderNamespace            *string                                      `json:"providerNamespace,omitempty"`
	RegistrationDate             *string                                      `json:"registrationDate,omitempty"`
	ReleaseDate                  *string                                      `json:"releaseDate,omitempty"`
	ShouldFeatureDisplayInPortal *bool                                        `json:"shouldFeatureDisplayInPortal,omitempty"`
	State                        *SubscriptionFeatureRegistrationState        `json:"state,omitempty"`
	SubscriptionId               *string                                      `json:"subscriptionId,omitempty"`
	TenantId                     *string                                      `json:"tenantId,omitempty"`
}

func (o *SubscriptionFeatureRegistrationProperties) GetRegistrationDateAsTime() (*time.Time, error) {
	if o.RegistrationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RegistrationDate, "2006-01-02T15:04:05Z07:00")
}

func (o *SubscriptionFeatureRegistrationProperties) SetRegistrationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RegistrationDate = &formatted
}

func (o *SubscriptionFeatureRegistrationProperties) GetReleaseDateAsTime() (*time.Time, error) {
	if o.ReleaseDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ReleaseDate, "2006-01-02T15:04:05Z07:00")
}

func (o *SubscriptionFeatureRegistrationProperties) SetReleaseDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ReleaseDate = &formatted
}
