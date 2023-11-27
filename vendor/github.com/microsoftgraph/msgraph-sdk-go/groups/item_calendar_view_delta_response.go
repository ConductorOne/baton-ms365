package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarViewDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarViewDeltaResponse struct {
    ItemCalendarViewDeltaGetResponse
}
// NewItemCalendarViewDeltaResponse instantiates a new ItemCalendarViewDeltaResponse and sets the default values.
func NewItemCalendarViewDeltaResponse()(*ItemCalendarViewDeltaResponse) {
    m := &ItemCalendarViewDeltaResponse{
        ItemCalendarViewDeltaGetResponse: *NewItemCalendarViewDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarViewDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarViewDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarViewDeltaResponse(), nil
}
// ItemCalendarViewDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarViewDeltaResponseable interface {
    ItemCalendarViewDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}