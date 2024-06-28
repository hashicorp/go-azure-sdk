package managedhsmkeys

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagedHsmKey
}

type ListVersionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ManagedHsmKey
}

type ListVersionsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListVersionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVersions ...
func (c ManagedHsmKeysClient) ListVersions(ctx context.Context, id KeyId) (result ListVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListVersionsCustomPager{},
		Path:       fmt.Sprintf("%s/versions", id.ID()),
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
		Values *[]ManagedHsmKey `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVersionsComplete retrieves all the results into a single object
func (c ManagedHsmKeysClient) ListVersionsComplete(ctx context.Context, id KeyId) (ListVersionsCompleteResult, error) {
	return c.ListVersionsCompleteMatchingPredicate(ctx, id, ManagedHsmKeyOperationPredicate{})
}

// ListVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedHsmKeysClient) ListVersionsCompleteMatchingPredicate(ctx context.Context, id KeyId, predicate ManagedHsmKeyOperationPredicate) (result ListVersionsCompleteResult, err error) {
	items := make([]ManagedHsmKey, 0)

	resp, err := c.ListVersions(ctx, id)
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

	result = ListVersionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
