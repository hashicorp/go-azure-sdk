package grouppolicymigrationreportunsupportedgrouppolicyextension

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

type ListGroupPolicyMigrationReportUnsupportedGroupPolicyExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnsupportedGroupPolicyExtension
}

type ListGroupPolicyMigrationReportUnsupportedGroupPolicyExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnsupportedGroupPolicyExtension
}

type ListGroupPolicyMigrationReportUnsupportedGroupPolicyExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGroupPolicyMigrationReportUnsupportedGroupPolicyExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGroupPolicyMigrationReportUnsupportedGroupPolicyExtensions ...
func (c GroupPolicyMigrationReportUnsupportedGroupPolicyExtensionClient) ListGroupPolicyMigrationReportUnsupportedGroupPolicyExtensions(ctx context.Context, id DeviceManagementGroupPolicyMigrationReportId) (result ListGroupPolicyMigrationReportUnsupportedGroupPolicyExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListGroupPolicyMigrationReportUnsupportedGroupPolicyExtensionsCustomPager{},
		Path:       fmt.Sprintf("%s/unsupportedGroupPolicyExtensions", id.ID()),
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
		Values *[]beta.UnsupportedGroupPolicyExtension `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupPolicyMigrationReportUnsupportedGroupPolicyExtensionsComplete retrieves all the results into a single object
func (c GroupPolicyMigrationReportUnsupportedGroupPolicyExtensionClient) ListGroupPolicyMigrationReportUnsupportedGroupPolicyExtensionsComplete(ctx context.Context, id DeviceManagementGroupPolicyMigrationReportId) (ListGroupPolicyMigrationReportUnsupportedGroupPolicyExtensionsCompleteResult, error) {
	return c.ListGroupPolicyMigrationReportUnsupportedGroupPolicyExtensionsCompleteMatchingPredicate(ctx, id, UnsupportedGroupPolicyExtensionOperationPredicate{})
}

// ListGroupPolicyMigrationReportUnsupportedGroupPolicyExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupPolicyMigrationReportUnsupportedGroupPolicyExtensionClient) ListGroupPolicyMigrationReportUnsupportedGroupPolicyExtensionsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementGroupPolicyMigrationReportId, predicate UnsupportedGroupPolicyExtensionOperationPredicate) (result ListGroupPolicyMigrationReportUnsupportedGroupPolicyExtensionsCompleteResult, err error) {
	items := make([]beta.UnsupportedGroupPolicyExtension, 0)

	resp, err := c.ListGroupPolicyMigrationReportUnsupportedGroupPolicyExtensions(ctx, id)
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

	result = ListGroupPolicyMigrationReportUnsupportedGroupPolicyExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
