package calendarevent

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateCalendarEventPermanentDeleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateCalendarEventPermanentDeleteOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateCalendarEventPermanentDeleteOperationOptions() CreateCalendarEventPermanentDeleteOperationOptions {
	return CreateCalendarEventPermanentDeleteOperationOptions{}
}

func (o CreateCalendarEventPermanentDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateCalendarEventPermanentDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateCalendarEventPermanentDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateCalendarEventPermanentDelete - Invoke action permanentDelete. Permanently delete an event and place it in the
// Purges folder in the dumpster in the user's mailbox. Email clients such as Outlook or the Outlook on the web can't
// access permanently deleted items. Unless there's a hold set on the mailbox, the items are permanently deleted after a
// set period of time. For more information about item retention, see Configure Deleted Item retention and Recoverable
// Items quotas.
func (c CalendarEventClient) CreateCalendarEventPermanentDelete(ctx context.Context, id beta.GroupIdCalendarEventId, options CreateCalendarEventPermanentDeleteOperationOptions) (result CreateCalendarEventPermanentDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/permanentDelete", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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
