package meprofileproject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeProfileProjectClient struct {
	Client *msgraph.Client
}

func NewMeProfileProjectClientWithBaseURI(api sdkEnv.Api) (*MeProfileProjectClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meprofileproject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeProfileProjectClient: %+v", err)
	}

	return &MeProfileProjectClient{
		Client: client,
	}, nil
}
