package flyweight

var (
	instanceList = make(map[string]*Instance, 0)
)

type Instance struct {
	Key string
}

func newInstance(key string) *Instance {
	return &Instance{
		Key: key,
	}
}

func GetOrCreateInstance(key string) *Instance {
	if v, ok := instanceList[key]; ok {
		return v
	}

	i := newInstance(key)
	instanceList[key] = i

	return i
}
