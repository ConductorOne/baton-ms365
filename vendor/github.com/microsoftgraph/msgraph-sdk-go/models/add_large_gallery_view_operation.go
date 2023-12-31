package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddLargeGalleryViewOperation 
type AddLargeGalleryViewOperation struct {
    CommsOperation
}
// NewAddLargeGalleryViewOperation instantiates a new addLargeGalleryViewOperation and sets the default values.
func NewAddLargeGalleryViewOperation()(*AddLargeGalleryViewOperation) {
    m := &AddLargeGalleryViewOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// CreateAddLargeGalleryViewOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAddLargeGalleryViewOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddLargeGalleryViewOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddLargeGalleryViewOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AddLargeGalleryViewOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// AddLargeGalleryViewOperationable 
type AddLargeGalleryViewOperationable interface {
    CommsOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
