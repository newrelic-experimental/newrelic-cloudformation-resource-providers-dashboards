package resource

import (
   "fmt"
   "github.com/newrelic/newrelic-cloudformation-resource-providers-common/client/nerdgraph"
   "github.com/newrelic/newrelic-cloudformation-resource-providers-common/model"
   log "github.com/sirupsen/logrus"
   "strings"
)

// ResultHandler at a minimum provides access to the default error processing.
// If required we can provide custom processing here via composition overrides https://go.dev/doc/effective_go#embedding
type ResultHandler struct {
   // Use Go composition to access the default implementation
   model.ResultHandler
}

func NewResultHandler() (h model.ResultHandler) {
   defer func() {
      log.Debugf("(tagging) errorHandler.NewErrorHandler: exit %p", h)
   }()
   // Initialize ourself with the common core
   h = &ResultHandler{ResultHandler: nerdgraph.NewResultHandler()}
   return
}

func (h *ResultHandler) Delete(m model.Model, b []byte) (err error) {
   key := "status"
   var v interface{}
   v, err = nerdgraph.FindKeyValue(b, key)
   if err != nil {
      log.Errorf("Create: error finding result key: %s in response: %s", key, string(b))
      return err
   }

   status := fmt.Sprintf("%v", v)
   if !strings.EqualFold(status, "success") {
      err = fmt.Errorf("dashboard delete failed: %s", status)
   }
   return
}
