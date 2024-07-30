package grouppolicyobjectfile

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

type ListGroupPolicyObjectFilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyObjectFile
}

type ListGroupPolicyObjectFilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyObjectFile
}

type ListGroupPolicyObjectFilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyObjectFilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyObjectFiles ...
func (c GroupPolicyObjectFileClient) ListGroupPolicyObjectFiles(ctx context.Context) (result ListGroupPolicyObjectFilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListGroupPolicyObjectFilesCustomPager{},
		Path:       "/deviceManagement/groupPolicyObjectFiles",
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
		Values *[]beta.GroupPolicyObjectFile `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupPolicyObjectFilesComplete retrieves all the results into a single object
func (c GroupPolicyObjectFileClient) ListGroupPolicyObjectFilesComplete(ctx context.Context) (ListGroupPolicyObjectFilesCompleteResult, error) {
	return c.ListGroupPolicyObjectFilesCompleteMatchingPredicate(ctx, GroupPolicyObjectFileOperationPredicate{})
}

// ListGroupPolicyObjectFilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyObjectFileClient) ListGroupPolicyObjectFilesCompleteMatchingPredicate(ctx context.Context, predicate GroupPolicyObjectFileOperationPredicate) (result ListGroupPolicyObjectFilesCompleteResult, err error) {
	items := make([]beta.GroupPolicyObjectFile, 0)

	resp, err := c.ListGroupPolicyObjectFiles(ctx)
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

	result = ListGroupPolicyObjectFilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
