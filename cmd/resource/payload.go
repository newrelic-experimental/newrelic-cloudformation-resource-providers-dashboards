package resource

import (
   "fmt"
   "github.com/newrelic/newrelic-cloudformation-resource-providers-common/model"
   log "github.com/sirupsen/logrus"
)

//
// Generic, should be able to leave these as-is
//

type Payload struct {
   model  *Model
   models []interface{}
}

func (p *Payload) SetIdentifier(g *string) {
   p.model.Guid = g
}

func (p *Payload) GetIdentifier() *string {
   return p.model.Guid
}

func (p *Payload) GetIdentifierKey(a model.Action) string {
   return "guid"
}

func (p *Payload) GetTagIdentifier() *string {
   return p.model.Guid
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

func (p *Payload) GetTags() map[string]string {
   return p.model.Tags
}

func (p *Payload) HasTags() bool {
   return p.model.Tags != nil
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
   return p.model.Dashboard
}

func (p *Payload) GetVariables() map[string]string {
   // FIXME Don't modify the original!
   vars := make(map[string]string)
   if p.model.Variables != nil {
      for k, v := range p.model.Variables {
         vars[k] = v
      }
   }

   if p.model.Guid != nil {
      vars["GUID"] = *p.model.Guid
   }

   if p.model.Dashboard != nil {
      vars["DASHBOARD"] = *p.model.Dashboard
   }

   lqf := ""
   if p.model.ListQueryFilter != nil {
      lqf = *p.model.ListQueryFilter
   }
   vars["LISTQUERYFILTER"] = lqf

   return vars
}

func (p *Payload) GetErrorKey() string {
   return "type"
}

func (p *Payload) GetCreateMutation() string {
   return `
mutation {
  dashboardCreate(accountId: {{{ACCOUNTID}}}, {{{DASHBOARD}}} ) {
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
  dashboardUpdate( {{{DASHBOARD}}} , guid: "{{{GUID}}}") {
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
      ... on DashboardEntity {
        domain
        entityType
        guid
        name
        type
      }
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
