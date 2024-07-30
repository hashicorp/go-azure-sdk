package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorStatusDetailsConnectorName string

const (
	ConnectorStatusDetailsConnectorName_AppleDepExpirationDateTime                        ConnectorStatusDetailsConnectorName = "appleDepExpirationDateTime"
	ConnectorStatusDetailsConnectorName_AppleDepLastSyncDateTime                          ConnectorStatusDetailsConnectorName = "appleDepLastSyncDateTime"
	ConnectorStatusDetailsConnectorName_ApplePushNotificationServiceExpirationDateTime    ConnectorStatusDetailsConnectorName = "applePushNotificationServiceExpirationDateTime"
	ConnectorStatusDetailsConnectorName_ChromebookLastDirectorySyncDateTime               ConnectorStatusDetailsConnectorName = "chromebookLastDirectorySyncDateTime"
	ConnectorStatusDetailsConnectorName_FutureValue                                       ConnectorStatusDetailsConnectorName = "futureValue"
	ConnectorStatusDetailsConnectorName_GooglePlayAppLastSyncDateTime                     ConnectorStatusDetailsConnectorName = "googlePlayAppLastSyncDateTime"
	ConnectorStatusDetailsConnectorName_GooglePlayConnectorLastModifiedDateTime           ConnectorStatusDetailsConnectorName = "googlePlayConnectorLastModifiedDateTime"
	ConnectorStatusDetailsConnectorName_JamfLastSyncDateTime                              ConnectorStatusDetailsConnectorName = "jamfLastSyncDateTime"
	ConnectorStatusDetailsConnectorName_MobileThreatDefenceConnectorLastHeartbeatDateTime ConnectorStatusDetailsConnectorName = "mobileThreatDefenceConnectorLastHeartbeatDateTime"
	ConnectorStatusDetailsConnectorName_NdesConnectorLastConnectionDateTime               ConnectorStatusDetailsConnectorName = "ndesConnectorLastConnectionDateTime"
	ConnectorStatusDetailsConnectorName_OnPremConnectorLastSyncDateTime                   ConnectorStatusDetailsConnectorName = "onPremConnectorLastSyncDateTime"
	ConnectorStatusDetailsConnectorName_VppTokenExpirationDateTime                        ConnectorStatusDetailsConnectorName = "vppTokenExpirationDateTime"
	ConnectorStatusDetailsConnectorName_VppTokenLastSyncDateTime                          ConnectorStatusDetailsConnectorName = "vppTokenLastSyncDateTime"
	ConnectorStatusDetailsConnectorName_WindowsAutopilotLastSyncDateTime                  ConnectorStatusDetailsConnectorName = "windowsAutopilotLastSyncDateTime"
	ConnectorStatusDetailsConnectorName_WindowsDefenderATPConnectorLastHeartbeatDateTime  ConnectorStatusDetailsConnectorName = "windowsDefenderATPConnectorLastHeartbeatDateTime"
	ConnectorStatusDetailsConnectorName_WindowsStoreForBusinessLastSyncDateTime           ConnectorStatusDetailsConnectorName = "windowsStoreForBusinessLastSyncDateTime"
)

func PossibleValuesForConnectorStatusDetailsConnectorName() []string {
	return []string{
		string(ConnectorStatusDetailsConnectorName_AppleDepExpirationDateTime),
		string(ConnectorStatusDetailsConnectorName_AppleDepLastSyncDateTime),
		string(ConnectorStatusDetailsConnectorName_ApplePushNotificationServiceExpirationDateTime),
		string(ConnectorStatusDetailsConnectorName_ChromebookLastDirectorySyncDateTime),
		string(ConnectorStatusDetailsConnectorName_FutureValue),
		string(ConnectorStatusDetailsConnectorName_GooglePlayAppLastSyncDateTime),
		string(ConnectorStatusDetailsConnectorName_GooglePlayConnectorLastModifiedDateTime),
		string(ConnectorStatusDetailsConnectorName_JamfLastSyncDateTime),
		string(ConnectorStatusDetailsConnectorName_MobileThreatDefenceConnectorLastHeartbeatDateTime),
		string(ConnectorStatusDetailsConnectorName_NdesConnectorLastConnectionDateTime),
		string(ConnectorStatusDetailsConnectorName_OnPremConnectorLastSyncDateTime),
		string(ConnectorStatusDetailsConnectorName_VppTokenExpirationDateTime),
		string(ConnectorStatusDetailsConnectorName_VppTokenLastSyncDateTime),
		string(ConnectorStatusDetailsConnectorName_WindowsAutopilotLastSyncDateTime),
		string(ConnectorStatusDetailsConnectorName_WindowsDefenderATPConnectorLastHeartbeatDateTime),
		string(ConnectorStatusDetailsConnectorName_WindowsStoreForBusinessLastSyncDateTime),
	}
}

func (s *ConnectorStatusDetailsConnectorName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectorStatusDetailsConnectorName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectorStatusDetailsConnectorName(input string) (*ConnectorStatusDetailsConnectorName, error) {
	vals := map[string]ConnectorStatusDetailsConnectorName{
		"appledepexpirationdatetime":                        ConnectorStatusDetailsConnectorName_AppleDepExpirationDateTime,
		"appledeplastsyncdatetime":                          ConnectorStatusDetailsConnectorName_AppleDepLastSyncDateTime,
		"applepushnotificationserviceexpirationdatetime":    ConnectorStatusDetailsConnectorName_ApplePushNotificationServiceExpirationDateTime,
		"chromebooklastdirectorysyncdatetime":               ConnectorStatusDetailsConnectorName_ChromebookLastDirectorySyncDateTime,
		"futurevalue":                                       ConnectorStatusDetailsConnectorName_FutureValue,
		"googleplayapplastsyncdatetime":                     ConnectorStatusDetailsConnectorName_GooglePlayAppLastSyncDateTime,
		"googleplayconnectorlastmodifieddatetime":           ConnectorStatusDetailsConnectorName_GooglePlayConnectorLastModifiedDateTime,
		"jamflastsyncdatetime":                              ConnectorStatusDetailsConnectorName_JamfLastSyncDateTime,
		"mobilethreatdefenceconnectorlastheartbeatdatetime": ConnectorStatusDetailsConnectorName_MobileThreatDefenceConnectorLastHeartbeatDateTime,
		"ndesconnectorlastconnectiondatetime":               ConnectorStatusDetailsConnectorName_NdesConnectorLastConnectionDateTime,
		"onpremconnectorlastsyncdatetime":                   ConnectorStatusDetailsConnectorName_OnPremConnectorLastSyncDateTime,
		"vpptokenexpirationdatetime":                        ConnectorStatusDetailsConnectorName_VppTokenExpirationDateTime,
		"vpptokenlastsyncdatetime":                          ConnectorStatusDetailsConnectorName_VppTokenLastSyncDateTime,
		"windowsautopilotlastsyncdatetime":                  ConnectorStatusDetailsConnectorName_WindowsAutopilotLastSyncDateTime,
		"windowsdefenderatpconnectorlastheartbeatdatetime":  ConnectorStatusDetailsConnectorName_WindowsDefenderATPConnectorLastHeartbeatDateTime,
		"windowsstoreforbusinesslastsyncdatetime":           ConnectorStatusDetailsConnectorName_WindowsStoreForBusinessLastSyncDateTime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectorStatusDetailsConnectorName(input)
	return &out, nil
}
