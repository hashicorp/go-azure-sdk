package usernotification

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserNotificationClient struct {
	Client *msgraph.Client
}

func NewUserNotificationClientWithBaseURI(api sdkEnv.Api) (*UserNotificationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usernotification", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserNotificationClient: %+v", err)
	}

	return &UserNotificationClient{
		Client: client,
	}, nil
}
