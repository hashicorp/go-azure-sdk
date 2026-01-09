package templatemigratabletosetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateMigratableToSettingClient struct {
	Client *msgraph.Client
}

func NewTemplateMigratableToSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*TemplateMigratableToSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "templatemigratabletosetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TemplateMigratableToSettingClient: %+v", err)
	}

	return &TemplateMigratableToSettingClient{
		Client: client,
	}, nil
}
