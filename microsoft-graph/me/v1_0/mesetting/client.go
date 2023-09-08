package mesetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeSettingClient struct {
	Client *msgraph.Client
}

func NewMeSettingClientWithBaseURI(api sdkEnv.Api) (*MeSettingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mesetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeSettingClient: %+v", err)
	}

	return &MeSettingClient{
		Client: client,
	}, nil
}
