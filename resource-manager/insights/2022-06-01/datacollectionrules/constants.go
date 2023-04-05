package datacollectionrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KnownColumnDefinitionType string

const (
	KnownColumnDefinitionTypeBoolean  KnownColumnDefinitionType = "boolean"
	KnownColumnDefinitionTypeDatetime KnownColumnDefinitionType = "datetime"
	KnownColumnDefinitionTypeDynamic  KnownColumnDefinitionType = "dynamic"
	KnownColumnDefinitionTypeInt      KnownColumnDefinitionType = "int"
	KnownColumnDefinitionTypeLong     KnownColumnDefinitionType = "long"
	KnownColumnDefinitionTypeReal     KnownColumnDefinitionType = "real"
	KnownColumnDefinitionTypeString   KnownColumnDefinitionType = "string"
)

func PossibleValuesForKnownColumnDefinitionType() []string {
	return []string{
		string(KnownColumnDefinitionTypeBoolean),
		string(KnownColumnDefinitionTypeDatetime),
		string(KnownColumnDefinitionTypeDynamic),
		string(KnownColumnDefinitionTypeInt),
		string(KnownColumnDefinitionTypeLong),
		string(KnownColumnDefinitionTypeReal),
		string(KnownColumnDefinitionTypeString),
	}
}

type KnownDataCollectionRuleProvisioningState string

const (
	KnownDataCollectionRuleProvisioningStateCanceled  KnownDataCollectionRuleProvisioningState = "Canceled"
	KnownDataCollectionRuleProvisioningStateCreating  KnownDataCollectionRuleProvisioningState = "Creating"
	KnownDataCollectionRuleProvisioningStateDeleting  KnownDataCollectionRuleProvisioningState = "Deleting"
	KnownDataCollectionRuleProvisioningStateFailed    KnownDataCollectionRuleProvisioningState = "Failed"
	KnownDataCollectionRuleProvisioningStateSucceeded KnownDataCollectionRuleProvisioningState = "Succeeded"
	KnownDataCollectionRuleProvisioningStateUpdating  KnownDataCollectionRuleProvisioningState = "Updating"
)

func PossibleValuesForKnownDataCollectionRuleProvisioningState() []string {
	return []string{
		string(KnownDataCollectionRuleProvisioningStateCanceled),
		string(KnownDataCollectionRuleProvisioningStateCreating),
		string(KnownDataCollectionRuleProvisioningStateDeleting),
		string(KnownDataCollectionRuleProvisioningStateFailed),
		string(KnownDataCollectionRuleProvisioningStateSucceeded),
		string(KnownDataCollectionRuleProvisioningStateUpdating),
	}
}

type KnownDataCollectionRuleResourceKind string

const (
	KnownDataCollectionRuleResourceKindLinux   KnownDataCollectionRuleResourceKind = "Linux"
	KnownDataCollectionRuleResourceKindWindows KnownDataCollectionRuleResourceKind = "Windows"
)

func PossibleValuesForKnownDataCollectionRuleResourceKind() []string {
	return []string{
		string(KnownDataCollectionRuleResourceKindLinux),
		string(KnownDataCollectionRuleResourceKindWindows),
	}
}

type KnownDataFlowStreams string

const (
	KnownDataFlowStreamsMicrosoftNegativeEvent           KnownDataFlowStreams = "Microsoft-Event"
	KnownDataFlowStreamsMicrosoftNegativeInsightsMetrics KnownDataFlowStreams = "Microsoft-InsightsMetrics"
	KnownDataFlowStreamsMicrosoftNegativePerf            KnownDataFlowStreams = "Microsoft-Perf"
	KnownDataFlowStreamsMicrosoftNegativeSyslog          KnownDataFlowStreams = "Microsoft-Syslog"
	KnownDataFlowStreamsMicrosoftNegativeWindowsEvent    KnownDataFlowStreams = "Microsoft-WindowsEvent"
)

func PossibleValuesForKnownDataFlowStreams() []string {
	return []string{
		string(KnownDataFlowStreamsMicrosoftNegativeEvent),
		string(KnownDataFlowStreamsMicrosoftNegativeInsightsMetrics),
		string(KnownDataFlowStreamsMicrosoftNegativePerf),
		string(KnownDataFlowStreamsMicrosoftNegativeSyslog),
		string(KnownDataFlowStreamsMicrosoftNegativeWindowsEvent),
	}
}

type KnownExtensionDataSourceStreams string

const (
	KnownExtensionDataSourceStreamsMicrosoftNegativeEvent           KnownExtensionDataSourceStreams = "Microsoft-Event"
	KnownExtensionDataSourceStreamsMicrosoftNegativeInsightsMetrics KnownExtensionDataSourceStreams = "Microsoft-InsightsMetrics"
	KnownExtensionDataSourceStreamsMicrosoftNegativePerf            KnownExtensionDataSourceStreams = "Microsoft-Perf"
	KnownExtensionDataSourceStreamsMicrosoftNegativeSyslog          KnownExtensionDataSourceStreams = "Microsoft-Syslog"
	KnownExtensionDataSourceStreamsMicrosoftNegativeWindowsEvent    KnownExtensionDataSourceStreams = "Microsoft-WindowsEvent"
)

func PossibleValuesForKnownExtensionDataSourceStreams() []string {
	return []string{
		string(KnownExtensionDataSourceStreamsMicrosoftNegativeEvent),
		string(KnownExtensionDataSourceStreamsMicrosoftNegativeInsightsMetrics),
		string(KnownExtensionDataSourceStreamsMicrosoftNegativePerf),
		string(KnownExtensionDataSourceStreamsMicrosoftNegativeSyslog),
		string(KnownExtensionDataSourceStreamsMicrosoftNegativeWindowsEvent),
	}
}

type KnownLogFileTextSettingsRecordStartTimestampFormat string

const (
	KnownLogFileTextSettingsRecordStartTimestampFormatDdMMMYyyyHHMmSsZzz               KnownLogFileTextSettingsRecordStartTimestampFormat = "dd/MMM/yyyy:HH:mm:ss zzz"
	KnownLogFileTextSettingsRecordStartTimestampFormatDdMMyyHHMmSs                     KnownLogFileTextSettingsRecordStartTimestampFormat = "ddMMyy HH:mm:ss"
	KnownLogFileTextSettingsRecordStartTimestampFormatISOEightSixZeroOne               KnownLogFileTextSettingsRecordStartTimestampFormat = "ISO 8601"
	KnownLogFileTextSettingsRecordStartTimestampFormatMDYYYYHHMMSSAMPM                 KnownLogFileTextSettingsRecordStartTimestampFormat = "M/D/YYYY HH:MM:SS AM/PM"
	KnownLogFileTextSettingsRecordStartTimestampFormatMMMDHhMmSs                       KnownLogFileTextSettingsRecordStartTimestampFormat = "MMM d hh:mm:ss"
	KnownLogFileTextSettingsRecordStartTimestampFormatMonDDYYYYHHMMSS                  KnownLogFileTextSettingsRecordStartTimestampFormat = "Mon DD, YYYY HH:MM:SS"
	KnownLogFileTextSettingsRecordStartTimestampFormatYYYYNegativeMMNegativeDDHHMMSS   KnownLogFileTextSettingsRecordStartTimestampFormat = "YYYY-MM-DD HH:MM:SS"
	KnownLogFileTextSettingsRecordStartTimestampFormatYyMMddHHMmSs                     KnownLogFileTextSettingsRecordStartTimestampFormat = "yyMMdd HH:mm:ss"
	KnownLogFileTextSettingsRecordStartTimestampFormatYyyyNegativeMMNegativeddTHHMmSsK KnownLogFileTextSettingsRecordStartTimestampFormat = "yyyy-MM-ddTHH:mm:ssK"
)

func PossibleValuesForKnownLogFileTextSettingsRecordStartTimestampFormat() []string {
	return []string{
		string(KnownLogFileTextSettingsRecordStartTimestampFormatDdMMMYyyyHHMmSsZzz),
		string(KnownLogFileTextSettingsRecordStartTimestampFormatDdMMyyHHMmSs),
		string(KnownLogFileTextSettingsRecordStartTimestampFormatISOEightSixZeroOne),
		string(KnownLogFileTextSettingsRecordStartTimestampFormatMDYYYYHHMMSSAMPM),
		string(KnownLogFileTextSettingsRecordStartTimestampFormatMMMDHhMmSs),
		string(KnownLogFileTextSettingsRecordStartTimestampFormatMonDDYYYYHHMMSS),
		string(KnownLogFileTextSettingsRecordStartTimestampFormatYYYYNegativeMMNegativeDDHHMMSS),
		string(KnownLogFileTextSettingsRecordStartTimestampFormatYyMMddHHMmSs),
		string(KnownLogFileTextSettingsRecordStartTimestampFormatYyyyNegativeMMNegativeddTHHMmSsK),
	}
}

type KnownLogFilesDataSourceFormat string

const (
	KnownLogFilesDataSourceFormatText KnownLogFilesDataSourceFormat = "text"
)

func PossibleValuesForKnownLogFilesDataSourceFormat() []string {
	return []string{
		string(KnownLogFilesDataSourceFormatText),
	}
}

type KnownPerfCounterDataSourceStreams string

const (
	KnownPerfCounterDataSourceStreamsMicrosoftNegativeInsightsMetrics KnownPerfCounterDataSourceStreams = "Microsoft-InsightsMetrics"
	KnownPerfCounterDataSourceStreamsMicrosoftNegativePerf            KnownPerfCounterDataSourceStreams = "Microsoft-Perf"
)

func PossibleValuesForKnownPerfCounterDataSourceStreams() []string {
	return []string{
		string(KnownPerfCounterDataSourceStreamsMicrosoftNegativeInsightsMetrics),
		string(KnownPerfCounterDataSourceStreamsMicrosoftNegativePerf),
	}
}

type KnownPrometheusForwarderDataSourceStreams string

const (
	KnownPrometheusForwarderDataSourceStreamsMicrosoftNegativePrometheusMetrics KnownPrometheusForwarderDataSourceStreams = "Microsoft-PrometheusMetrics"
)

func PossibleValuesForKnownPrometheusForwarderDataSourceStreams() []string {
	return []string{
		string(KnownPrometheusForwarderDataSourceStreamsMicrosoftNegativePrometheusMetrics),
	}
}

type KnownSyslogDataSourceFacilityNames string

const (
	KnownSyslogDataSourceFacilityNamesAny        KnownSyslogDataSourceFacilityNames = "*"
	KnownSyslogDataSourceFacilityNamesAuth       KnownSyslogDataSourceFacilityNames = "auth"
	KnownSyslogDataSourceFacilityNamesAuthpriv   KnownSyslogDataSourceFacilityNames = "authpriv"
	KnownSyslogDataSourceFacilityNamesCron       KnownSyslogDataSourceFacilityNames = "cron"
	KnownSyslogDataSourceFacilityNamesDaemon     KnownSyslogDataSourceFacilityNames = "daemon"
	KnownSyslogDataSourceFacilityNamesKern       KnownSyslogDataSourceFacilityNames = "kern"
	KnownSyslogDataSourceFacilityNamesLocalFive  KnownSyslogDataSourceFacilityNames = "local5"
	KnownSyslogDataSourceFacilityNamesLocalFour  KnownSyslogDataSourceFacilityNames = "local4"
	KnownSyslogDataSourceFacilityNamesLocalOne   KnownSyslogDataSourceFacilityNames = "local1"
	KnownSyslogDataSourceFacilityNamesLocalSeven KnownSyslogDataSourceFacilityNames = "local7"
	KnownSyslogDataSourceFacilityNamesLocalSix   KnownSyslogDataSourceFacilityNames = "local6"
	KnownSyslogDataSourceFacilityNamesLocalThree KnownSyslogDataSourceFacilityNames = "local3"
	KnownSyslogDataSourceFacilityNamesLocalTwo   KnownSyslogDataSourceFacilityNames = "local2"
	KnownSyslogDataSourceFacilityNamesLocalZero  KnownSyslogDataSourceFacilityNames = "local0"
	KnownSyslogDataSourceFacilityNamesLpr        KnownSyslogDataSourceFacilityNames = "lpr"
	KnownSyslogDataSourceFacilityNamesMail       KnownSyslogDataSourceFacilityNames = "mail"
	KnownSyslogDataSourceFacilityNamesMark       KnownSyslogDataSourceFacilityNames = "mark"
	KnownSyslogDataSourceFacilityNamesNews       KnownSyslogDataSourceFacilityNames = "news"
	KnownSyslogDataSourceFacilityNamesSyslog     KnownSyslogDataSourceFacilityNames = "syslog"
	KnownSyslogDataSourceFacilityNamesUser       KnownSyslogDataSourceFacilityNames = "user"
	KnownSyslogDataSourceFacilityNamesUucp       KnownSyslogDataSourceFacilityNames = "uucp"
)

func PossibleValuesForKnownSyslogDataSourceFacilityNames() []string {
	return []string{
		string(KnownSyslogDataSourceFacilityNamesAny),
		string(KnownSyslogDataSourceFacilityNamesAuth),
		string(KnownSyslogDataSourceFacilityNamesAuthpriv),
		string(KnownSyslogDataSourceFacilityNamesCron),
		string(KnownSyslogDataSourceFacilityNamesDaemon),
		string(KnownSyslogDataSourceFacilityNamesKern),
		string(KnownSyslogDataSourceFacilityNamesLocalFive),
		string(KnownSyslogDataSourceFacilityNamesLocalFour),
		string(KnownSyslogDataSourceFacilityNamesLocalOne),
		string(KnownSyslogDataSourceFacilityNamesLocalSeven),
		string(KnownSyslogDataSourceFacilityNamesLocalSix),
		string(KnownSyslogDataSourceFacilityNamesLocalThree),
		string(KnownSyslogDataSourceFacilityNamesLocalTwo),
		string(KnownSyslogDataSourceFacilityNamesLocalZero),
		string(KnownSyslogDataSourceFacilityNamesLpr),
		string(KnownSyslogDataSourceFacilityNamesMail),
		string(KnownSyslogDataSourceFacilityNamesMark),
		string(KnownSyslogDataSourceFacilityNamesNews),
		string(KnownSyslogDataSourceFacilityNamesSyslog),
		string(KnownSyslogDataSourceFacilityNamesUser),
		string(KnownSyslogDataSourceFacilityNamesUucp),
	}
}

type KnownSyslogDataSourceLogLevels string

const (
	KnownSyslogDataSourceLogLevelsAlert     KnownSyslogDataSourceLogLevels = "Alert"
	KnownSyslogDataSourceLogLevelsAny       KnownSyslogDataSourceLogLevels = "*"
	KnownSyslogDataSourceLogLevelsCritical  KnownSyslogDataSourceLogLevels = "Critical"
	KnownSyslogDataSourceLogLevelsDebug     KnownSyslogDataSourceLogLevels = "Debug"
	KnownSyslogDataSourceLogLevelsEmergency KnownSyslogDataSourceLogLevels = "Emergency"
	KnownSyslogDataSourceLogLevelsError     KnownSyslogDataSourceLogLevels = "Error"
	KnownSyslogDataSourceLogLevelsInfo      KnownSyslogDataSourceLogLevels = "Info"
	KnownSyslogDataSourceLogLevelsNotice    KnownSyslogDataSourceLogLevels = "Notice"
	KnownSyslogDataSourceLogLevelsWarning   KnownSyslogDataSourceLogLevels = "Warning"
)

func PossibleValuesForKnownSyslogDataSourceLogLevels() []string {
	return []string{
		string(KnownSyslogDataSourceLogLevelsAlert),
		string(KnownSyslogDataSourceLogLevelsAny),
		string(KnownSyslogDataSourceLogLevelsCritical),
		string(KnownSyslogDataSourceLogLevelsDebug),
		string(KnownSyslogDataSourceLogLevelsEmergency),
		string(KnownSyslogDataSourceLogLevelsError),
		string(KnownSyslogDataSourceLogLevelsInfo),
		string(KnownSyslogDataSourceLogLevelsNotice),
		string(KnownSyslogDataSourceLogLevelsWarning),
	}
}

type KnownSyslogDataSourceStreams string

const (
	KnownSyslogDataSourceStreamsMicrosoftNegativeSyslog KnownSyslogDataSourceStreams = "Microsoft-Syslog"
)

func PossibleValuesForKnownSyslogDataSourceStreams() []string {
	return []string{
		string(KnownSyslogDataSourceStreamsMicrosoftNegativeSyslog),
	}
}

type KnownWindowsEventLogDataSourceStreams string

const (
	KnownWindowsEventLogDataSourceStreamsMicrosoftNegativeEvent        KnownWindowsEventLogDataSourceStreams = "Microsoft-Event"
	KnownWindowsEventLogDataSourceStreamsMicrosoftNegativeWindowsEvent KnownWindowsEventLogDataSourceStreams = "Microsoft-WindowsEvent"
)

func PossibleValuesForKnownWindowsEventLogDataSourceStreams() []string {
	return []string{
		string(KnownWindowsEventLogDataSourceStreamsMicrosoftNegativeEvent),
		string(KnownWindowsEventLogDataSourceStreamsMicrosoftNegativeWindowsEvent),
	}
}
