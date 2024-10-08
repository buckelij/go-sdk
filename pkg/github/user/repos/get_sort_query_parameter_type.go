package repos
type GetSortQueryParameterType int

const (
    CREATED_GETSORTQUERYPARAMETERTYPE GetSortQueryParameterType = iota
    UPDATED_GETSORTQUERYPARAMETERTYPE
    PUSHED_GETSORTQUERYPARAMETERTYPE
    FULL_NAME_GETSORTQUERYPARAMETERTYPE
)

func (i GetSortQueryParameterType) String() string {
    return []string{"created", "updated", "pushed", "full_name"}[i]
}
func ParseGetSortQueryParameterType(v string) (any, error) {
    result := CREATED_GETSORTQUERYPARAMETERTYPE
    switch v {
        case "created":
            result = CREATED_GETSORTQUERYPARAMETERTYPE
        case "updated":
            result = UPDATED_GETSORTQUERYPARAMETERTYPE
        case "pushed":
            result = PUSHED_GETSORTQUERYPARAMETERTYPE
        case "full_name":
            result = FULL_NAME_GETSORTQUERYPARAMETERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetSortQueryParameterType(values []GetSortQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetSortQueryParameterType) isMultiValue() bool {
    return false
}
