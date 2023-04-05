package tenantactivitylogs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventLevel string

const (
	EventLevelCritical      EventLevel = "Critical"
	EventLevelError         EventLevel = "Error"
	EventLevelInformational EventLevel = "Informational"
	EventLevelVerbose       EventLevel = "Verbose"
	EventLevelWarning       EventLevel = "Warning"
)

func PossibleValuesForEventLevel() []string {
	return []string{
		string(EventLevelCritical),
		string(EventLevelError),
		string(EventLevelInformational),
		string(EventLevelVerbose),
		string(EventLevelWarning),
	}
}
