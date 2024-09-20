package triggers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByDataBoxEdgeDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Trigger
}

type ListByDataBoxEdgeDeviceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Trigger
}

type ListByDataBoxEdgeDeviceOperationOptions struct {
	Filter *string
}

func DefaultListByDataBoxEdgeDeviceOperationOptions() ListByDataBoxEdgeDeviceOperationOptions {
	return ListByDataBoxEdgeDeviceOperationOptions{}
}

func (o ListByDataBoxEdgeDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByDataBoxEdgeDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListByDataBoxEdgeDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type ListByDataBoxEdgeDeviceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByDataBoxEdgeDeviceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByDataBoxEdgeDevice ...
func (c TriggersClient) ListByDataBoxEdgeDevice(ctx context.Context, id DataBoxEdgeDeviceId, options ListByDataBoxEdgeDeviceOperationOptions) (result ListByDataBoxEdgeDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByDataBoxEdgeDeviceCustomPager{},
		Path:          fmt.Sprintf("%s/triggers", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]Trigger, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := UnmarshalTriggerImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for Trigger (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListByDataBoxEdgeDeviceComplete retrieves all the results into a single object
func (c TriggersClient) ListByDataBoxEdgeDeviceComplete(ctx context.Context, id DataBoxEdgeDeviceId, options ListByDataBoxEdgeDeviceOperationOptions) (ListByDataBoxEdgeDeviceCompleteResult, error) {
	return c.ListByDataBoxEdgeDeviceCompleteMatchingPredicate(ctx, id, options, TriggerOperationPredicate{})
}

// ListByDataBoxEdgeDeviceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TriggersClient) ListByDataBoxEdgeDeviceCompleteMatchingPredicate(ctx context.Context, id DataBoxEdgeDeviceId, options ListByDataBoxEdgeDeviceOperationOptions, predicate TriggerOperationPredicate) (result ListByDataBoxEdgeDeviceCompleteResult, err error) {
	items := make([]Trigger, 0)

	resp, err := c.ListByDataBoxEdgeDevice(ctx, id, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListByDataBoxEdgeDeviceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
