package manageddatabasesecurityevents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEventProperties struct {
	ApplicationName                               *string                                        `json:"applicationName,omitempty"`
	ClientIP                                      *string                                        `json:"clientIp,omitempty"`
	Database                                      *string                                        `json:"database,omitempty"`
	EventTime                                     *string                                        `json:"eventTime,omitempty"`
	PrincipalName                                 *string                                        `json:"principalName,omitempty"`
	SecurityEventSqlInjectionAdditionalProperties *SecurityEventSqlInjectionAdditionalProperties `json:"securityEventSqlInjectionAdditionalProperties,omitempty"`
	SecurityEventType                             *SecurityEventType                             `json:"securityEventType,omitempty"`
	Server                                        *string                                        `json:"server,omitempty"`
	Subscription                                  *string                                        `json:"subscription,omitempty"`
}
