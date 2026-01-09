package settingcontactmergesuggestion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingContactMergeSuggestionClient struct {
	Client *msgraph.Client
}

func NewSettingContactMergeSuggestionClientWithBaseURI(sdkApi sdkEnv.Api) (*SettingContactMergeSuggestionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "settingcontactmergesuggestion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SettingContactMergeSuggestionClient: %+v", err)
	}

	return &SettingContactMergeSuggestionClient{
		Client: client,
	}, nil
}
