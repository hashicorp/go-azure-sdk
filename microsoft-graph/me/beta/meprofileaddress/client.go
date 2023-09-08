package meprofileaddress

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeProfileAddressClient struct {
	Client *msgraph.Client
}

func NewMeProfileAddressClientWithBaseURI(api sdkEnv.Api) (*MeProfileAddressClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meprofileaddress", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeProfileAddressClient: %+v", err)
	}

	return &MeProfileAddressClient{
		Client: client,
	}, nil
}
