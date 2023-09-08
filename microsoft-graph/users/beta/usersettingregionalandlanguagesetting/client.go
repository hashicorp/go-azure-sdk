package usersettingregionalandlanguagesetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSettingRegionalAndLanguageSettingClient struct {
	Client *msgraph.Client
}

func NewUserSettingRegionalAndLanguageSettingClientWithBaseURI(api sdkEnv.Api) (*UserSettingRegionalAndLanguageSettingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usersettingregionalandlanguagesetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserSettingRegionalAndLanguageSettingClient: %+v", err)
	}

	return &UserSettingRegionalAndLanguageSettingClient{
		Client: client,
	}, nil
}
