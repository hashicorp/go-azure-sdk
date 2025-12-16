package dynatracesinglesignonresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SingleSignOnListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DynatraceSingleSignOnResource
}

type SingleSignOnListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DynatraceSingleSignOnResource
}

type SingleSignOnListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SingleSignOnListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SingleSignOnList ...
func (c DynatraceSingleSignOnResourcesClient) SingleSignOnList(ctx context.Context, id MonitorId) (result SingleSignOnListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &SingleSignOnListCustomPager{},
		Path:       fmt.Sprintf("%s/singleSignOnConfigurations", id.ID()),
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
		Values *[]DynatraceSingleSignOnResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SingleSignOnListComplete retrieves all the results into a single object
func (c DynatraceSingleSignOnResourcesClient) SingleSignOnListComplete(ctx context.Context, id MonitorId) (SingleSignOnListCompleteResult, error) {
	return c.SingleSignOnListCompleteMatchingPredicate(ctx, id, DynatraceSingleSignOnResourceOperationPredicate{})
}

// SingleSignOnListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DynatraceSingleSignOnResourcesClient) SingleSignOnListCompleteMatchingPredicate(ctx context.Context, id MonitorId, predicate DynatraceSingleSignOnResourceOperationPredicate) (result SingleSignOnListCompleteResult, err error) {
	items := make([]DynatraceSingleSignOnResource, 0)

	resp, err := c.SingleSignOnList(ctx, id)
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

	result = SingleSignOnListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
