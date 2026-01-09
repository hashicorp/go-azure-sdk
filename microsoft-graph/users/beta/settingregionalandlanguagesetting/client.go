package settingregionalandlanguagesetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingRegionalAndLanguageSettingClient struct {
	Client *msgraph.Client
}

func NewSettingRegionalAndLanguageSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SettingRegionalAndLanguageSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "settingregionalandlanguagesetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SettingRegionalAndLanguageSettingClient: %+v", err)
	}

	return &SettingRegionalAndLanguageSettingClient{
		Client: client,
	}, nil
}
