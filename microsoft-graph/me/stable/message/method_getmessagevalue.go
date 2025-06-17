package message

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMessageValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetMessageValueOperationOptions struct {
	IncludeHiddenMessages *string
	Metadata              *odata.Metadata
	RetryFunc             client.RequestRetryFunc
}

func DefaultGetMessageValueOperationOptions() GetMessageValueOperationOptions {
	return GetMessageValueOperationOptions{}
}

func (o GetMessageValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetMessageValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetMessageValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IncludeHiddenMessages != nil {
		out.Append("includeHiddenMessages", fmt.Sprintf("%v", *o.IncludeHiddenMessages))
	}
	return &out
}

// GetMessageValue - List messages. Get the messages in the signed-in user's mailbox (including the Deleted Items and
// Clutter folders). Depending on the page size and mailbox data, getting messages from a mailbox can incur multiple
// requests. The default page size is 10 messages. Use $top to customize the page size, within the range of 1 and 1000.
// To improve the operation response time, use $select to specify the exact properties you need; see example 1 below.
// Fine-tune the values for $select and $top, especially when you must use a larger page size, as returning a page with
// hundreds of messages each with a full response payload may trigger the gateway timeout (HTTP 504). To get the next
// page of messages, simply apply the entire URL returned in @odata.nextLink to the next get-messages request. This URL
// includes any query parameters you may have specified in the initial request. Do not try to extract the $skip value
// from the @odata.nextLink URL to manipulate responses. This API uses the $skip value to keep count of all the items it
// has gone through in the user's mailbox to return a page of message-type items. It's therefore possible that even in
// the initial response, the $skip value is larger than the page size. For more information, see Paging Microsoft Graph
// data in your app. Currently, this operation returns message bodies in only HTML format. There are two scenarios where
// an app can get messages in another user's mail folder
func (c MessageClient) GetMessageValue(ctx context.Context, id stable.MeMessageId, options GetMessageValueOperationOptions) (result GetMessageValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$value", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
