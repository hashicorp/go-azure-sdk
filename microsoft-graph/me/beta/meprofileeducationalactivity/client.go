package meprofileeducationalactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeProfileEducationalActivityClient struct {
	Client *msgraph.Client
}

func NewMeProfileEducationalActivityClientWithBaseURI(api sdkEnv.Api) (*MeProfileEducationalActivityClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meprofileeducationalactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeProfileEducationalActivityClient: %+v", err)
	}

	return &MeProfileEducationalActivityClient{
		Client: client,
	}, nil
}
