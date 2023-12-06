package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// OnlineMeetingsItemRecordingsRequestBuilder provides operations to manage the recordings property of the microsoft.graph.onlineMeeting entity.
type OnlineMeetingsItemRecordingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlineMeetingsItemRecordingsRequestBuilderGetQueryParameters get recordings from communications
type OnlineMeetingsItemRecordingsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// OnlineMeetingsItemRecordingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingsItemRecordingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnlineMeetingsItemRecordingsRequestBuilderGetQueryParameters
}
// OnlineMeetingsItemRecordingsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingsItemRecordingsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCallRecordingId provides operations to manage the recordings property of the microsoft.graph.onlineMeeting entity.
func (m *OnlineMeetingsItemRecordingsRequestBuilder) ByCallRecordingId(callRecordingId string)(*OnlineMeetingsItemRecordingsCallRecordingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if callRecordingId != "" {
        urlTplParams["callRecording%2Did"] = callRecordingId
    }
    return NewOnlineMeetingsItemRecordingsCallRecordingItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewOnlineMeetingsItemRecordingsRequestBuilderInternal instantiates a new RecordingsRequestBuilder and sets the default values.
func NewOnlineMeetingsItemRecordingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingsItemRecordingsRequestBuilder) {
    m := &OnlineMeetingsItemRecordingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/onlineMeetings/{onlineMeeting%2Did}/recordings{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewOnlineMeetingsItemRecordingsRequestBuilder instantiates a new RecordingsRequestBuilder and sets the default values.
func NewOnlineMeetingsItemRecordingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingsItemRecordingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlineMeetingsItemRecordingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *OnlineMeetingsItemRecordingsRequestBuilder) Count()(*OnlineMeetingsItemRecordingsCountRequestBuilder) {
    return NewOnlineMeetingsItemRecordingsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get recordings from communications
func (m *OnlineMeetingsItemRecordingsRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlineMeetingsItemRecordingsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallRecordingCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCallRecordingCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallRecordingCollectionResponseable), nil
}
// Post create new navigation property to recordings for communications
func (m *OnlineMeetingsItemRecordingsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallRecordingable, requestConfiguration *OnlineMeetingsItemRecordingsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallRecordingable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCallRecordingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallRecordingable), nil
}
// ToGetRequestInformation get recordings from communications
func (m *OnlineMeetingsItemRecordingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlineMeetingsItemRecordingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to recordings for communications
func (m *OnlineMeetingsItemRecordingsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallRecordingable, requestConfiguration *OnlineMeetingsItemRecordingsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *OnlineMeetingsItemRecordingsRequestBuilder) WithUrl(rawUrl string)(*OnlineMeetingsItemRecordingsRequestBuilder) {
    return NewOnlineMeetingsItemRecordingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
