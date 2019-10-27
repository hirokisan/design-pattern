package flyweight

var (
	instanceList = make(map[string]*instance, 0)
)

type instance struct {
	Key string
}

func (i *instance) Access() string {
	return "accessible"
}

func newInstance(key string) *instance {
	return &instance{
		Key: key,
	}
}

func GetOrCreateInstance(key string) *instance {
	if v, ok := instanceList[key]; ok {
		return v
	}

	i := newInstance(key)
	instanceList[key] = i

	return i
}
