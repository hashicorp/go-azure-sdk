package alerts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertSeverity string

const (
	AlertSeverityCritical      AlertSeverity = "Critical"
	AlertSeverityInformational AlertSeverity = "Informational"
	AlertSeverityWarning       AlertSeverity = "Warning"
)

func PossibleValuesForAlertSeverity() []string {
	return []string{
		string(AlertSeverityCritical),
		string(AlertSeverityInformational),
		string(AlertSeverityWarning),
	}
}
