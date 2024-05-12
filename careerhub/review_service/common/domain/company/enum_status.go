package company

import "github.com/jae2274/goutils/enum"

type StatusValues struct{}

type ExistStatus enum.Enum[StatusValues]

const (
	Exist    = ExistStatus("exist")
	NotExist = ExistStatus("notExist")
	Unknown  = ExistStatus("unknown")
)

func (StatusValues) Values() []string {
	return []string{
		string(Exist),
		string(NotExist),
		string(Unknown),
	}
}

type CrawlingStatusValues struct{}

type CrawlingStatus enum.Enum[CrawlingStatusValues]

const (
	NotCrawled = CrawlingStatus("notCrawled")
	Crawled    = CrawlingStatus("crawled")
)

func (CrawlingStatusValues) Values() []string {
	return []string{
		string(NotCrawled),
		string(Crawled),
	}
}
