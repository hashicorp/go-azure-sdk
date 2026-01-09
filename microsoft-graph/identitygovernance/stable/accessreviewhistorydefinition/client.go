package accessreviewhistorydefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewHistoryDefinitionClient struct {
	Client *msgraph.Client
}

func NewAccessReviewHistoryDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewHistoryDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewhistorydefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewHistoryDefinitionClient: %+v", err)
	}

	return &AccessReviewHistoryDefinitionClient{
		Client: client,
	}, nil
}
