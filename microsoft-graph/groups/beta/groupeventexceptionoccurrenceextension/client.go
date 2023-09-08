package groupeventexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEventExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupEventExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupEventExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupeventexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEventExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &GroupEventExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
