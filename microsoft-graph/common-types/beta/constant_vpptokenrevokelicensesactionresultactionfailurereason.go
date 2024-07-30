package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenRevokeLicensesActionResultActionFailureReason string

const (
	VppTokenRevokeLicensesActionResultActionFailureReason_AppleFailure                            VppTokenRevokeLicensesActionResultActionFailureReason = "appleFailure"
	VppTokenRevokeLicensesActionResultActionFailureReason_ExpiredApplePushNotificationCertificate VppTokenRevokeLicensesActionResultActionFailureReason = "expiredApplePushNotificationCertificate"
	VppTokenRevokeLicensesActionResultActionFailureReason_ExpiredVppToken                         VppTokenRevokeLicensesActionResultActionFailureReason = "expiredVppToken"
	VppTokenRevokeLicensesActionResultActionFailureReason_InternalError                           VppTokenRevokeLicensesActionResultActionFailureReason = "internalError"
	VppTokenRevokeLicensesActionResultActionFailureReason_None                                    VppTokenRevokeLicensesActionResultActionFailureReason = "none"
)

func PossibleValuesForVppTokenRevokeLicensesActionResultActionFailureReason() []string {
	return []string{
		string(VppTokenRevokeLicensesActionResultActionFailureReason_AppleFailure),
		string(VppTokenRevokeLicensesActionResultActionFailureReason_ExpiredApplePushNotificationCertificate),
		string(VppTokenRevokeLicensesActionResultActionFailureReason_ExpiredVppToken),
		string(VppTokenRevokeLicensesActionResultActionFailureReason_InternalError),
		string(VppTokenRevokeLicensesActionResultActionFailureReason_None),
	}
}

func (s *VppTokenRevokeLicensesActionResultActionFailureReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVppTokenRevokeLicensesActionResultActionFailureReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVppTokenRevokeLicensesActionResultActionFailureReason(input string) (*VppTokenRevokeLicensesActionResultActionFailureReason, error) {
	vals := map[string]VppTokenRevokeLicensesActionResultActionFailureReason{
		"applefailure": VppTokenRevokeLicensesActionResultActionFailureReason_AppleFailure,
		"expiredapplepushnotificationcertificate": VppTokenRevokeLicensesActionResultActionFailureReason_ExpiredApplePushNotificationCertificate,
		"expiredvpptoken":                         VppTokenRevokeLicensesActionResultActionFailureReason_ExpiredVppToken,
		"internalerror":                           VppTokenRevokeLicensesActionResultActionFailureReason_InternalError,
		"none":                                    VppTokenRevokeLicensesActionResultActionFailureReason_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VppTokenRevokeLicensesActionResultActionFailureReason(input)
	return &out, nil
}
