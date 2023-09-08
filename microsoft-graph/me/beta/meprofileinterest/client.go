package meprofileinterest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeProfileInterestClient struct {
	Client *msgraph.Client
}

func NewMeProfileInterestClientWithBaseURI(api sdkEnv.Api) (*MeProfileInterestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meprofileinterest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeProfileInterestClient: %+v", err)
	}

	return &MeProfileInterestClient{
		Client: client,
	}, nil
}
