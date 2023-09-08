package usereventexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEventExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewUserEventExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*UserEventExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usereventexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEventExceptionOccurrenceClient: %+v", err)
	}

	return &UserEventExceptionOccurrenceClient{
		Client: client,
	}, nil
}
