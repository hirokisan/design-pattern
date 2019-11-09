package hook_operation

// NOTE : 前後処理は外部から受け取るようにする

type Interface interface {
	Run() error
	pre() error
	post() error
}

type Object struct {
	preHooks  []func() error
	postHooks []func() error
}

func NewObject() *Object {
	return &Object{}
}

func (o *Object) WithPreHooks(hooks ...func() error) *Object {
	o.preHooks = hooks
	return o
}
func (o *Object) WithPostHooks(hooks ...func() error) *Object {
	o.postHooks = hooks
	return o
}

// Run : 実施前後に何らかの処理をする
func (o *Object) Run() error {
	if err := o.pre(); err != nil {
		return err
	}
	// do something
	if err := o.post(); err != nil {
		return err
	}
	return nil
}

func (o *Object) pre() error {
	for _, hook := range o.preHooks {
		if err := hook(); err != nil {
			return err
		}
	}
	return nil
}

func (o *Object) post() error {
	for _, hook := range o.postHooks {
		if err := hook(); err != nil {
			return err
		}
	}
	return nil
}
