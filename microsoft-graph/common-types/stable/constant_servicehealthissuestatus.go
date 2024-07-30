package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceHealthIssueStatus string

const (
	ServiceHealthIssueStatus_Confirmed                   ServiceHealthIssueStatus = "confirmed"
	ServiceHealthIssueStatus_ExtendedRecovery            ServiceHealthIssueStatus = "extendedRecovery"
	ServiceHealthIssueStatus_FalsePositive               ServiceHealthIssueStatus = "falsePositive"
	ServiceHealthIssueStatus_Investigating               ServiceHealthIssueStatus = "investigating"
	ServiceHealthIssueStatus_InvestigationSuspended      ServiceHealthIssueStatus = "investigationSuspended"
	ServiceHealthIssueStatus_Mitigated                   ServiceHealthIssueStatus = "mitigated"
	ServiceHealthIssueStatus_MitigatedExternal           ServiceHealthIssueStatus = "mitigatedExternal"
	ServiceHealthIssueStatus_PostIncidentReviewPublished ServiceHealthIssueStatus = "postIncidentReviewPublished"
	ServiceHealthIssueStatus_Reported                    ServiceHealthIssueStatus = "reported"
	ServiceHealthIssueStatus_Resolved                    ServiceHealthIssueStatus = "resolved"
	ServiceHealthIssueStatus_ResolvedExternal            ServiceHealthIssueStatus = "resolvedExternal"
	ServiceHealthIssueStatus_RestoringService            ServiceHealthIssueStatus = "restoringService"
	ServiceHealthIssueStatus_ServiceDegradation          ServiceHealthIssueStatus = "serviceDegradation"
	ServiceHealthIssueStatus_ServiceInterruption         ServiceHealthIssueStatus = "serviceInterruption"
	ServiceHealthIssueStatus_ServiceOperational          ServiceHealthIssueStatus = "serviceOperational"
	ServiceHealthIssueStatus_ServiceRestored             ServiceHealthIssueStatus = "serviceRestored"
	ServiceHealthIssueStatus_VerifyingService            ServiceHealthIssueStatus = "verifyingService"
)

func PossibleValuesForServiceHealthIssueStatus() []string {
	return []string{
		string(ServiceHealthIssueStatus_Confirmed),
		string(ServiceHealthIssueStatus_ExtendedRecovery),
		string(ServiceHealthIssueStatus_FalsePositive),
		string(ServiceHealthIssueStatus_Investigating),
		string(ServiceHealthIssueStatus_InvestigationSuspended),
		string(ServiceHealthIssueStatus_Mitigated),
		string(ServiceHealthIssueStatus_MitigatedExternal),
		string(ServiceHealthIssueStatus_PostIncidentReviewPublished),
		string(ServiceHealthIssueStatus_Reported),
		string(ServiceHealthIssueStatus_Resolved),
		string(ServiceHealthIssueStatus_ResolvedExternal),
		string(ServiceHealthIssueStatus_RestoringService),
		string(ServiceHealthIssueStatus_ServiceDegradation),
		string(ServiceHealthIssueStatus_ServiceInterruption),
		string(ServiceHealthIssueStatus_ServiceOperational),
		string(ServiceHealthIssueStatus_ServiceRestored),
		string(ServiceHealthIssueStatus_VerifyingService),
	}
}

func (s *ServiceHealthIssueStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceHealthIssueStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceHealthIssueStatus(input string) (*ServiceHealthIssueStatus, error) {
	vals := map[string]ServiceHealthIssueStatus{
		"confirmed":                   ServiceHealthIssueStatus_Confirmed,
		"extendedrecovery":            ServiceHealthIssueStatus_ExtendedRecovery,
		"falsepositive":               ServiceHealthIssueStatus_FalsePositive,
		"investigating":               ServiceHealthIssueStatus_Investigating,
		"investigationsuspended":      ServiceHealthIssueStatus_InvestigationSuspended,
		"mitigated":                   ServiceHealthIssueStatus_Mitigated,
		"mitigatedexternal":           ServiceHealthIssueStatus_MitigatedExternal,
		"postincidentreviewpublished": ServiceHealthIssueStatus_PostIncidentReviewPublished,
		"reported":                    ServiceHealthIssueStatus_Reported,
		"resolved":                    ServiceHealthIssueStatus_Resolved,
		"resolvedexternal":            ServiceHealthIssueStatus_ResolvedExternal,
		"restoringservice":            ServiceHealthIssueStatus_RestoringService,
		"servicedegradation":          ServiceHealthIssueStatus_ServiceDegradation,
		"serviceinterruption":         ServiceHealthIssueStatus_ServiceInterruption,
		"serviceoperational":          ServiceHealthIssueStatus_ServiceOperational,
		"servicerestored":             ServiceHealthIssueStatus_ServiceRestored,
		"verifyingservice":            ServiceHealthIssueStatus_VerifyingService,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceHealthIssueStatus(input)
	return &out, nil
}
