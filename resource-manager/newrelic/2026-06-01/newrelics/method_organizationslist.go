package newrelics

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

type OrganizationsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]OrganizationResource
}

type OrganizationsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []OrganizationResource
}

type OrganizationsListOperationOptions struct {
	Location  *string
	UserEmail *string
}

func DefaultOrganizationsListOperationOptions() OrganizationsListOperationOptions {
	return OrganizationsListOperationOptions{}
}

func (o OrganizationsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o OrganizationsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o OrganizationsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Location != nil {
		out.Append("location", fmt.Sprintf("%v", *o.Location))
	}
	if o.UserEmail != nil {
		out.Append("userEmail", fmt.Sprintf("%v", *o.UserEmail))
	}
	return &out
}

type OrganizationsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *OrganizationsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// OrganizationsList ...
func (c NewRelicsClient) OrganizationsList(ctx context.Context, id commonids.SubscriptionId, options OrganizationsListOperationOptions) (result OrganizationsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &OrganizationsListCustomPager{},
		Path:          fmt.Sprintf("%s/providers/NewRelic.Observability/organizations", id.ID()),
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
		Values *[]OrganizationResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// OrganizationsListComplete retrieves all the results into a single object
func (c NewRelicsClient) OrganizationsListComplete(ctx context.Context, id commonids.SubscriptionId, options OrganizationsListOperationOptions) (OrganizationsListCompleteResult, error) {
	return c.OrganizationsListCompleteMatchingPredicate(ctx, id, options, OrganizationResourceOperationPredicate{})
}

// OrganizationsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NewRelicsClient) OrganizationsListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options OrganizationsListOperationOptions, predicate OrganizationResourceOperationPredicate) (result OrganizationsListCompleteResult, err error) {
	items := make([]OrganizationResource, 0)

	resp, err := c.OrganizationsList(ctx, id, options)
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

	result = OrganizationsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
