package templatedevicetemplatedeviceinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateDeviceTemplateDeviceInstanceClient struct {
	Client *msgraph.Client
}

func NewTemplateDeviceTemplateDeviceInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*TemplateDeviceTemplateDeviceInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "templatedevicetemplatedeviceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TemplateDeviceTemplateDeviceInstanceClient: %+v", err)
	}

	return &TemplateDeviceTemplateDeviceInstanceClient{
		Client: client,
	}, nil
}
