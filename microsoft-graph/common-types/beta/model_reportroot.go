package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ReportRoot{}

type ReportRoot struct {
	// Represents a collection of sign-in activities of application credentials.
	AppCredentialSignInActivities *[]AppCredentialSignInActivity `json:"appCredentialSignInActivities,omitempty"`

	// Represents a detailed summary of an application sign-in.
	ApplicationSignInDetailedSummary *[]ApplicationSignInDetailedSummary `json:"applicationSignInDetailedSummary,omitempty"`

	// Container for navigation properties for Microsoft Entra authentication methods resources.
	AuthenticationMethods *AuthenticationMethodsRoot `json:"authenticationMethods,omitempty"`

	// Details of the usage of self-service password reset and multifactor authentication (MFA) for all registered users.
	CredentialUserRegistrationDetails *[]CredentialUserRegistrationDetails `json:"credentialUserRegistrationDetails,omitempty"`

	DailyPrintUsage *[]PrintUsage `json:"dailyPrintUsage,omitempty"`

	// Retrieve a list of daily print usage summaries, grouped by printer.
	DailyPrintUsageByPrinter *[]PrintUsageByPrinter `json:"dailyPrintUsageByPrinter,omitempty"`

	// Retrieve a list of daily print usage summaries, grouped by user.
	DailyPrintUsageByUser *[]PrintUsageByUser `json:"dailyPrintUsageByUser,omitempty"`

	DailyPrintUsageSummariesByPrinter *[]PrintUsageByPrinter `json:"dailyPrintUsageSummariesByPrinter,omitempty"`
	DailyPrintUsageSummariesByUser    *[]PrintUsageByUser    `json:"dailyPrintUsageSummariesByUser,omitempty"`

	// Retrieve a list of monthly print usage summaries, grouped by printer.
	MonthlyPrintUsageByPrinter *[]PrintUsageByPrinter `json:"monthlyPrintUsageByPrinter,omitempty"`

	// Retrieve a list of monthly print usage summaries, grouped by user.
	MonthlyPrintUsageByUser *[]PrintUsageByUser `json:"monthlyPrintUsageByUser,omitempty"`

	MonthlyPrintUsageSummariesByPrinter *[]PrintUsageByPrinter `json:"monthlyPrintUsageSummariesByPrinter,omitempty"`
	MonthlyPrintUsageSummariesByUser    *[]PrintUsageByUser    `json:"monthlyPrintUsageSummariesByUser,omitempty"`

	// Represents billing details for a Microsoft direct partner.
	Partners *Partners `json:"partners,omitempty"`

	// Provides the ability to launch a simulated phishing attack that organizations can learn from.
	Security *SecurityReportsRoot `json:"security,omitempty"`

	// Reports that relate to tenant-level authentication activities in Microsoft Entra.
	ServiceActivity *ServiceActivity `json:"serviceActivity,omitempty"`

	// Represents a collection of sign-in activities of service principals.
	ServicePrincipalSignInActivities *[]ServicePrincipalSignInActivity `json:"servicePrincipalSignInActivities,omitempty"`

	// Reports that relate to tenant-level Microsoft Entra SLA attainment.
	Sla *ServiceLevelAgreementRoot `json:"sla,omitempty"`

	// Represents the self-service password reset (SSPR) usage for a given tenant.
	UserCredentialUsageDetails *[]UserCredentialUsageDetails `json:"userCredentialUsageDetails,omitempty"`

	// Represents a collection of user activities on applications in a tenant that is configured for Microsoft Entra
	// External ID for customers.
	UserInsights *UserInsightsRoot `json:"userInsights,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ReportRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ReportRoot{}

func (s ReportRoot) MarshalJSON() ([]byte, error) {
	type wrapper ReportRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ReportRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ReportRoot: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.reportRoot"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ReportRoot: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ReportRoot{}

func (s *ReportRoot) UnmarshalJSON(bytes []byte) error {
	type alias ReportRoot
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ReportRoot: %+v", err)
	}

	s.AppCredentialSignInActivities = decoded.AppCredentialSignInActivities
	s.ApplicationSignInDetailedSummary = decoded.ApplicationSignInDetailedSummary
	s.AuthenticationMethods = decoded.AuthenticationMethods
	s.CredentialUserRegistrationDetails = decoded.CredentialUserRegistrationDetails
	s.DailyPrintUsageByPrinter = decoded.DailyPrintUsageByPrinter
	s.DailyPrintUsageByUser = decoded.DailyPrintUsageByUser
	s.DailyPrintUsageSummariesByPrinter = decoded.DailyPrintUsageSummariesByPrinter
	s.DailyPrintUsageSummariesByUser = decoded.DailyPrintUsageSummariesByUser
	s.Id = decoded.Id
	s.MonthlyPrintUsageByPrinter = decoded.MonthlyPrintUsageByPrinter
	s.MonthlyPrintUsageByUser = decoded.MonthlyPrintUsageByUser
	s.MonthlyPrintUsageSummariesByPrinter = decoded.MonthlyPrintUsageSummariesByPrinter
	s.MonthlyPrintUsageSummariesByUser = decoded.MonthlyPrintUsageSummariesByUser
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Partners = decoded.Partners
	s.Security = decoded.Security
	s.ServiceActivity = decoded.ServiceActivity
	s.ServicePrincipalSignInActivities = decoded.ServicePrincipalSignInActivities
	s.Sla = decoded.Sla
	s.UserCredentialUsageDetails = decoded.UserCredentialUsageDetails
	s.UserInsights = decoded.UserInsights

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ReportRoot into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["dailyPrintUsage"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DailyPrintUsage into list []json.RawMessage: %+v", err)
		}

		output := make([]PrintUsage, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalPrintUsageImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DailyPrintUsage' for 'ReportRoot': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DailyPrintUsage = &output
	}
	return nil
}
