package logger

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LoggerType string

const (
	LoggerTypeApplicationInsights LoggerType = "applicationInsights"
	LoggerTypeAzureEventHub       LoggerType = "azureEventHub"
	LoggerTypeAzureMonitor        LoggerType = "azureMonitor"
)

func PossibleValuesForLoggerType() []string {
	return []string{
		string(LoggerTypeApplicationInsights),
		string(LoggerTypeAzureEventHub),
		string(LoggerTypeAzureMonitor),
	}
}

func parseLoggerType(input string) (*LoggerType, error) {
	vals := map[string]LoggerType{
		"applicationinsights": LoggerTypeApplicationInsights,
		"azureeventhub":       LoggerTypeAzureEventHub,
		"azuremonitor":        LoggerTypeAzureMonitor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LoggerType(input)
	return &out, nil
}
