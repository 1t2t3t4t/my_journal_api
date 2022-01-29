package utils

func ContainKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

func FindMapValue[K comparable, V any](m map[K]V, finder func(V) bool) (K, V, bool) {
	for k, v := range m {
		if finder(v) {
			return k, v, true
		}
	}

	return *new(K), *new(V), false
}
