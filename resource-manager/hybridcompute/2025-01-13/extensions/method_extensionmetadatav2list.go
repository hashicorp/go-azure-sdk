package extensions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtensionMetadataV2ListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ExtensionValueV2
}

type ExtensionMetadataV2ListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ExtensionValueV2
}

type ExtensionMetadataV2ListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ExtensionMetadataV2ListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ExtensionMetadataV2List ...
func (c ExtensionsClient) ExtensionMetadataV2List(ctx context.Context, id ExtensionTypeId) (result ExtensionMetadataV2ListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ExtensionMetadataV2ListCustomPager{},
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
		Values *[]ExtensionValueV2 `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ExtensionMetadataV2ListComplete retrieves all the results into a single object
func (c ExtensionsClient) ExtensionMetadataV2ListComplete(ctx context.Context, id ExtensionTypeId) (ExtensionMetadataV2ListCompleteResult, error) {
	return c.ExtensionMetadataV2ListCompleteMatchingPredicate(ctx, id, ExtensionValueV2OperationPredicate{})
}

// ExtensionMetadataV2ListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExtensionsClient) ExtensionMetadataV2ListCompleteMatchingPredicate(ctx context.Context, id ExtensionTypeId, predicate ExtensionValueV2OperationPredicate) (result ExtensionMetadataV2ListCompleteResult, err error) {
	items := make([]ExtensionValueV2, 0)

	resp, err := c.ExtensionMetadataV2List(ctx, id)
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

	result = ExtensionMetadataV2ListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
