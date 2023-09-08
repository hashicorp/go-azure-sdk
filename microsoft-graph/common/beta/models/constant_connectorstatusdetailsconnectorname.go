package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorStatusDetailsConnectorName string

const (
	ConnectorStatusDetailsConnectorNameappleDepExpirationDateTime                        ConnectorStatusDetailsConnectorName = "AppleDepExpirationDateTime"
	ConnectorStatusDetailsConnectorNameappleDepLastSyncDateTime                          ConnectorStatusDetailsConnectorName = "AppleDepLastSyncDateTime"
	ConnectorStatusDetailsConnectorNameapplePushNotificationServiceExpirationDateTime    ConnectorStatusDetailsConnectorName = "ApplePushNotificationServiceExpirationDateTime"
	ConnectorStatusDetailsConnectorNamechromebookLastDirectorySyncDateTime               ConnectorStatusDetailsConnectorName = "ChromebookLastDirectorySyncDateTime"
	ConnectorStatusDetailsConnectorNamefutureValue                                       ConnectorStatusDetailsConnectorName = "FutureValue"
	ConnectorStatusDetailsConnectorNamegooglePlayAppLastSyncDateTime                     ConnectorStatusDetailsConnectorName = "GooglePlayAppLastSyncDateTime"
	ConnectorStatusDetailsConnectorNamegooglePlayConnectorLastModifiedDateTime           ConnectorStatusDetailsConnectorName = "GooglePlayConnectorLastModifiedDateTime"
	ConnectorStatusDetailsConnectorNamejamfLastSyncDateTime                              ConnectorStatusDetailsConnectorName = "JamfLastSyncDateTime"
	ConnectorStatusDetailsConnectorNamemobileThreatDefenceConnectorLastHeartbeatDateTime ConnectorStatusDetailsConnectorName = "MobileThreatDefenceConnectorLastHeartbeatDateTime"
	ConnectorStatusDetailsConnectorNamendesConnectorLastConnectionDateTime               ConnectorStatusDetailsConnectorName = "NdesConnectorLastConnectionDateTime"
	ConnectorStatusDetailsConnectorNameonPremConnectorLastSyncDateTime                   ConnectorStatusDetailsConnectorName = "OnPremConnectorLastSyncDateTime"
	ConnectorStatusDetailsConnectorNamevppTokenExpirationDateTime                        ConnectorStatusDetailsConnectorName = "VppTokenExpirationDateTime"
	ConnectorStatusDetailsConnectorNamevppTokenLastSyncDateTime                          ConnectorStatusDetailsConnectorName = "VppTokenLastSyncDateTime"
	ConnectorStatusDetailsConnectorNamewindowsAutopilotLastSyncDateTime                  ConnectorStatusDetailsConnectorName = "WindowsAutopilotLastSyncDateTime"
	ConnectorStatusDetailsConnectorNamewindowsDefenderATPConnectorLastHeartbeatDateTime  ConnectorStatusDetailsConnectorName = "WindowsDefenderATPConnectorLastHeartbeatDateTime"
	ConnectorStatusDetailsConnectorNamewindowsStoreForBusinessLastSyncDateTime           ConnectorStatusDetailsConnectorName = "WindowsStoreForBusinessLastSyncDateTime"
)

func PossibleValuesForConnectorStatusDetailsConnectorName() []string {
	return []string{
		string(ConnectorStatusDetailsConnectorNameappleDepExpirationDateTime),
		string(ConnectorStatusDetailsConnectorNameappleDepLastSyncDateTime),
		string(ConnectorStatusDetailsConnectorNameapplePushNotificationServiceExpirationDateTime),
		string(ConnectorStatusDetailsConnectorNamechromebookLastDirectorySyncDateTime),
		string(ConnectorStatusDetailsConnectorNamefutureValue),
		string(ConnectorStatusDetailsConnectorNamegooglePlayAppLastSyncDateTime),
		string(ConnectorStatusDetailsConnectorNamegooglePlayConnectorLastModifiedDateTime),
		string(ConnectorStatusDetailsConnectorNamejamfLastSyncDateTime),
		string(ConnectorStatusDetailsConnectorNamemobileThreatDefenceConnectorLastHeartbeatDateTime),
		string(ConnectorStatusDetailsConnectorNamendesConnectorLastConnectionDateTime),
		string(ConnectorStatusDetailsConnectorNameonPremConnectorLastSyncDateTime),
		string(ConnectorStatusDetailsConnectorNamevppTokenExpirationDateTime),
		string(ConnectorStatusDetailsConnectorNamevppTokenLastSyncDateTime),
		string(ConnectorStatusDetailsConnectorNamewindowsAutopilotLastSyncDateTime),
		string(ConnectorStatusDetailsConnectorNamewindowsDefenderATPConnectorLastHeartbeatDateTime),
		string(ConnectorStatusDetailsConnectorNamewindowsStoreForBusinessLastSyncDateTime),
	}
}

func parseConnectorStatusDetailsConnectorName(input string) (*ConnectorStatusDetailsConnectorName, error) {
	vals := map[string]ConnectorStatusDetailsConnectorName{
		"appledepexpirationdatetime":                        ConnectorStatusDetailsConnectorNameappleDepExpirationDateTime,
		"appledeplastsyncdatetime":                          ConnectorStatusDetailsConnectorNameappleDepLastSyncDateTime,
		"applepushnotificationserviceexpirationdatetime":    ConnectorStatusDetailsConnectorNameapplePushNotificationServiceExpirationDateTime,
		"chromebooklastdirectorysyncdatetime":               ConnectorStatusDetailsConnectorNamechromebookLastDirectorySyncDateTime,
		"futurevalue":                                       ConnectorStatusDetailsConnectorNamefutureValue,
		"googleplayapplastsyncdatetime":                     ConnectorStatusDetailsConnectorNamegooglePlayAppLastSyncDateTime,
		"googleplayconnectorlastmodifieddatetime":           ConnectorStatusDetailsConnectorNamegooglePlayConnectorLastModifiedDateTime,
		"jamflastsyncdatetime":                              ConnectorStatusDetailsConnectorNamejamfLastSyncDateTime,
		"mobilethreatdefenceconnectorlastheartbeatdatetime": ConnectorStatusDetailsConnectorNamemobileThreatDefenceConnectorLastHeartbeatDateTime,
		"ndesconnectorlastconnectiondatetime":               ConnectorStatusDetailsConnectorNamendesConnectorLastConnectionDateTime,
		"onpremconnectorlastsyncdatetime":                   ConnectorStatusDetailsConnectorNameonPremConnectorLastSyncDateTime,
		"vpptokenexpirationdatetime":                        ConnectorStatusDetailsConnectorNamevppTokenExpirationDateTime,
		"vpptokenlastsyncdatetime":                          ConnectorStatusDetailsConnectorNamevppTokenLastSyncDateTime,
		"windowsautopilotlastsyncdatetime":                  ConnectorStatusDetailsConnectorNamewindowsAutopilotLastSyncDateTime,
		"windowsdefenderatpconnectorlastheartbeatdatetime":  ConnectorStatusDetailsConnectorNamewindowsDefenderATPConnectorLastHeartbeatDateTime,
		"windowsstoreforbusinesslastsyncdatetime":           ConnectorStatusDetailsConnectorNamewindowsStoreForBusinessLastSyncDateTime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectorStatusDetailsConnectorName(input)
	return &out, nil
}
