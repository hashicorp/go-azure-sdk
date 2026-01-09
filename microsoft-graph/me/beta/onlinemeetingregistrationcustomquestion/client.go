package onlinemeetingregistrationcustomquestion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingRegistrationCustomQuestionClient struct {
	Client *msgraph.Client
}

func NewOnlineMeetingRegistrationCustomQuestionClientWithBaseURI(sdkApi sdkEnv.Api) (*OnlineMeetingRegistrationCustomQuestionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onlinemeetingregistrationcustomquestion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnlineMeetingRegistrationCustomQuestionClient: %+v", err)
	}

	return &OnlineMeetingRegistrationCustomQuestionClient{
		Client: client,
	}, nil
}
