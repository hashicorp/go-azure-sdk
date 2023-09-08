package groupsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSettingClient struct {
	Client *msgraph.Client
}

func NewGroupSettingClientWithBaseURI(api sdkEnv.Api) (*GroupSettingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSettingClient: %+v", err)
	}

	return &GroupSettingClient{
		Client: client,
	}, nil
}
