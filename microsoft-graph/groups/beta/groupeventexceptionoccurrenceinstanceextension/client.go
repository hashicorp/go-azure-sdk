package groupeventexceptionoccurrenceinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEventExceptionOccurrenceInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupEventExceptionOccurrenceInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupeventexceptionoccurrenceinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEventExceptionOccurrenceInstanceExtensionClient: %+v", err)
	}

	return &GroupEventExceptionOccurrenceInstanceExtensionClient{
		Client: client,
	}, nil
}
