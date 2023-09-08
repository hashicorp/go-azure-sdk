package meprofileaward

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeProfileAwardClient struct {
	Client *msgraph.Client
}

func NewMeProfileAwardClientWithBaseURI(api sdkEnv.Api) (*MeProfileAwardClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meprofileaward", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeProfileAwardClient: %+v", err)
	}

	return &MeProfileAwardClient{
		Client: client,
	}, nil
}
