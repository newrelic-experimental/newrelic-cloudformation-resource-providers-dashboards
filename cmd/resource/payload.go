package resource

import (
   "fmt"
   "github.com/newrelic-experimental/newrelic-cloudformation-resource-providers-common/model"
   log "github.com/sirupsen/logrus"
)

//
// Generic, should be able to leave these as-is
//

type Payload struct {
   model  *Model
   models []interface{}
}

func NewPayload(m *Model) *Payload {
   return &Payload{
      model:  m,
      models: make([]interface{}, 0),
   }
}

func (p *Payload) GetResourceModel() interface{} {
   return p.model
}

func (p *Payload) GetResourceModels() []interface{} {
   log.Debugf("GetResourceModels: returning %+v", p.models)
   return p.models
}

func (p *Payload) AppendToResourceModels(m model.Model) {
   p.models = append(p.models, m.GetResourceModel())
}

//
// These are API specific, must be configured per API
//

var typeName = "NewRelic::Observability::Dashboards"

func (p *Payload) NewModelFromGuid(g interface{}) (m model.Model) {
   s := fmt.Sprintf("%s", g)
   return NewPayload(&Model{Guid: &s})
}

func (p *Payload) GetGraphQLFragment() *string {
   return p.model.DashboardInput
}

func (p *Payload) SetGuid(g *string) {
   p.model.Guid = g
   log.Debugf("SetGuid: %s", *p.model.Guid)
}

func (p *Payload) GetGuid() *string {
   return p.model.Guid
}

func (p *Payload) GetCreateMutation() string {
   return `
mutation {
  dashboardCreate(accountId: {{{ACCOUNTID}}}, dashboard: { {{{DASHBOARD}}} }) {
    entityResult {
      guid
    }
    errors {
      description
      type
    }
  }
}
`
}

func (p *Payload) GetDeleteMutation() string {
   return `
mutation {
  dashboardDelete(guid: "{{{GUID}}}") {
    errors {
      description
      type
    }
    status
  }
}
`
}

func (p *Payload) GetUpdateMutation() string {
   return `
mutation {
  dashboardUpdate(dashboard: { {{{DASHBOARD}}} }, guid: "{{{GUID}}}") {
    entityResult {
      guid
    }
    errors {
      description
      type
    }
  }
}
`
}

func (p *Payload) GetReadQuery() string {
   return `
{
  actor {
    entity(guid: "{{{GUID}}}") {
      domain
      entityType
      guid
      name
      type
    }
  }
}
`
}

func (p *Payload) GetListQuery() string {
   return `
{
  actor {
    entitySearch(query: "accountId = '{{{ACCOUNTID}}}' AND type = 'DASHBOARD' {{{LISTQUERYFILTERS}}}") {
      count
      results {
        entities {
          guid
        }
        nextCursor
      }
    }
  }
}
`
}

func (p *Payload) GetListQueryNextCursor() string {
   return `
{
  actor {
    entitySearch(query: "accountId = '{{{ACCOUNTID}}}' AND type = 'DASHBOARD' {{{LISTQUERYFILTER}}}") {
      count
      results(cursor: "{{{NEXTCURSOR}}}") {
        entities {
          guid
        }
        nextCursor
      }
    }
  }
}
`
}

// func (p *Payload) GetListQueryFilter() *string {
//    return p.model.ListQueryFilter
// }

func (p *Payload) GetGuidKey() string {
   return "guid"
}

func (p *Payload) GetVariables() map[string]string {
   // ACCOUNTID comes from the configuration
   // NEXTCURSOR is a _convention_

   if p.model.Variables == nil {
      p.model.Variables = make(map[string]string)
   }

   if p.model.Guid != nil {
      p.model.Variables["GUID"] = *p.model.Guid
   }

   if p.model.DashboardInput != nil {
      p.model.Variables["DASHBOARD"] = *p.model.DashboardInput
   }

   lqf := ""
   if p.model.ListQueryFilter != nil {
      lqf = *p.model.ListQueryFilter
   }
   p.model.Variables["LISTQUERYFILTER"] = lqf

   return p.model.Variables
}

func (p *Payload) GetErrorKey() string {
   return "type"
}

func (p *Payload) GetResultKey(a model.Action) string {
   switch a {
   case model.Delete:
      return ""
   default:
      return p.GetGuidKey()
   }
}

func (p *Payload) NeedsPropagationDelay(a model.Action) bool {
   return true
}
