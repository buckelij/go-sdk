package repos
type GetTypeQueryParameterType int

const (
    ALL_GETTYPEQUERYPARAMETERTYPE GetTypeQueryParameterType = iota
    OWNER_GETTYPEQUERYPARAMETERTYPE
    MEMBER_GETTYPEQUERYPARAMETERTYPE
)

func (i GetTypeQueryParameterType) String() string {
    return []string{"all", "owner", "member"}[i]
}
func ParseGetTypeQueryParameterType(v string) (any, error) {
    result := ALL_GETTYPEQUERYPARAMETERTYPE
    switch v {
        case "all":
            result = ALL_GETTYPEQUERYPARAMETERTYPE
        case "owner":
            result = OWNER_GETTYPEQUERYPARAMETERTYPE
        case "member":
            result = MEMBER_GETTYPEQUERYPARAMETERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetTypeQueryParameterType(values []GetTypeQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetTypeQueryParameterType) isMultiValue() bool {
    return false
}
