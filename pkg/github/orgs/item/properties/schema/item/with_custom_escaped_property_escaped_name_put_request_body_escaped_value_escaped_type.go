package item
// The type of the value for the property
type WithCustom_property_namePutRequestBody_value_type int

const (
    STRING_WITHCUSTOM_PROPERTY_NAMEPUTREQUESTBODY_VALUE_TYPE WithCustom_property_namePutRequestBody_value_type = iota
    SINGLE_SELECT_WITHCUSTOM_PROPERTY_NAMEPUTREQUESTBODY_VALUE_TYPE
    MULTI_SELECT_WITHCUSTOM_PROPERTY_NAMEPUTREQUESTBODY_VALUE_TYPE
    TRUE_FALSE_WITHCUSTOM_PROPERTY_NAMEPUTREQUESTBODY_VALUE_TYPE
)

func (i WithCustom_property_namePutRequestBody_value_type) String() string {
    return []string{"string", "single_select", "multi_select", "true_false"}[i]
}
func ParseWithCustom_property_namePutRequestBody_value_type(v string) (any, error) {
    result := STRING_WITHCUSTOM_PROPERTY_NAMEPUTREQUESTBODY_VALUE_TYPE
    switch v {
        case "string":
            result = STRING_WITHCUSTOM_PROPERTY_NAMEPUTREQUESTBODY_VALUE_TYPE
        case "single_select":
            result = SINGLE_SELECT_WITHCUSTOM_PROPERTY_NAMEPUTREQUESTBODY_VALUE_TYPE
        case "multi_select":
            result = MULTI_SELECT_WITHCUSTOM_PROPERTY_NAMEPUTREQUESTBODY_VALUE_TYPE
        case "true_false":
            result = TRUE_FALSE_WITHCUSTOM_PROPERTY_NAMEPUTREQUESTBODY_VALUE_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWithCustom_property_namePutRequestBody_value_type(values []WithCustom_property_namePutRequestBody_value_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WithCustom_property_namePutRequestBody_value_type) isMultiValue() bool {
    return false
}
