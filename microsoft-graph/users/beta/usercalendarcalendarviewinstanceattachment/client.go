package usercalendarcalendarviewinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarCalendarViewInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserCalendarCalendarViewInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserCalendarCalendarViewInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarcalendarviewinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarCalendarViewInstanceAttachmentClient: %+v", err)
	}

	return &UserCalendarCalendarViewInstanceAttachmentClient{
		Client: client,
	}, nil
}
