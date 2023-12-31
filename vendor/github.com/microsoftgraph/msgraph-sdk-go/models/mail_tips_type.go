package models
import (
    "errors"
    "strings"
)
// 
type MailTipsType int

const (
    AUTOMATICREPLIES_MAILTIPSTYPE MailTipsType = iota
    MAILBOXFULLSTATUS_MAILTIPSTYPE
    CUSTOMMAILTIP_MAILTIPSTYPE
    EXTERNALMEMBERCOUNT_MAILTIPSTYPE
    TOTALMEMBERCOUNT_MAILTIPSTYPE
    MAXMESSAGESIZE_MAILTIPSTYPE
    DELIVERYRESTRICTION_MAILTIPSTYPE
    MODERATIONSTATUS_MAILTIPSTYPE
    RECIPIENTSCOPE_MAILTIPSTYPE
    RECIPIENTSUGGESTIONS_MAILTIPSTYPE
)

func (i MailTipsType) String() string {
    var values []string
    for p := MailTipsType(1); p <= RECIPIENTSUGGESTIONS_MAILTIPSTYPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"automaticReplies", "mailboxFullStatus", "customMailTip", "externalMemberCount", "totalMemberCount", "maxMessageSize", "deliveryRestriction", "moderationStatus", "recipientScope", "recipientSuggestions"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseMailTipsType(v string) (any, error) {
    var result MailTipsType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "automaticReplies":
                result |= AUTOMATICREPLIES_MAILTIPSTYPE
            case "mailboxFullStatus":
                result |= MAILBOXFULLSTATUS_MAILTIPSTYPE
            case "customMailTip":
                result |= CUSTOMMAILTIP_MAILTIPSTYPE
            case "externalMemberCount":
                result |= EXTERNALMEMBERCOUNT_MAILTIPSTYPE
            case "totalMemberCount":
                result |= TOTALMEMBERCOUNT_MAILTIPSTYPE
            case "maxMessageSize":
                result |= MAXMESSAGESIZE_MAILTIPSTYPE
            case "deliveryRestriction":
                result |= DELIVERYRESTRICTION_MAILTIPSTYPE
            case "moderationStatus":
                result |= MODERATIONSTATUS_MAILTIPSTYPE
            case "recipientScope":
                result |= RECIPIENTSCOPE_MAILTIPSTYPE
            case "recipientSuggestions":
                result |= RECIPIENTSUGGESTIONS_MAILTIPSTYPE
            default:
                return 0, errors.New("Unknown MailTipsType value: " + v)
        }
    }
    return &result, nil
}
func SerializeMailTipsType(values []MailTipsType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MailTipsType) isMultiValue() bool {
    return true
}
