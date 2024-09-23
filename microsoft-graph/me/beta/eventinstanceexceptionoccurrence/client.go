package eventinstanceexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventInstanceExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewEventInstanceExceptionOccurrenceClientWithBaseURI(sdkApi sdkEnv.Api) (*EventInstanceExceptionOccurrenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "eventinstanceexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventInstanceExceptionOccurrenceClient: %+v", err)
	}

	return &EventInstanceExceptionOccurrenceClient{
		Client: client,
	}, nil
}
