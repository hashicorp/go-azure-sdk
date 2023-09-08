package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceHealthIssueStatus string

const (
	ServiceHealthIssueStatusconfirmed                   ServiceHealthIssueStatus = "Confirmed"
	ServiceHealthIssueStatusextendedRecovery            ServiceHealthIssueStatus = "ExtendedRecovery"
	ServiceHealthIssueStatusfalsePositive               ServiceHealthIssueStatus = "FalsePositive"
	ServiceHealthIssueStatusinvestigating               ServiceHealthIssueStatus = "Investigating"
	ServiceHealthIssueStatusinvestigationSuspended      ServiceHealthIssueStatus = "InvestigationSuspended"
	ServiceHealthIssueStatusmitigated                   ServiceHealthIssueStatus = "Mitigated"
	ServiceHealthIssueStatusmitigatedExternal           ServiceHealthIssueStatus = "MitigatedExternal"
	ServiceHealthIssueStatuspostIncidentReviewPublished ServiceHealthIssueStatus = "PostIncidentReviewPublished"
	ServiceHealthIssueStatusreported                    ServiceHealthIssueStatus = "Reported"
	ServiceHealthIssueStatusresolved                    ServiceHealthIssueStatus = "Resolved"
	ServiceHealthIssueStatusresolvedExternal            ServiceHealthIssueStatus = "ResolvedExternal"
	ServiceHealthIssueStatusrestoringService            ServiceHealthIssueStatus = "RestoringService"
	ServiceHealthIssueStatusserviceDegradation          ServiceHealthIssueStatus = "ServiceDegradation"
	ServiceHealthIssueStatusserviceInterruption         ServiceHealthIssueStatus = "ServiceInterruption"
	ServiceHealthIssueStatusserviceOperational          ServiceHealthIssueStatus = "ServiceOperational"
	ServiceHealthIssueStatusserviceRestored             ServiceHealthIssueStatus = "ServiceRestored"
	ServiceHealthIssueStatusverifyingService            ServiceHealthIssueStatus = "VerifyingService"
)

func PossibleValuesForServiceHealthIssueStatus() []string {
	return []string{
		string(ServiceHealthIssueStatusconfirmed),
		string(ServiceHealthIssueStatusextendedRecovery),
		string(ServiceHealthIssueStatusfalsePositive),
		string(ServiceHealthIssueStatusinvestigating),
		string(ServiceHealthIssueStatusinvestigationSuspended),
		string(ServiceHealthIssueStatusmitigated),
		string(ServiceHealthIssueStatusmitigatedExternal),
		string(ServiceHealthIssueStatuspostIncidentReviewPublished),
		string(ServiceHealthIssueStatusreported),
		string(ServiceHealthIssueStatusresolved),
		string(ServiceHealthIssueStatusresolvedExternal),
		string(ServiceHealthIssueStatusrestoringService),
		string(ServiceHealthIssueStatusserviceDegradation),
		string(ServiceHealthIssueStatusserviceInterruption),
		string(ServiceHealthIssueStatusserviceOperational),
		string(ServiceHealthIssueStatusserviceRestored),
		string(ServiceHealthIssueStatusverifyingService),
	}
}

func parseServiceHealthIssueStatus(input string) (*ServiceHealthIssueStatus, error) {
	vals := map[string]ServiceHealthIssueStatus{
		"confirmed":                   ServiceHealthIssueStatusconfirmed,
		"extendedrecovery":            ServiceHealthIssueStatusextendedRecovery,
		"falsepositive":               ServiceHealthIssueStatusfalsePositive,
		"investigating":               ServiceHealthIssueStatusinvestigating,
		"investigationsuspended":      ServiceHealthIssueStatusinvestigationSuspended,
		"mitigated":                   ServiceHealthIssueStatusmitigated,
		"mitigatedexternal":           ServiceHealthIssueStatusmitigatedExternal,
		"postincidentreviewpublished": ServiceHealthIssueStatuspostIncidentReviewPublished,
		"reported":                    ServiceHealthIssueStatusreported,
		"resolved":                    ServiceHealthIssueStatusresolved,
		"resolvedexternal":            ServiceHealthIssueStatusresolvedExternal,
		"restoringservice":            ServiceHealthIssueStatusrestoringService,
		"servicedegradation":          ServiceHealthIssueStatusserviceDegradation,
		"serviceinterruption":         ServiceHealthIssueStatusserviceInterruption,
		"serviceoperational":          ServiceHealthIssueStatusserviceOperational,
		"servicerestored":             ServiceHealthIssueStatusserviceRestored,
		"verifyingservice":            ServiceHealthIssueStatusverifyingService,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceHealthIssueStatus(input)
	return &out, nil
}
