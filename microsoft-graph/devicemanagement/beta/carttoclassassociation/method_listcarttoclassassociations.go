package carttoclassassociation

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

type ListCartToClassAssociationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CartToClassAssociation
}

type ListCartToClassAssociationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CartToClassAssociation
}

type ListCartToClassAssociationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCartToClassAssociationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCartToClassAssociations ...
func (c CartToClassAssociationClient) ListCartToClassAssociations(ctx context.Context) (result ListCartToClassAssociationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCartToClassAssociationsCustomPager{},
		Path:       "/deviceManagement/cartToClassAssociations",
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
		Values *[]beta.CartToClassAssociation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCartToClassAssociationsComplete retrieves all the results into a single object
func (c CartToClassAssociationClient) ListCartToClassAssociationsComplete(ctx context.Context) (ListCartToClassAssociationsCompleteResult, error) {
	return c.ListCartToClassAssociationsCompleteMatchingPredicate(ctx, CartToClassAssociationOperationPredicate{})
}

// ListCartToClassAssociationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CartToClassAssociationClient) ListCartToClassAssociationsCompleteMatchingPredicate(ctx context.Context, predicate CartToClassAssociationOperationPredicate) (result ListCartToClassAssociationsCompleteResult, err error) {
	items := make([]beta.CartToClassAssociation, 0)

	resp, err := c.ListCartToClassAssociations(ctx)
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

	result = ListCartToClassAssociationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
