package useronlinemeetingregistration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnlineMeetingRegistrationClient struct {
	Client *msgraph.Client
}

func NewUserOnlineMeetingRegistrationClientWithBaseURI(api sdkEnv.Api) (*UserOnlineMeetingRegistrationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronlinemeetingregistration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnlineMeetingRegistrationClient: %+v", err)
	}

	return &UserOnlineMeetingRegistrationClient{
		Client: client,
	}, nil
}
