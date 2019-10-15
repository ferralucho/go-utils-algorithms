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
   p.isExecuting.Lock() // Check is we can reserve it
   p.isExecuting.Unlock() // Because it may exist too many gets
   return p.responsePointer, p.error
}
