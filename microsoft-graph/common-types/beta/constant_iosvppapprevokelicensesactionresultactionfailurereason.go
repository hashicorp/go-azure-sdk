package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppAppRevokeLicensesActionResultActionFailureReason string

const (
	IosVppAppRevokeLicensesActionResultActionFailureReason_AppleFailure                            IosVppAppRevokeLicensesActionResultActionFailureReason = "appleFailure"
	IosVppAppRevokeLicensesActionResultActionFailureReason_ExpiredApplePushNotificationCertificate IosVppAppRevokeLicensesActionResultActionFailureReason = "expiredApplePushNotificationCertificate"
	IosVppAppRevokeLicensesActionResultActionFailureReason_ExpiredVppToken                         IosVppAppRevokeLicensesActionResultActionFailureReason = "expiredVppToken"
	IosVppAppRevokeLicensesActionResultActionFailureReason_InternalError                           IosVppAppRevokeLicensesActionResultActionFailureReason = "internalError"
	IosVppAppRevokeLicensesActionResultActionFailureReason_None                                    IosVppAppRevokeLicensesActionResultActionFailureReason = "none"
)

func PossibleValuesForIosVppAppRevokeLicensesActionResultActionFailureReason() []string {
	return []string{
		string(IosVppAppRevokeLicensesActionResultActionFailureReason_AppleFailure),
		string(IosVppAppRevokeLicensesActionResultActionFailureReason_ExpiredApplePushNotificationCertificate),
		string(IosVppAppRevokeLicensesActionResultActionFailureReason_ExpiredVppToken),
		string(IosVppAppRevokeLicensesActionResultActionFailureReason_InternalError),
		string(IosVppAppRevokeLicensesActionResultActionFailureReason_None),
	}
}

func (s *IosVppAppRevokeLicensesActionResultActionFailureReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosVppAppRevokeLicensesActionResultActionFailureReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosVppAppRevokeLicensesActionResultActionFailureReason(input string) (*IosVppAppRevokeLicensesActionResultActionFailureReason, error) {
	vals := map[string]IosVppAppRevokeLicensesActionResultActionFailureReason{
		"applefailure": IosVppAppRevokeLicensesActionResultActionFailureReason_AppleFailure,
		"expiredapplepushnotificationcertificate": IosVppAppRevokeLicensesActionResultActionFailureReason_ExpiredApplePushNotificationCertificate,
		"expiredvpptoken":                         IosVppAppRevokeLicensesActionResultActionFailureReason_ExpiredVppToken,
		"internalerror":                           IosVppAppRevokeLicensesActionResultActionFailureReason_InternalError,
		"none":                                    IosVppAppRevokeLicensesActionResultActionFailureReason_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVppAppRevokeLicensesActionResultActionFailureReason(input)
	return &out, nil
}
