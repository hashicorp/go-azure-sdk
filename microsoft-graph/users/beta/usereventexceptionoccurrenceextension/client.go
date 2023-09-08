package usereventexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEventExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewUserEventExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*UserEventExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usereventexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEventExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &UserEventExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
