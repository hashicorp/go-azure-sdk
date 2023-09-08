package meprofileemail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeProfileEmailClient struct {
	Client *msgraph.Client
}

func NewMeProfileEmailClientWithBaseURI(api sdkEnv.Api) (*MeProfileEmailClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meprofileemail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeProfileEmailClient: %+v", err)
	}

	return &MeProfileEmailClient{
		Client: client,
	}, nil
}
