package embeddedsimactivationcodepool

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

type ListEmbeddedSIMActivationCodePoolsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.EmbeddedSIMActivationCodePool
}

type ListEmbeddedSIMActivationCodePoolsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.EmbeddedSIMActivationCodePool
}

type ListEmbeddedSIMActivationCodePoolsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEmbeddedSIMActivationCodePoolsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEmbeddedSIMActivationCodePools ...
func (c EmbeddedSIMActivationCodePoolClient) ListEmbeddedSIMActivationCodePools(ctx context.Context) (result ListEmbeddedSIMActivationCodePoolsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEmbeddedSIMActivationCodePoolsCustomPager{},
		Path:       "/deviceManagement/embeddedSIMActivationCodePools",
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
		Values *[]beta.EmbeddedSIMActivationCodePool `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEmbeddedSIMActivationCodePoolsComplete retrieves all the results into a single object
func (c EmbeddedSIMActivationCodePoolClient) ListEmbeddedSIMActivationCodePoolsComplete(ctx context.Context) (ListEmbeddedSIMActivationCodePoolsCompleteResult, error) {
	return c.ListEmbeddedSIMActivationCodePoolsCompleteMatchingPredicate(ctx, EmbeddedSIMActivationCodePoolOperationPredicate{})
}

// ListEmbeddedSIMActivationCodePoolsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EmbeddedSIMActivationCodePoolClient) ListEmbeddedSIMActivationCodePoolsCompleteMatchingPredicate(ctx context.Context, predicate EmbeddedSIMActivationCodePoolOperationPredicate) (result ListEmbeddedSIMActivationCodePoolsCompleteResult, err error) {
	items := make([]beta.EmbeddedSIMActivationCodePool, 0)

	resp, err := c.ListEmbeddedSIMActivationCodePools(ctx)
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

	result = ListEmbeddedSIMActivationCodePoolsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
