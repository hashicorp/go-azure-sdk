package grouppolicydefinitionfile

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

type ListGroupPolicyDefinitionFilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.GroupPolicyDefinitionFile
}

type ListGroupPolicyDefinitionFilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.GroupPolicyDefinitionFile
}

type ListGroupPolicyDefinitionFilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyDefinitionFilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyDefinitionFiles ...
func (c GroupPolicyDefinitionFileClient) ListGroupPolicyDefinitionFiles(ctx context.Context) (result ListGroupPolicyDefinitionFilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListGroupPolicyDefinitionFilesCustomPager{},
		Path:       "/deviceManagement/groupPolicyDefinitionFiles",
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
		Values *[]beta.GroupPolicyDefinitionFile `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupPolicyDefinitionFilesComplete retrieves all the results into a single object
func (c GroupPolicyDefinitionFileClient) ListGroupPolicyDefinitionFilesComplete(ctx context.Context) (ListGroupPolicyDefinitionFilesCompleteResult, error) {
	return c.ListGroupPolicyDefinitionFilesCompleteMatchingPredicate(ctx, GroupPolicyDefinitionFileOperationPredicate{})
}

// ListGroupPolicyDefinitionFilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyDefinitionFileClient) ListGroupPolicyDefinitionFilesCompleteMatchingPredicate(ctx context.Context, predicate GroupPolicyDefinitionFileOperationPredicate) (result ListGroupPolicyDefinitionFilesCompleteResult, err error) {
	items := make([]beta.GroupPolicyDefinitionFile, 0)

	resp, err := c.ListGroupPolicyDefinitionFiles(ctx)
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

	result = ListGroupPolicyDefinitionFilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
