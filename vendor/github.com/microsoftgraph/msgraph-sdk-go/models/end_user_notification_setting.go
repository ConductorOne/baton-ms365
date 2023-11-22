package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// EndUserNotificationSetting 
type EndUserNotificationSetting struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewEndUserNotificationSetting instantiates a new endUserNotificationSetting and sets the default values.
func NewEndUserNotificationSetting()(*EndUserNotificationSetting) {
    m := &EndUserNotificationSetting{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEndUserNotificationSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEndUserNotificationSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.noTrainingNotificationSetting":
                        return NewNoTrainingNotificationSetting(), nil
                    case "#microsoft.graph.trainingNotificationSetting":
                        return NewTrainingNotificationSetting(), nil
                }
            }
        }
    }
    return NewEndUserNotificationSetting(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EndUserNotificationSetting) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *EndUserNotificationSetting) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EndUserNotificationSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["notificationPreference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEndUserNotificationPreference)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationPreference(val.(*EndUserNotificationPreference))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["positiveReinforcement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePositiveReinforcementNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPositiveReinforcement(val.(PositiveReinforcementNotificationable))
        }
        return nil
    }
    res["settingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEndUserNotificationSettingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingType(val.(*EndUserNotificationSettingType))
        }
        return nil
    }
    return res
}
// GetNotificationPreference gets the notificationPreference property value. Notification preference. Possible values are: unknown, microsoft, custom, unknownFutureValue.
func (m *EndUserNotificationSetting) GetNotificationPreference()(*EndUserNotificationPreference) {
    val, err := m.GetBackingStore().Get("notificationPreference")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EndUserNotificationPreference)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EndUserNotificationSetting) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPositiveReinforcement gets the positiveReinforcement property value. Positive reinforcement detail.
func (m *EndUserNotificationSetting) GetPositiveReinforcement()(PositiveReinforcementNotificationable) {
    val, err := m.GetBackingStore().Get("positiveReinforcement")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PositiveReinforcementNotificationable)
    }
    return nil
}
// GetSettingType gets the settingType property value. End user notification type. Possible values are: unknown, noTraining, trainingSelected, noNotification, unknownFutureValue.
func (m *EndUserNotificationSetting) GetSettingType()(*EndUserNotificationSettingType) {
    val, err := m.GetBackingStore().Get("settingType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EndUserNotificationSettingType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EndUserNotificationSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetNotificationPreference() != nil {
        cast := (*m.GetNotificationPreference()).String()
        err := writer.WriteStringValue("notificationPreference", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("positiveReinforcement", m.GetPositiveReinforcement())
        if err != nil {
            return err
        }
    }
    if m.GetSettingType() != nil {
        cast := (*m.GetSettingType()).String()
        err := writer.WriteStringValue("settingType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EndUserNotificationSetting) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *EndUserNotificationSetting) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetNotificationPreference sets the notificationPreference property value. Notification preference. Possible values are: unknown, microsoft, custom, unknownFutureValue.
func (m *EndUserNotificationSetting) SetNotificationPreference(value *EndUserNotificationPreference)() {
    err := m.GetBackingStore().Set("notificationPreference", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EndUserNotificationSetting) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPositiveReinforcement sets the positiveReinforcement property value. Positive reinforcement detail.
func (m *EndUserNotificationSetting) SetPositiveReinforcement(value PositiveReinforcementNotificationable)() {
    err := m.GetBackingStore().Set("positiveReinforcement", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingType sets the settingType property value. End user notification type. Possible values are: unknown, noTraining, trainingSelected, noNotification, unknownFutureValue.
func (m *EndUserNotificationSetting) SetSettingType(value *EndUserNotificationSettingType)() {
    err := m.GetBackingStore().Set("settingType", value)
    if err != nil {
        panic(err)
    }
}
// EndUserNotificationSettingable 
type EndUserNotificationSettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetNotificationPreference()(*EndUserNotificationPreference)
    GetOdataType()(*string)
    GetPositiveReinforcement()(PositiveReinforcementNotificationable)
    GetSettingType()(*EndUserNotificationSettingType)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetNotificationPreference(value *EndUserNotificationPreference)()
    SetOdataType(value *string)()
    SetPositiveReinforcement(value PositiveReinforcementNotificationable)()
    SetSettingType(value *EndUserNotificationSettingType)()
}
