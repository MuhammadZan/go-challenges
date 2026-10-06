package analyzer

import (
	"sort"
)

type RequestByMethod struct {
	Method string
	Count int
}

type RequestByCode struct {
	Code  int
	Count int
}

type RequestByUrl struct {
	Url string
	Count int
}

type RequestByIP struct {
	IP    string
	Count int
}

type LogFields struct {
	Method string
	IpAddress string
	TimeStamp string
	URL string
	StatusCode int
	ResponseSize int
	HTTPprotocol string
}

type Analysis struct {
	byMethod map[string]int
	byCode map[int]int
	byUrl map[string]int
	byIP map[string]int
	TotalResponseSize int
	LogsByMethod []RequestByMethod
	LogsByCode []RequestByCode
	LogsByUrl []RequestByUrl
	LogsByIP []RequestByIP
}

func NewAnalyzer() *Analysis {
	return &Analysis{
		byMethod: make(map[string]int),
		byCode:   make(map[int]int),
		byUrl:    make(map[string]int),
		byIP:     make(map[string]int),
		TotalResponseSize: 0,
		LogsByMethod: []RequestByMethod{},
		LogsByCode: []RequestByCode{},
		LogsByUrl: []RequestByUrl{},
		LogsByIP: []RequestByIP{},
	}
}

func (a *Analysis) Process(log *LogFields){
	a.byMethod[log.Method]++
	a.byCode[log.StatusCode]++
	a.byUrl[log.URL]++
	a.byIP[log.IpAddress]++
	a.TotalResponseSize += log.ResponseSize
}
func (a *Analysis) Sort(){
	for k,v := range a.byMethod{
		a.LogsByMethod = append(a.LogsByMethod, RequestByMethod{Method:k,Count: v})
	}
	for k,v := range a.byCode{
		a.LogsByCode = append(a.LogsByCode, RequestByCode{Code:k,Count: v})
	}
	for k,v := range a.byUrl{
		a.LogsByUrl = append(a.LogsByUrl, RequestByUrl{Url:k,Count: v})
	}
	for k,v := range a.byIP{
		a.LogsByIP = append(a.LogsByIP, RequestByIP{IP:k,Count: v})
	}
	sort.Slice(a.LogsByMethod, func(i, j int) bool {
		return a.LogsByMethod[i].Count > a.LogsByMethod[j].Count
	})
	sort.Slice(a.LogsByCode, func(i, j int) bool {
		return a.LogsByCode[i].Count > a.LogsByCode[j].Count
	})
	sort.Slice(a.LogsByUrl, func(i, j int) bool {
		return a.LogsByUrl[i].Count > a.LogsByUrl[j].Count
	})
	sort.Slice(a.LogsByIP, func(i, j int) bool {
		return a.LogsByIP[i].Count > a.LogsByIP[j].Count
	})
}
