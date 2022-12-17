package dict

// SimpleDict 不是线程安全的map
type SimpleDict struct {
	m map[string]interface{}
}

func MakeSimpleDict() *SimpleDict {
	return &SimpleDict{
		m: make(map[string]interface{}),
	}

}

func (dict *SimpleDict) Get(key string) (val interface{}, exists bool) {
	val, ok := dict.m[key]
	return val, ok
}

func (dict *SimpleDict) Len() int {
	if dict.m == nil {
		panic("m is nil")
	}
	return len(dict.m)
}

func (dict *SimpleDict) Put(key string, val interface{}) (result int) {
	_, existed := dict.m[key]
	dict.m[key] = val
	if existed {
		return 0
	}
	return 1
}

// PutIfAbsent puts value if the key is not exists and returns the number of updated key-value
func (dict *SimpleDict) PutIfAbsent(key string, val interface{}) (result int) {
	_, existed := dict.m[key]
	if existed {
		return 0
	}
	dict.m[key] = val
	return 1
}

func (dict *SimpleDict) PutIfExists(key string, val interface{}) (result int) {
	_, existed := dict.m[key]
	if existed {
		dict.m[key] = val
		return 1
	}
	return 0
}

func (dict *SimpleDict) Remove(key string) (result int) {
	_, existed := dict.m[key]
	delete(dict.m, key)
	if existed {
		return 1
	}
	return 0
}

func (dict *SimpleDict) ForEach(consumer consumer) {
	for k, v := range dict.m {
		if !consumer(k, v) {
			break
		}
	}
}

func (dict *SimpleDict) Keys() []string {
	result := make([]string, len(dict.m))
	i := 0
	for k := range dict.m {
		result[i] = k
		i++
	}
	return result
}

// RandomKeys randomly returns keys of the given number, may contain duplicated key
func (dict *SimpleDict) RandomKeys(limit int) []string {
	result := make([]string, limit)
	for i := 0; i < limit; i++ {
		for k := range dict.m {
			result[i] = k
			break
		}
	}
	return result
}

func (dict *SimpleDict) RandomDistinctKeys(limit int) []string {
	size := limit
	if size > len(dict.m) {
		size = len(dict.m)
	}
	result := make([]string, size)
	i := 0
	for k := range dict.m {
		if i == limit {
			break
		}
		result[i] = k
		i++
	}
	return result
}

func (dict *SimpleDict) Clear() {
	*dict = *MakeSimpleDict()
}
