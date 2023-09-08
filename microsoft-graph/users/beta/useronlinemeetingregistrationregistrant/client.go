package useronlinemeetingregistrationregistrant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnlineMeetingRegistrationRegistrantClient struct {
	Client *msgraph.Client
}

func NewUserOnlineMeetingRegistrationRegistrantClientWithBaseURI(api sdkEnv.Api) (*UserOnlineMeetingRegistrationRegistrantClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronlinemeetingregistrationregistrant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnlineMeetingRegistrationRegistrantClient: %+v", err)
	}

	return &UserOnlineMeetingRegistrationRegistrantClient{
		Client: client,
	}, nil
}
