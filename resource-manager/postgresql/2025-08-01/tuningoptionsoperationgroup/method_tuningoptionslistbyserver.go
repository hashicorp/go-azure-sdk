package tuningoptionsoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TuningOptionsListByServerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Resource
}

type TuningOptionsListByServerCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Resource
}

type TuningOptionsListByServerCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *TuningOptionsListByServerCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// TuningOptionsListByServer ...
func (c TuningOptionsOperationGroupClient) TuningOptionsListByServer(ctx context.Context, id FlexibleServerId) (result TuningOptionsListByServerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &TuningOptionsListByServerCustomPager{},
		Path:       fmt.Sprintf("%s/tuningOptions", id.ID()),
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
		Values *[]Resource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// TuningOptionsListByServerComplete retrieves all the results into a single object
func (c TuningOptionsOperationGroupClient) TuningOptionsListByServerComplete(ctx context.Context, id FlexibleServerId) (TuningOptionsListByServerCompleteResult, error) {
	return c.TuningOptionsListByServerCompleteMatchingPredicate(ctx, id, ResourceOperationPredicate{})
}

// TuningOptionsListByServerCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TuningOptionsOperationGroupClient) TuningOptionsListByServerCompleteMatchingPredicate(ctx context.Context, id FlexibleServerId, predicate ResourceOperationPredicate) (result TuningOptionsListByServerCompleteResult, err error) {
	items := make([]Resource, 0)

	resp, err := c.TuningOptionsListByServer(ctx, id)
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

	result = TuningOptionsListByServerCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
