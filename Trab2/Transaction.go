package Transaction

type Transaction interface {
	
}

type TransactionImpl struct {
	Label     string //nome da transação
	TrId      int    //timestamp
	Status    string //"ativa", "concluida", "esperando" ou "abortada"
}

// Construtor nova transação
func newTransaction(label string, id int) *TransactionImpl {
	return &TransactionImpl{
		Label: label
		TrId: id,
		Status: "ativa"
	}
}

