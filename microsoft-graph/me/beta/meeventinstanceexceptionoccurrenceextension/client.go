package meeventinstanceexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEventInstanceExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeEventInstanceExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meeventinstanceexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEventInstanceExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &MeEventInstanceExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
