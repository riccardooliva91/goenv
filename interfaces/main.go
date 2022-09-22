package interfaces

type EnvInterface interface {
	Isset(k string) bool
	Get(k string) string
	Set(k string, v string)
	Delete(k string)
}
