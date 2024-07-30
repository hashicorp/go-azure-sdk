package calendarcalendarviewexceptionoccurrence

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateCalendarCalendarViewExceptionOccurrenceForwardOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateCalendarCalendarViewExceptionOccurrenceForward ...
func (c CalendarCalendarViewExceptionOccurrenceClient) CreateCalendarCalendarViewExceptionOccurrenceForward(ctx context.Context, id GroupIdCalendarCalendarViewIdExceptionOccurrenceId, input CreateCalendarCalendarViewExceptionOccurrenceForwardRequest) (result CreateCalendarCalendarViewExceptionOccurrenceForwardOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/forward", id.ID()),
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
