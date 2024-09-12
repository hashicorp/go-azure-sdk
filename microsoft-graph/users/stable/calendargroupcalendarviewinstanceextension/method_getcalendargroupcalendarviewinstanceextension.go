package calendargroupcalendarviewinstanceextension

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetCalendarGroupCalendarViewInstanceExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.Extension
}

type GetCalendarGroupCalendarViewInstanceExtensionOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetCalendarGroupCalendarViewInstanceExtensionOperationOptions() GetCalendarGroupCalendarViewInstanceExtensionOperationOptions {
	return GetCalendarGroupCalendarViewInstanceExtensionOperationOptions{}
}

func (o GetCalendarGroupCalendarViewInstanceExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCalendarGroupCalendarViewInstanceExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetCalendarGroupCalendarViewInstanceExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetCalendarGroupCalendarViewInstanceExtension - Get extensions from users. The collection of open extensions defined
// for the event. Nullable.
func (c CalendarGroupCalendarViewInstanceExtensionClient) GetCalendarGroupCalendarViewInstanceExtension(ctx context.Context, id stable.UserIdCalendarGroupIdCalendarIdCalendarViewIdInstanceIdExtensionId, options GetCalendarGroupCalendarViewInstanceExtensionOperationOptions) (result GetCalendarGroupCalendarViewInstanceExtensionOperationResponse, err error) {
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
	model, err := stable.UnmarshalExtensionImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
