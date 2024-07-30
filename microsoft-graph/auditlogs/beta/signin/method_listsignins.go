package signin

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

type ListSignInsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SignIn
}

type ListSignInsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SignIn
}

type ListSignInsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSignInsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSignIns ...
func (c SignInClient) ListSignIns(ctx context.Context) (result ListSignInsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSignInsCustomPager{},
		Path:       "/auditLogs/signIns",
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
		Values *[]beta.SignIn `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSignInsComplete retrieves all the results into a single object
func (c SignInClient) ListSignInsComplete(ctx context.Context) (ListSignInsCompleteResult, error) {
	return c.ListSignInsCompleteMatchingPredicate(ctx, SignInOperationPredicate{})
}

// ListSignInsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SignInClient) ListSignInsCompleteMatchingPredicate(ctx context.Context, predicate SignInOperationPredicate) (result ListSignInsCompleteResult, err error) {
	items := make([]beta.SignIn, 0)

	resp, err := c.ListSignIns(ctx)
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

	result = ListSignInsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
