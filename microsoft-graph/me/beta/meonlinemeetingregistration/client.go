package meonlinemeetingregistration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnlineMeetingRegistrationClient struct {
	Client *msgraph.Client
}

func NewMeOnlineMeetingRegistrationClientWithBaseURI(api sdkEnv.Api) (*MeOnlineMeetingRegistrationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonlinemeetingregistration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnlineMeetingRegistrationClient: %+v", err)
	}

	return &MeOnlineMeetingRegistrationClient{
		Client: client,
	}, nil
}
