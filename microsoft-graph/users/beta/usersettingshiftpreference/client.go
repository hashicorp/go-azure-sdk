package usersettingshiftpreference

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSettingShiftPreferenceClient struct {
	Client *msgraph.Client
}

func NewUserSettingShiftPreferenceClientWithBaseURI(api sdkEnv.Api) (*UserSettingShiftPreferenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usersettingshiftpreference", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserSettingShiftPreferenceClient: %+v", err)
	}

	return &UserSettingShiftPreferenceClient{
		Client: client,
	}, nil
}
