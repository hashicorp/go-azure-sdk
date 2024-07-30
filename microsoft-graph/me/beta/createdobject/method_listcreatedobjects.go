package createdobject

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

type ListCreatedObjectsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListCreatedObjectsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListCreatedObjectsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCreatedObjectsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCreatedObjects ...
func (c CreatedObjectClient) ListCreatedObjects(ctx context.Context) (result ListCreatedObjectsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCreatedObjectsCustomPager{},
		Path:       "/me/createdObjects",
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
		Values *[]beta.DirectoryObject `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCreatedObjectsComplete retrieves all the results into a single object
func (c CreatedObjectClient) ListCreatedObjectsComplete(ctx context.Context) (ListCreatedObjectsCompleteResult, error) {
	return c.ListCreatedObjectsCompleteMatchingPredicate(ctx, DirectoryObjectOperationPredicate{})
}

// ListCreatedObjectsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CreatedObjectClient) ListCreatedObjectsCompleteMatchingPredicate(ctx context.Context, predicate DirectoryObjectOperationPredicate) (result ListCreatedObjectsCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListCreatedObjects(ctx)
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

	result = ListCreatedObjectsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
