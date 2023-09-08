package mesettingiteminsight

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeSettingItemInsightClient struct {
	Client *msgraph.Client
}

func NewMeSettingItemInsightClientWithBaseURI(api sdkEnv.Api) (*MeSettingItemInsightClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mesettingiteminsight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeSettingItemInsightClient: %+v", err)
	}

	return &MeSettingItemInsightClient{
		Client: client,
	}, nil
}
