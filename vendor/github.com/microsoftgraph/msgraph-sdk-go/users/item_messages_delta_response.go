package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMessagesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemMessagesDeltaResponse struct {
    ItemMessagesDeltaGetResponse
}
// NewItemMessagesDeltaResponse instantiates a new ItemMessagesDeltaResponse and sets the default values.
func NewItemMessagesDeltaResponse()(*ItemMessagesDeltaResponse) {
    m := &ItemMessagesDeltaResponse{
        ItemMessagesDeltaGetResponse: *NewItemMessagesDeltaGetResponse(),
    }
    return m
}
// CreateItemMessagesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMessagesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMessagesDeltaResponse(), nil
}
// ItemMessagesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemMessagesDeltaResponseable interface {
    ItemMessagesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
