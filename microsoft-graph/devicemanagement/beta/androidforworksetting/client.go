package androidforworksetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkSettingClient struct {
	Client *msgraph.Client
}

func NewAndroidForWorkSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*AndroidForWorkSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "androidforworksetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AndroidForWorkSettingClient: %+v", err)
	}

	return &AndroidForWorkSettingClient{
		Client: client,
	}, nil
}
