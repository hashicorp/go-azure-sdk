package managedinstancedtcs

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

type ListByManagedInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagedInstanceDtc
}

type ListByManagedInstanceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ManagedInstanceDtc
}

type ListByManagedInstanceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByManagedInstanceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByManagedInstance ...
func (c ManagedInstanceDtcsClient) ListByManagedInstance(ctx context.Context, id commonids.SqlManagedInstanceId) (result ListByManagedInstanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByManagedInstanceCustomPager{},
		Path:       fmt.Sprintf("%s/dtc", id.ID()),
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
		Values *[]ManagedInstanceDtc `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByManagedInstanceComplete retrieves all the results into a single object
func (c ManagedInstanceDtcsClient) ListByManagedInstanceComplete(ctx context.Context, id commonids.SqlManagedInstanceId) (ListByManagedInstanceCompleteResult, error) {
	return c.ListByManagedInstanceCompleteMatchingPredicate(ctx, id, ManagedInstanceDtcOperationPredicate{})
}

// ListByManagedInstanceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedInstanceDtcsClient) ListByManagedInstanceCompleteMatchingPredicate(ctx context.Context, id commonids.SqlManagedInstanceId, predicate ManagedInstanceDtcOperationPredicate) (result ListByManagedInstanceCompleteResult, err error) {
	items := make([]ManagedInstanceDtc, 0)

	resp, err := c.ListByManagedInstance(ctx, id)
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

	result = ListByManagedInstanceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
