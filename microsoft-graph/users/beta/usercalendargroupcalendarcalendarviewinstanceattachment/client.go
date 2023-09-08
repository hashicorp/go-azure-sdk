package usercalendargroupcalendarcalendarviewinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarCalendarViewInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarCalendarViewInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarCalendarViewInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendarcalendarviewinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarCalendarViewInstanceAttachmentClient: %+v", err)
	}

	return &UserCalendarGroupCalendarCalendarViewInstanceAttachmentClient{
		Client: client,
	}, nil
}
