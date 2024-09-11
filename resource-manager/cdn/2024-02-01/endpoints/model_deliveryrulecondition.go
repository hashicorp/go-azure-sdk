package endpoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeliveryRuleCondition interface {
}

// RawDeliveryRuleConditionImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeliveryRuleConditionImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalDeliveryRuleConditionImplementation(input []byte) (DeliveryRuleCondition, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeliveryRuleCondition into map[string]interface: %+v", err)
	}

	value, ok := temp["name"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "ClientPort") {
		var out DeliveryRuleClientPortCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleClientPortCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Cookies") {
		var out DeliveryRuleCookiesCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleCookiesCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "HttpVersion") {
		var out DeliveryRuleHTTPVersionCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleHTTPVersionCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "HostName") {
		var out DeliveryRuleHostNameCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleHostNameCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "IsDevice") {
		var out DeliveryRuleIsDeviceCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleIsDeviceCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "PostArgs") {
		var out DeliveryRulePostArgsCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRulePostArgsCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "QueryString") {
		var out DeliveryRuleQueryStringCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleQueryStringCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "RemoteAddress") {
		var out DeliveryRuleRemoteAddressCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleRemoteAddressCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "RequestBody") {
		var out DeliveryRuleRequestBodyCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleRequestBodyCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "RequestHeader") {
		var out DeliveryRuleRequestHeaderCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleRequestHeaderCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "RequestMethod") {
		var out DeliveryRuleRequestMethodCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleRequestMethodCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "RequestScheme") {
		var out DeliveryRuleRequestSchemeCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleRequestSchemeCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "RequestUri") {
		var out DeliveryRuleRequestUriCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleRequestUriCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ServerPort") {
		var out DeliveryRuleServerPortCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleServerPortCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SocketAddr") {
		var out DeliveryRuleSocketAddrCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleSocketAddrCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SslProtocol") {
		var out DeliveryRuleSslProtocolCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleSslProtocolCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "UrlFileExtension") {
		var out DeliveryRuleUrlFileExtensionCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleUrlFileExtensionCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "UrlFileName") {
		var out DeliveryRuleUrlFileNameCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleUrlFileNameCondition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "UrlPath") {
		var out DeliveryRuleUrlPathCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryRuleUrlPathCondition: %+v", err)
		}
		return out, nil
	}

	out := RawDeliveryRuleConditionImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
