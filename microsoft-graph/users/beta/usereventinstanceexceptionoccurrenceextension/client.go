package usereventinstanceexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEventInstanceExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewUserEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*UserEventInstanceExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usereventinstanceexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEventInstanceExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &UserEventInstanceExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
