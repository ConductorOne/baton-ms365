package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse 
type EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponse
}
// NewEntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse instantiates a new EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse and sets the default values.
func NewEntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse()(*EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse) {
    m := &EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse{
        BaseCollectionPaginationCountResponse: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateEntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse) GetValue()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponse) SetValue(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponseable 
type EntitlementManagementAssignmentsAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdGetResponseable interface {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable)
    SetValue(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable)()
}
