package incidententities

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Entity interface {
	Entity() BaseEntityImpl
}

var _ Entity = BaseEntityImpl{}

type BaseEntityImpl struct {
	Id         *string                `json:"id,omitempty"`
	Kind       EntityKindEnum         `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s BaseEntityImpl) Entity() BaseEntityImpl {
	return s
}

var _ Entity = RawEntityImpl{}

// RawEntityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEntityImpl struct {
	entity BaseEntityImpl
	Type   string
	Values map[string]interface{}
}

func (s RawEntityImpl) Entity() BaseEntityImpl {
	return s.entity
}

func UnmarshalEntityImplementation(input []byte) (Entity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Entity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Account") {
		var out AccountEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccountEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureResource") {
		var out AzureResourceEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureResourceEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "CloudApplication") {
		var out CloudApplicationEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudApplicationEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DnsResolution") {
		var out DnsEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DnsEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "File") {
		var out FileEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FileEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "FileHash") {
		var out FileHashEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FileHashEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Host") {
		var out HostEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HostEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Bookmark") {
		var out HuntingBookmark
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HuntingBookmark: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Ip") {
		var out IPEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IPEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "IoTDevice") {
		var out IoTDeviceEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IoTDeviceEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MailCluster") {
		var out MailClusterEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailClusterEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MailMessage") {
		var out MailMessageEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailMessageEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Mailbox") {
		var out MailboxEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailboxEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Malware") {
		var out MalwareEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MalwareEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Process") {
		var out ProcessEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProcessEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "RegistryKey") {
		var out RegistryKeyEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RegistryKeyEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "RegistryValue") {
		var out RegistryValueEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RegistryValueEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SecurityAlert") {
		var out SecurityAlert
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAlert: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SecurityGroup") {
		var out SecurityGroupEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityGroupEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SubmissionMail") {
		var out SubmissionMailEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SubmissionMailEntity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Url") {
		var out URLEntity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into URLEntity: %+v", err)
		}
		return out, nil
	}

	var parent BaseEntityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEntityImpl: %+v", err)
	}

	return RawEntityImpl{
		entity: parent,
		Type:   value,
		Values: temp,
	}, nil

}
