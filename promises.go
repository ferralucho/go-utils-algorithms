package promises
 
import "sync"
 
type Promise struct {
   asyncFunction func() (interface{}, error)
   isExecuting sync.Mutex
   responsePointer interface{}
   error error
}
 
func New(asyncFunction func() (interface{}, error)) *Promise {
   promise := & Promise{asyncFunction: asyncFunction}
   promise.execute()
   return promise
}
 
func (p *Promise) executeFunctionAndUnlock() {
   p.responsePointer, p.error = p.asyncFunction()
   p.isExecuting.Unlock()
}
 
func (p *Promise) execute() {
   p.isExecuting.Lock()
   go p.executeFunctionAndUnlock()
}
 
func (p *Promise) Get() (interface{}, error) {
   p.isExecuting.Lock() // Chequeamos que se pueda reservar.
   p.isExecuting.Unlock() // Lo desbloqueamos porque podrían haber muchos Get's
   return p.responsePointer, p.error
}
