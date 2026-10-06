package analyzer
import "testing"

func TestAnalyzeProcess(t *testing.T){
	az := NewAnalyzer()
	log := &LogFields{
		Method:       "GET",
		StatusCode:   200,
		URL:          "/index.html",
		IpAddress:    "192.168.1.1",
		ResponseSize: 1024,
	}
	az.Process(log)
	if az.byMethod["GET"] != 1 {
		t.Errorf("Expected count of GET to be 1, got %d", az.byMethod["GET"])
	}
	if az.byCode[200] != 1 {
		t.Errorf("Expected count of 200 to be 1, got %d", az.byCode[200])
	}
	if az.byUrl["/index.html"] != 1 {
		t.Errorf("Expected count of /index.html to be 1, got %d", az.byUrl["/index.html"])
	}
	if az.byIP["192.168.1.1"] != 1 {
		t.Errorf("Expected count of 192.168.1.1 to be 1, got %d", az.byIP["192.168.1.1"])
	}
	if az.TotalResponseSize != 1024 {
		t.Errorf("Expected total response size to be 1024, got %d", az.TotalResponseSize)
	}
}