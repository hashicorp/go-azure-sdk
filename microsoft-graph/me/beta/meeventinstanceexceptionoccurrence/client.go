package meeventinstanceexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEventInstanceExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewMeEventInstanceExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*MeEventInstanceExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meeventinstanceexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEventInstanceExceptionOccurrenceClient: %+v", err)
	}

	return &MeEventInstanceExceptionOccurrenceClient{
		Client: client,
	}, nil
}
