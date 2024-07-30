package calendargroupcalendareventinstanceexceptionoccurrence

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceTentativelyAcceptOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceTentativelyAccept ...
func (c CalendarGroupCalendarEventInstanceExceptionOccurrenceClient) CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceTentativelyAccept(ctx context.Context, id MeCalendarGroupIdCalendarIdEventIdInstanceIdExceptionOccurrenceId, input CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceTentativelyAcceptRequest) (result CreateCalendarGroupCalendarEventInstanceExceptionOccurrenceTentativelyAcceptOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/tentativelyAccept", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
