package scripts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptPackagesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ScriptPackage
}

type ScriptPackagesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ScriptPackage
}

type ScriptPackagesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ScriptPackagesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ScriptPackagesList ...
func (c ScriptsClient) ScriptPackagesList(ctx context.Context, id PrivateCloudId) (result ScriptPackagesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ScriptPackagesListCustomPager{},
		Path:       fmt.Sprintf("%s/scriptPackages", id.ID()),
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
		Values *[]ScriptPackage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ScriptPackagesListComplete retrieves all the results into a single object
func (c ScriptsClient) ScriptPackagesListComplete(ctx context.Context, id PrivateCloudId) (ScriptPackagesListCompleteResult, error) {
	return c.ScriptPackagesListCompleteMatchingPredicate(ctx, id, ScriptPackageOperationPredicate{})
}

// ScriptPackagesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ScriptsClient) ScriptPackagesListCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate ScriptPackageOperationPredicate) (result ScriptPackagesListCompleteResult, err error) {
	items := make([]ScriptPackage, 0)

	resp, err := c.ScriptPackagesList(ctx, id)
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

	result = ScriptPackagesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
