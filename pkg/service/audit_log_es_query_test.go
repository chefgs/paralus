package service

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	v1 "github.com/paralus/paralus/proto/rpc/audit"
	v3 "github.com/paralus/paralus/proto/types/commonpb/v3"
)

type md struct {
	Source []string `json:"_source"`
	Aggs   struct {
		GroupByProject struct {
			Aggs struct {
				GroupByType struct {
					Terms struct {
						Field string `json:"field"`
						Size  int    `json:"size"`
					} `json:"terms"`
				} `json:"group_by_type"`
				GroupByUsername struct {
					Terms struct {
						Field string `json:"field"`
						Size  int    `json:"size"`
					} `json:"terms"`
				} `json:"group_by_username"`
			} `json:"aggs"`
			Terms struct {
				Field string `json:"field"`
				Size  int    `json:"size"`
			} `json:"terms"`
		} `json:"group_by_project"`
		GroupByType struct {
			Terms struct {
				Field string `json:"field"`
			} `json:"terms"`
		} `json:"group_by_type"`
		GroupByUsername struct {
			Terms struct {
				Field string `json:"field"`
			} `json:"terms"`
		} `json:"group_by_username"`
	} `json:"aggs"`
	Query struct {
		Bool struct {
			Filter struct {
				Range struct {
					JSONTimestamp struct {
						Gte string `json:"gte"`
						Lt  string `json:"lt"`
					} `json:"json.timestamp"`
				} `json:"range"`
			} `json:"filter"`
			Must []struct {
				Term struct {
					JSONCategory string `json:"json.category"`
				} `json:"term,omitempty"`
				Terms struct {
					JSONProject []string `json:"json.project"`
				} `json:"terms,omitempty"`
				QueryString struct {
					Query string `json:"query"`
				} `json:"query_string,omitempty"`
			} `json:"must"`
		} `json:"bool"`
	} `json:"query"`
	Size int `json:"size"`
	Sort struct {
		JSONTimestamp struct {
			Order string `json:"order"`
		} `json:"json.timestamp"`
	} `json:"sort"`
}

type mockElasticSearchQuery struct {
	msg []bytes.Buffer
}

func (m *mockElasticSearchQuery) Handle(msg bytes.Buffer) (map[string]interface{}, error) {
	m.msg = append(m.msg, msg)
	return map[string]interface{}{}, nil
}

func TestGetAuditLogByProjectsSimple(t *testing.T) {
	esq := &mockElasticSearchQuery{}
	al := &auditLogElasticSearchService{auditQuery: esq}
	req := v1.GetAuditLogSearchRequest{
		Filter: &v1.AuditLogQueryFilter{
			QueryString:   "query-string",
			Projects:      []string{"project-one", "project-two"},
			Timefrom:      "now-1h",
			Type:          "fake-type",
			User:          "fake-user",
			Client:        "fake-client",
			DashboardData: true,
		},
	}
	_, err := al.GetAuditLogByProjects(&req)
	if err != nil {
		t.Error("unable to get audit logs")
	}
	if len(esq.msg) != 1 {
		t.Fatalf("incorrect number of searches; expected '%v', got '%v'", 1, len(esq.msg))
	}
	m := &md{}
	err = json.Unmarshal(esq.msg[0].Bytes(), m)
	if err != nil {
		t.Fatal("unable to unmarshall es request")
	}
	expected := `{"_source":["json"],"aggs":{"group_by_project":{"aggs":{"group_by_type":{"terms":{"field":"json.type","size":1000}},"group_by_username":{"terms":{"field":"json.actor.account.username","size":1000}}},"terms":{"field":"json.project","size":1000}},"group_by_type":{"terms":{"field":"json.type"}},"group_by_username":{"terms":{"field":"json.actor.account.username"}}},"query":{"bool":{"filter":{"range":{"json.timestamp":{"gte":"now-1h","lt":"now"}}},"must":[{"term":{"json.category":"AUDIT"}},{"term":{"json.type":"fake-type"}},{"term":{"json.actor.account.username":"fake-user"}},{"term":{"json.client.type":"fake-client"}},{"terms":{"json.project":["project-one","project-two"]}},{"query_string":{"query":"query-string"}}]}},"size":0,"sort":{"json.timestamp":{"order":"desc"}}}`
	if strings.TrimSpace(esq.msg[0].String()) != expected {
		t.Errorf("incorrect es query; expected '%v', got '%v'", expected, strings.TrimSpace(esq.msg[0].String()))
	}
}

func TestGetAuditLogByProjectsNoProject(t *testing.T) {
	esq := &mockElasticSearchQuery{}
	al := &auditLogElasticSearchService{auditQuery: esq}
	req := v1.GetAuditLogSearchRequest{
		Metadata: &v3.Metadata{UrlScope: "url/project"},
		Filter: &v1.AuditLogQueryFilter{
			QueryString: "query-string",
		},
	}
	_, err := al.GetAuditLog(&req)
	if err != nil {
		t.Error("unable to get audit logs", err)
	}
	if len(esq.msg) != 1 {
		t.Fatalf("incorrect number of searches; expected '%v', got '%v'", 1, len(esq.msg))
	}
	m := &md{}
	err = json.Unmarshal(esq.msg[0].Bytes(), m)
	if err != nil {
		t.Fatal("unable to unmarshall es request")
	}

	expected := `{"_source":["json"],"aggs":{"group_by_type":{"terms":{"field":"json.type"}},"group_by_username":{"terms":{"field":"json.actor.account.username"}}},"query":{"bool":{"must":[{"term":{"json.category":"AUDIT"}},{"terms":{"json.project":["project"]}},{"query_string":{"query":"query-string"}}]}},"size":500,"sort":{"json.timestamp":{"order":"desc"}}}`
	if strings.TrimSpace(esq.msg[0].String()) != expected {
		t.Errorf("incorrect es query; expected '%v', got '%v'", expected, strings.TrimSpace(esq.msg[0].String()))
	}
}
