package meonlinemeetingregistrationregistrant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnlineMeetingRegistrationRegistrantClient struct {
	Client *msgraph.Client
}

func NewMeOnlineMeetingRegistrationRegistrantClientWithBaseURI(api sdkEnv.Api) (*MeOnlineMeetingRegistrationRegistrantClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonlinemeetingregistrationregistrant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnlineMeetingRegistrationRegistrantClient: %+v", err)
	}

	return &MeOnlineMeetingRegistrationRegistrantClient{
		Client: client,
	}, nil
}
