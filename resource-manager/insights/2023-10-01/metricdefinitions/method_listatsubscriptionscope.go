package metricdefinitions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAtSubscriptionScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *SubscriptionScopeMetricDefinitionCollection
}

type ListAtSubscriptionScopeOperationOptions struct {
	Metricnamespace *string
	Region          *string
}

func DefaultListAtSubscriptionScopeOperationOptions() ListAtSubscriptionScopeOperationOptions {
	return ListAtSubscriptionScopeOperationOptions{}
}

func (o ListAtSubscriptionScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAtSubscriptionScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListAtSubscriptionScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Metricnamespace != nil {
		out.Append("metricnamespace", fmt.Sprintf("%v", *o.Metricnamespace))
	}
	if o.Region != nil {
		out.Append("region", fmt.Sprintf("%v", *o.Region))
	}
	return &out
}

// ListAtSubscriptionScope ...
func (c MetricDefinitionsClient) ListAtSubscriptionScope(ctx context.Context, id commonids.SubscriptionId, options ListAtSubscriptionScopeOperationOptions) (result ListAtSubscriptionScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.Insights/metricDefinitions", id.ID()),
		OptionsObject: options,
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

	var model SubscriptionScopeMetricDefinitionCollection
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
