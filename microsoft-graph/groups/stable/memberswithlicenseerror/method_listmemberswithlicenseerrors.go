package memberswithlicenseerror

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMembersWithLicenseErrorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListMembersWithLicenseErrorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListMembersWithLicenseErrorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMembersWithLicenseErrorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMembersWithLicenseErrors ...
func (c MembersWithLicenseErrorClient) ListMembersWithLicenseErrors(ctx context.Context, id GroupId) (result ListMembersWithLicenseErrorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMembersWithLicenseErrorsCustomPager{},
		Path:       fmt.Sprintf("%s/membersWithLicenseErrors", id.ID()),
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
		Values *[]stable.DirectoryObject `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMembersWithLicenseErrorsComplete retrieves all the results into a single object
func (c MembersWithLicenseErrorClient) ListMembersWithLicenseErrorsComplete(ctx context.Context, id GroupId) (ListMembersWithLicenseErrorsCompleteResult, error) {
	return c.ListMembersWithLicenseErrorsCompleteMatchingPredicate(ctx, id, DirectoryObjectOperationPredicate{})
}

// ListMembersWithLicenseErrorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MembersWithLicenseErrorClient) ListMembersWithLicenseErrorsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate DirectoryObjectOperationPredicate) (result ListMembersWithLicenseErrorsCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListMembersWithLicenseErrors(ctx, id)
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

	result = ListMembersWithLicenseErrorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
