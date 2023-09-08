package groupeventexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEventExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewGroupEventExceptionOccurrenceInstanceClientWithBaseURI(api sdkEnv.Api) (*GroupEventExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupeventexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEventExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &GroupEventExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
