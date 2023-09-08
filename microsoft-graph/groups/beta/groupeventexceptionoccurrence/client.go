package groupeventexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEventExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewGroupEventExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*GroupEventExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupeventexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEventExceptionOccurrenceClient: %+v", err)
	}

	return &GroupEventExceptionOccurrenceClient{
		Client: client,
	}, nil
}
