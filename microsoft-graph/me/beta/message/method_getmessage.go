package message

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMessageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.Message
}

type GetMessageOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetMessageOperationOptions() GetMessageOperationOptions {
	return GetMessageOperationOptions{}
}

func (o GetMessageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetMessageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetMessageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetMessage - Get singleValueLegacyExtendedProperty. You can get a single resource instance expanded with a specific
// extended property, or a collection of resource instances that include extended properties matching a filter. Using
// the query parameter $expand allows you to get the specified resource instance expanded with a specific extended
// property. Use a $filter and eq operator on the id property to specify the extended property. This is currently the
// only way to get the singleValueLegacyExtendedProperty object that represents an extended property. To get resource
// instances that have certain extended properties, use the $filter query parameter and apply an eq operator on the id
// property. In addition, for numeric extended properties, apply one of the following operators on the value property:
// eq, ne,ge, gt, le, or lt. For string-typed extended properties, apply a contains, startswith, eq, or ne operator on
// value. Filtering the string name (Name) in the id of an extended property is case-sensitive. Filtering the value
// property of an extended property is case-insensitive. The following user resources are supported: As well as the
// following group resources: See Extended properties overview for more information about when to use open extensions or
// extended properties, and how to specify extended properties.
func (c MessageClient) GetMessage(ctx context.Context, id beta.MeMessageId, options GetMessageOperationOptions) (result GetMessageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalMessageImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
