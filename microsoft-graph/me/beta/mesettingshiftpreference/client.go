package mesettingshiftpreference

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeSettingShiftPreferenceClient struct {
	Client *msgraph.Client
}

func NewMeSettingShiftPreferenceClientWithBaseURI(api sdkEnv.Api) (*MeSettingShiftPreferenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mesettingshiftpreference", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeSettingShiftPreferenceClient: %+v", err)
	}

	return &MeSettingShiftPreferenceClient{
		Client: client,
	}, nil
}
