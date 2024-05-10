package company

import "github.com/jae2274/goutils/enum"

type StatusValues struct{}

type Status enum.Enum[StatusValues]

const (
	Exist    = Status("exist")
	NotExist = Status("notExist")
	Unknown  = Status("unknown")
)

func (StatusValues) Values() []string {
	return []string{
		string(Exist),
		string(NotExist),
		string(Unknown),
	}
}
