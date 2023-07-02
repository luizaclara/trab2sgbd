package Block

type Block interface {
}

type BlockImpl struct {
	Item       string
	TrId       int
	Duration   string //"longa" ou  "curta"
	TypeBlock  int    //"R" ou "W"
	isAnalyzed int    // 1 ou 0
}

func newBlock(item string, idtr int, dur string, typeblock string) *BlockImpl {
	return &BlockImpl{
		Item:       item,
		TrId:       idtr,
		Duration:   dur,
		TypeBlock:  typeblock,
		isAnalyzed: 0,
	}
}
