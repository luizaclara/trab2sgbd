package Operation

type Operation interface {
}

type OperationImpl struct {
	Tipo string // Tipo da operação: BT, r, w, C
	Item string // Item da operação (x, y, z)
	TrId int    // Identificador da transação
	Exec int    //Flag para indicar se já foi executada
}

func newOperation(tipo string, idTrans int, item string) *OperationImpl {
	return &OperationImpl{
		Tipo: tipo,
		Item: item,
		TrId: idTrans,
		Exec: 0,
	}
}
