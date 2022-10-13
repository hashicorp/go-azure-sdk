package hybrididentitymetadata

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByVmOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]HybridIdentityMetadata

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByVmOperationResponse, error)
}

type ListByVmCompleteResult struct {
	Items []HybridIdentityMetadata
}

func (r ListByVmOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByVmOperationResponse) LoadMore(ctx context.Context) (resp ListByVmOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByVm ...
func (c HybridIdentityMetadataClient) ListByVm(ctx context.Context, id VirtualMachineId) (resp ListByVmOperationResponse, err error) {
	req, err := c.preparerForListByVm(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybrididentitymetadata.HybridIdentityMetadataClient", "ListByVm", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybrididentitymetadata.HybridIdentityMetadataClient", "ListByVm", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByVm(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybrididentitymetadata.HybridIdentityMetadataClient", "ListByVm", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByVm prepares the ListByVm request.
func (c HybridIdentityMetadataClient) preparerForListByVm(ctx context.Context, id VirtualMachineId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/hybridIdentityMetadata", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByVmWithNextLink prepares the ListByVm request with the given nextLink token.
func (c HybridIdentityMetadataClient) preparerForListByVmWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByVm handles the response to the ListByVm request. The method always
// closes the http.Response Body.
func (c HybridIdentityMetadataClient) responderForListByVm(resp *http.Response) (result ListByVmOperationResponse, err error) {
	type page struct {
		Values   []HybridIdentityMetadata `json:"value"`
		NextLink *string                  `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByVmOperationResponse, err error) {
			req, err := c.preparerForListByVmWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "hybrididentitymetadata.HybridIdentityMetadataClient", "ListByVm", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "hybrididentitymetadata.HybridIdentityMetadataClient", "ListByVm", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByVm(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "hybrididentitymetadata.HybridIdentityMetadataClient", "ListByVm", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByVmComplete retrieves all of the results into a single object
func (c HybridIdentityMetadataClient) ListByVmComplete(ctx context.Context, id VirtualMachineId) (ListByVmCompleteResult, error) {
	return c.ListByVmCompleteMatchingPredicate(ctx, id, HybridIdentityMetadataOperationPredicate{})
}

// ListByVmCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c HybridIdentityMetadataClient) ListByVmCompleteMatchingPredicate(ctx context.Context, id VirtualMachineId, predicate HybridIdentityMetadataOperationPredicate) (resp ListByVmCompleteResult, err error) {
	items := make([]HybridIdentityMetadata, 0)

	page, err := c.ListByVm(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := ListByVmCompleteResult{
		Items: items,
	}
	return out, nil
}
