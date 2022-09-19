package env

type Env map[string]string

func New(m map[string]string) Env {
	return Env(m)
}

func (e Env) Isset(k string) bool {
	for v := range e {
		if v == k {
			return true
		}
	}

	return false
}

func (e Env) Get(k string) string {
	return e[k]
}

func (e Env) Set(k string, v string) {
	e[k] = v
}

func (e Env) Delete(k string) {
	delete(e, k)
}
