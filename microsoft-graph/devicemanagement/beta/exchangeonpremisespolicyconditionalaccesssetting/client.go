package exchangeonpremisespolicyconditionalaccesssetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExchangeOnPremisesPolicyConditionalAccessSettingClient struct {
	Client *msgraph.Client
}

func NewExchangeOnPremisesPolicyConditionalAccessSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*ExchangeOnPremisesPolicyConditionalAccessSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "exchangeonpremisespolicyconditionalaccesssetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExchangeOnPremisesPolicyConditionalAccessSettingClient: %+v", err)
	}

	return &ExchangeOnPremisesPolicyConditionalAccessSettingClient{
		Client: client,
	}, nil
}
