package dedicatedhosts

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

type ListAvailableSizesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]string
}

type ListAvailableSizesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []string
}

type ListAvailableSizesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListAvailableSizesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAvailableSizes ...
func (c DedicatedHostsClient) ListAvailableSizes(ctx context.Context, id commonids.DedicatedHostId) (result ListAvailableSizesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAvailableSizesCustomPager{},
		Path:       fmt.Sprintf("%s/hostSizes", id.ID()),
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
		Values *[]string `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAvailableSizesComplete retrieves all the results into a single object
func (c DedicatedHostsClient) ListAvailableSizesComplete(ctx context.Context, id commonids.DedicatedHostId) (result ListAvailableSizesCompleteResult, err error) {
	items := make([]string, 0)

	resp, err := c.ListAvailableSizes(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			items = append(items, v)
		}
	}

	result = ListAvailableSizesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
