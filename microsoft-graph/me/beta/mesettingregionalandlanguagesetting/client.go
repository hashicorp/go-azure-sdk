package mesettingregionalandlanguagesetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeSettingRegionalAndLanguageSettingClient struct {
	Client *msgraph.Client
}

func NewMeSettingRegionalAndLanguageSettingClientWithBaseURI(api sdkEnv.Api) (*MeSettingRegionalAndLanguageSettingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mesettingregionalandlanguagesetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeSettingRegionalAndLanguageSettingClient: %+v", err)
	}

	return &MeSettingRegionalAndLanguageSettingClient{
		Client: client,
	}, nil
}
