package dataconnectors

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CcpAuthConfig interface {
	CcpAuthConfig() BaseCcpAuthConfigImpl
}

var _ CcpAuthConfig = BaseCcpAuthConfigImpl{}

type BaseCcpAuthConfigImpl struct {
	Type CcpAuthType `json:"type"`
}

func (s BaseCcpAuthConfigImpl) CcpAuthConfig() BaseCcpAuthConfigImpl {
	return s
}

var _ CcpAuthConfig = RawCcpAuthConfigImpl{}

// RawCcpAuthConfigImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCcpAuthConfigImpl struct {
	ccpAuthConfig BaseCcpAuthConfigImpl
	Type          string
	Values        map[string]interface{}
}

func (s RawCcpAuthConfigImpl) CcpAuthConfig() BaseCcpAuthConfigImpl {
	return s.ccpAuthConfig
}

func UnmarshalCcpAuthConfigImplementation(input []byte) (CcpAuthConfig, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CcpAuthConfig into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AWS") {
		var out AWSAuthModel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AWSAuthModel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "APIKey") {
		var out ApiKeyAuthModel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApiKeyAuthModel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Basic") {
		var out BasicAuthModel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BasicAuthModel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "GCP") {
		var out GCPAuthModel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GCPAuthModel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ServiceBus") {
		var out GenericBlobSbsAuthModel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GenericBlobSbsAuthModel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "GitHub") {
		var out GitHubAuthModel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GitHubAuthModel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "JwtToken") {
		var out JwtAuthModel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JwtAuthModel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "None") {
		var out NoneAuthModel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NoneAuthModel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "OAuth2") {
		var out OAuthModel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OAuthModel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Oracle") {
		var out OracleAuthModel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OracleAuthModel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Session") {
		var out SessionAuthModel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SessionAuthModel: %+v", err)
		}
		return out, nil
	}

	var parent BaseCcpAuthConfigImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCcpAuthConfigImpl: %+v", err)
	}

	return RawCcpAuthConfigImpl{
		ccpAuthConfig: parent,
		Type:          value,
		Values:        temp,
	}, nil

}
