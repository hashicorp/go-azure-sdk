package logger

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
