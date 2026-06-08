package appattachpackageinfo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AppAttachPackage
}

type ImportCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AppAttachPackage
}

type ImportCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ImportCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// Import ...
func (c AppAttachPackageInfoClient) Import(ctx context.Context, id HostPoolId, input ImportPackageInfoRequest) (result ImportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &ImportCustomPager{},
		Path:       fmt.Sprintf("%s/importAppAttachPackageInfo", id.ID()),
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
		Values *[]AppAttachPackage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ImportComplete retrieves all the results into a single object
func (c AppAttachPackageInfoClient) ImportComplete(ctx context.Context, id HostPoolId, input ImportPackageInfoRequest) (ImportCompleteResult, error) {
	return c.ImportCompleteMatchingPredicate(ctx, id, input, AppAttachPackageOperationPredicate{})
}

// ImportCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppAttachPackageInfoClient) ImportCompleteMatchingPredicate(ctx context.Context, id HostPoolId, input ImportPackageInfoRequest, predicate AppAttachPackageOperationPredicate) (result ImportCompleteResult, err error) {
	items := make([]AppAttachPackage, 0)

	resp, err := c.Import(ctx, id, input)
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

	result = ImportCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
