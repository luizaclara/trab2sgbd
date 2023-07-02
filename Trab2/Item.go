package Item

type Item interface {
}

type ItemImpl struct {
	Label     string //nome do item
	BlockExcl string    //id transação que tem block escrita
	BlockComp []string  //id transações que tem block leitura
	WaitList  []string  //id transações aguardando concessão
}

func newItem(label string) *OperationImpl {
	return &OperationImpl{
		Label: label,
		BlockExcl: "empty",
		BlockComp []string{},
		WaitList  []string{},
	}
}
