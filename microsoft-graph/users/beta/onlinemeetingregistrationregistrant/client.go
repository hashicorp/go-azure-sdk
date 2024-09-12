package onlinemeetingregistrationregistrant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingRegistrationRegistrantClient struct {
	Client *msgraph.Client
}

func NewOnlineMeetingRegistrationRegistrantClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineMeetingRegistrationRegistrantClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onlinemeetingregistrationregistrant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineMeetingRegistrationRegistrantClient: %+v", err)
	}

	return &OnlineMeetingRegistrationRegistrantClient{
		Client: client,
	}, nil
}
