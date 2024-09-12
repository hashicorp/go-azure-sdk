package profilephone

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfilePhoneClient struct {
	Client *msgraph.Client
}

func NewProfilePhoneClientWithBaseURI(sdkApi sdkEnv.Api) (*ProfilePhoneClient, error) {
	client, err := msgraph.NewClient(sdkApi, "profilephone", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProfilePhoneClient: %+v", err)
	}

	return &ProfilePhoneClient{
		Client: client,
	}, nil
}
