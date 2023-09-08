package meprofilephone

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeProfilePhoneClient struct {
	Client *msgraph.Client
}

func NewMeProfilePhoneClientWithBaseURI(api sdkEnv.Api) (*MeProfilePhoneClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meprofilephone", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeProfilePhoneClient: %+v", err)
	}

	return &MeProfilePhoneClient{
		Client: client,
	}, nil
}
