package issues
type GetSortQueryParameterType int

const (
    CREATED_GETSORTQUERYPARAMETERTYPE GetSortQueryParameterType = iota
    UPDATED_GETSORTQUERYPARAMETERTYPE
    COMMENTS_GETSORTQUERYPARAMETERTYPE
)

func (i GetSortQueryParameterType) String() string {
    return []string{"created", "updated", "comments"}[i]
}
func ParseGetSortQueryParameterType(v string) (any, error) {
    result := CREATED_GETSORTQUERYPARAMETERTYPE
    switch v {
        case "created":
            result = CREATED_GETSORTQUERYPARAMETERTYPE
        case "updated":
            result = UPDATED_GETSORTQUERYPARAMETERTYPE
        case "comments":
            result = COMMENTS_GETSORTQUERYPARAMETERTYPE
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
