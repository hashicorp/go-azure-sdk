package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer string

const (
	Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer_BlockImageAndTextFile Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer = "blockImageAndTextFile"
	Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer_BlockImageFile        Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer = "blockImageFile"
	Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer_BlockNone             Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer = "blockNone"
	Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer_BlockTextFile         Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer = "blockTextFile"
	Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer_NotConfigured         Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer = "notConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer_BlockImageAndTextFile),
		string(Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer_BlockImageFile),
		string(Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer_BlockNone),
		string(Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer_BlockTextFile),
		string(Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer_NotConfigured),
	}
}

func (s *Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer(input string) (*Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer{
		"blockimageandtextfile": Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer_BlockImageAndTextFile,
		"blockimagefile":        Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer_BlockImageFile,
		"blocknone":             Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer_BlockNone,
		"blocktextfile":         Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer_BlockTextFile,
		"notconfigured":         Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer(input)
	return &out, nil
}
