package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOsVppAppRevokeLicensesActionResultActionFailureReason string

const (
	MacOsVppAppRevokeLicensesActionResultActionFailureReason_AppleFailure                            MacOsVppAppRevokeLicensesActionResultActionFailureReason = "appleFailure"
	MacOsVppAppRevokeLicensesActionResultActionFailureReason_ExpiredApplePushNotificationCertificate MacOsVppAppRevokeLicensesActionResultActionFailureReason = "expiredApplePushNotificationCertificate"
	MacOsVppAppRevokeLicensesActionResultActionFailureReason_ExpiredVppToken                         MacOsVppAppRevokeLicensesActionResultActionFailureReason = "expiredVppToken"
	MacOsVppAppRevokeLicensesActionResultActionFailureReason_InternalError                           MacOsVppAppRevokeLicensesActionResultActionFailureReason = "internalError"
	MacOsVppAppRevokeLicensesActionResultActionFailureReason_None                                    MacOsVppAppRevokeLicensesActionResultActionFailureReason = "none"
)

func PossibleValuesForMacOsVppAppRevokeLicensesActionResultActionFailureReason() []string {
	return []string{
		string(MacOsVppAppRevokeLicensesActionResultActionFailureReason_AppleFailure),
		string(MacOsVppAppRevokeLicensesActionResultActionFailureReason_ExpiredApplePushNotificationCertificate),
		string(MacOsVppAppRevokeLicensesActionResultActionFailureReason_ExpiredVppToken),
		string(MacOsVppAppRevokeLicensesActionResultActionFailureReason_InternalError),
		string(MacOsVppAppRevokeLicensesActionResultActionFailureReason_None),
	}
}

func (s *MacOsVppAppRevokeLicensesActionResultActionFailureReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOsVppAppRevokeLicensesActionResultActionFailureReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOsVppAppRevokeLicensesActionResultActionFailureReason(input string) (*MacOsVppAppRevokeLicensesActionResultActionFailureReason, error) {
	vals := map[string]MacOsVppAppRevokeLicensesActionResultActionFailureReason{
		"applefailure": MacOsVppAppRevokeLicensesActionResultActionFailureReason_AppleFailure,
		"expiredapplepushnotificationcertificate": MacOsVppAppRevokeLicensesActionResultActionFailureReason_ExpiredApplePushNotificationCertificate,
		"expiredvpptoken":                         MacOsVppAppRevokeLicensesActionResultActionFailureReason_ExpiredVppToken,
		"internalerror":                           MacOsVppAppRevokeLicensesActionResultActionFailureReason_InternalError,
		"none":                                    MacOsVppAppRevokeLicensesActionResultActionFailureReason_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOsVppAppRevokeLicensesActionResultActionFailureReason(input)
	return &out, nil
}
