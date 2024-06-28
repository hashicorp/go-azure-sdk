package managedvmsizes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedUnsupportedVMSizesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagedVMSize
}

type ManagedUnsupportedVMSizesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ManagedVMSize
}

type ManagedUnsupportedVMSizesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ManagedUnsupportedVMSizesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ManagedUnsupportedVMSizesList ...
func (c ManagedVMSizesClient) ManagedUnsupportedVMSizesList(ctx context.Context, id LocationId) (result ManagedUnsupportedVMSizesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ManagedUnsupportedVMSizesListCustomPager{},
		Path:       fmt.Sprintf("%s/managedUnsupportedVMSizes", id.ID()),
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
		Values *[]ManagedVMSize `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ManagedUnsupportedVMSizesListComplete retrieves all the results into a single object
func (c ManagedVMSizesClient) ManagedUnsupportedVMSizesListComplete(ctx context.Context, id LocationId) (ManagedUnsupportedVMSizesListCompleteResult, error) {
	return c.ManagedUnsupportedVMSizesListCompleteMatchingPredicate(ctx, id, ManagedVMSizeOperationPredicate{})
}

// ManagedUnsupportedVMSizesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedVMSizesClient) ManagedUnsupportedVMSizesListCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate ManagedVMSizeOperationPredicate) (result ManagedUnsupportedVMSizesListCompleteResult, err error) {
	items := make([]ManagedVMSize, 0)

	resp, err := c.ManagedUnsupportedVMSizesList(ctx, id)
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

	result = ManagedUnsupportedVMSizesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
