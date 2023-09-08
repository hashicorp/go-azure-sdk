package memailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewMeMailboxSettingClientWithBaseURI(api sdkEnv.Api) (*MeMailboxSettingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeMailboxSettingClient: %+v", err)
	}

	return &MeMailboxSettingClient{
		Client: client,
	}, nil
}
