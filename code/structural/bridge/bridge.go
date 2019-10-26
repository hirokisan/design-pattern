package bridge

// Bridge : interfaceにmethodを定義してBridgeにセットする
type Bridge struct {
	MethodA MethodInterfaceA
	MethodB MethodInterfaceB
}

// NewBridge : この辺はfactoryパターンを用いたりすると柔軟性を持たせられるはず
func NewBridge(methodA MethodInterfaceA, methodB MethodInterfaceB) *Bridge {
	return &Bridge{
		MethodA: methodA,
		MethodB: methodB,
	}
}

func (b *Bridge) RunMethodA() string {
	return b.MethodA.Run()
}

func (b *Bridge) RunMethodB() string {
	return b.MethodB.Run()
}

type MethodInterfaceA interface {
	Run() string
}

type MethodInterfaceB interface {
	Run() string
}

// MethodInplA : MethodInterfaceAを満たす実装
type MethodInplA struct{}

func NewMethodImplA() MethodInterfaceA {
	return &MethodInplA{}
}

func (i *MethodInplA) Run() string {
	return "this is MethodA"
}

// MethodInplB : MethodInterfaceBを満たす実装
type MethodInplB struct{}

func NewMethodImplB() MethodInterfaceB {
	return &MethodInplB{}
}

func (i *MethodInplB) Run() string {
	return "this is MethodB"
}
