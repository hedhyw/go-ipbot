package interfaces

type IFConfig interface {
	GetIP() (string, error)
}
