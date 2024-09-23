package eventexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewEventExceptionOccurrenceClientWithBaseURI(sdkApi sdkEnv.Api) (*EventExceptionOccurrenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "eventexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventExceptionOccurrenceClient: %+v", err)
	}

	return &EventExceptionOccurrenceClient{
		Client: client,
	}, nil
}
