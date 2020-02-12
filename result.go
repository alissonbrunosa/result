package result

type Result interface {
	Value() interface{}
	IsOkay() bool
	IsError() bool
	Or(Result) Result
	And(Result) Result
	AndThen(res) Result
	OrElse(res) Result
}

type res func(interface{}) Result

type okay struct {
	value interface{}
}

func (o *okay) Value() interface{}       { return o.value }
func (o *okay) IsOkay() bool             { return true }
func (o *okay) IsError() bool            { return !o.IsOkay() }
func (o *okay) Or(result Result) Result  { return o }
func (o *okay) And(result Result) Result { return result }
func (o *okay) AndThen(fn res) Result    { return fn(o.value) }
func (o *okay) OrElse(fn res) Result     { return o }

type err struct {
	value interface{}
}

func (e *err) Value() interface{}       { return e.value }
func (e *err) IsOkay() bool             { return !e.IsError() }
func (e *err) IsError() bool            { return true }
func (e *err) Or(result Result) Result  { return result }
func (e *err) And(result Result) Result { return e }
func (e *err) AndThen(fn res) Result    { return e }
func (e *err) OrElse(fn res) Result     { return fn(e.value) }

func Ok(value interface{}) Result {
	return &okay{value: value}
}

func Err(value interface{}) Result {
	return &err{value: value}
}
