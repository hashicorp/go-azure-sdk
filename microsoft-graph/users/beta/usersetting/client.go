package usersetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSettingClient struct {
	Client *msgraph.Client
}

func NewUserSettingClientWithBaseURI(api sdkEnv.Api) (*UserSettingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usersetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserSettingClient: %+v", err)
	}

	return &UserSettingClient{
		Client: client,
	}, nil
}
