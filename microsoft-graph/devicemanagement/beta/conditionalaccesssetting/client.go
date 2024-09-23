package conditionalaccesssetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessSettingClient struct {
	Client *msgraph.Client
}

func NewConditionalAccessSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*ConditionalAccessSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conditionalaccesssetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConditionalAccessSettingClient: %+v", err)
	}

	return &ConditionalAccessSettingClient{
		Client: client,
	}, nil
}
