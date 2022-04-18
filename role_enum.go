package modelrole

type RoleEnum string

const (
	RoleUnauthorized RoleEnum = "unauthorized"
	RoleUser         RoleEnum = "user"
	RoleAdmin        RoleEnum = "admin"
	RoleSuper        RoleEnum = "super"
)

func (e RoleEnum) Level() int {
	levels := map[RoleEnum]int{
		RoleUnauthorized: 0,
		RoleUser:         1,
		RoleAdmin:        2,
		RoleSuper:        3,
	}

	return levels[e]
}

func (e *RoleEnum) Scan(value interface{}) error {
	*e = RoleEnum(value.(string))

	return nil
}

func (e RoleEnum) Value() (string, error) {
	return string(e), nil
}
