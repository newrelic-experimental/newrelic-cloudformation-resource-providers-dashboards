package resource

type Payload struct {
   model *Model
}

func NewPayload(m *Model) *Payload {
   return &Payload{model: m}
}

func (p *Payload) GetResourceModel() interface{} {
   return p.model
}

func (p *Payload) GetResourceModels() []interface{} {
   return []interface{}{p.model}
}

func (p *Payload) GetGraphQL() *string {
   return p.model.DashboardInput
}

func (p *Payload) SetGuid(g *string) {
   p.model.Guid = g
}

func (p *Payload) GetGuid() *string {
   return p.model.Guid
}

func (p *Payload) GetCreateMutation() string {
   return ""
}

func (p *Payload) GetDeleteMutation() string {
   return ""
}

func (p *Payload) GetUpdateMutation() string {
   return ""
}

func (p *Payload) GetReadQuery() string {
   return ""
}

func (p *Payload) GetListQuery() string {
   return ""
}

func (p *Payload) GetCreateResponse() interface{} {
   return ""
}

func (p *Payload) GetDeleteResponse() interface{} {
   return ""
}

func (p *Payload) GetUpdateResponse() interface{} {
   return ""
}

func (p *Payload) GetReadResponse() interface{} {
   return ""
}

func (p *Payload) GetListResponse() interface{} {
   return ""
}

func (p *Payload) GetListQueryFilter() *string {
   return p.model.ListQueryFilter
}
