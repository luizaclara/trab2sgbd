package Scheduler

import (
	"Trab2/Operation"
	"Trab2/Block"
	"Trab2/Transaction"
	"fmt"
	"strconv"
)

type Scheduler interface {
	SchedulerSerializable()
}

type SchedulerImpl struct {
	ScheduleIn  string //escalonamento de entrada
	ScheduleOut string //escalonamento de saída
	OperationsIn  []*Operation.OperationImpl{} //operações separadas de entrada
	OperationsOut  []*Operation.OperationImpl{} //operações separadas de saída
}


//INICIALIZANDO ESCALONADOR
func newScheduler(schedule string) {
	operationsin := splitOperacoes(schedule)
	operationsout :=[]*Operation.OperationImpl{}

	return &SchedulerImpl{
		ScheduleIn:    schedule,
		ScheduleOut:   " ",
		OperationsIn:  operationsin,
		OperationsOut: operationsout,
	}
}

func splitOperacoes(escalonamento string) []Operation {
	estruturas := strings.Split(escalonamento, ")")

	operacoes := make([]Operation, 0)
	for _, estrutura := range estruturas {
		if estrutura != "" {
			op := parseOperacao(estrutura)
			operacoes = append(operacoes, op)
		}
	}

	return operacoes
}

func parseOperacao(estrutura string) Operation {
	parts := strings.SplitN(estrutura, "(", 2)
	fmt.Println(parts[0])

	tipo, trans, item := extractTransactionID(parts)

	operation := OperationImpl.newOperation(tipo, trans, item)
	
	return op
}

func extractTransactionID(transaction []string) (string, int, string) {
	firstPart := strings.TrimSpace(transaction[0])
	secondPart := strings.TrimSuffix(strings.TrimSpace(transaction[1]), ")")
	opTipo := string(firstPart[0])
	trans := 0
	item := ""

	if opTipo == "B" {
		opTipo += "T"
	} else if opTipo == "C" {
		trans, _ = strconv.Atoi(secondPart)
	} else {
		trans, _ = strconv.Atoi(firstPart[1:])
		item = secondPart
	}

	return opTipo, trans, item
}


//NÍVEIS DE ISOLAMENTO
func (sche *scheduler) SchedulerReadUncommitted() {

}

func (sche *scheduler) SchedulerReadCommitted() {

}

func (sche *scheduler) SchedulerRepeatableRead() {

}

//FALTA IMPLEMENTAR
//operação de uma transação que já está aguardando
//verificar bloqueios incompativeis
//wait die ---> colocar operação em espera + abort
//transformar operationsOut em schedule Out [!!!]

func (sche *SchedulerImpl) SchedulerSerializable() {
	listaTr := [] *Transaction //lista de transações
	listaBlock := [] *Block.BlockImpl //lista de bloqueios concedidos
	listaWait := [] *Block.BlockImpl //lista de bloqueios aguardando
	timeTr := 0 //contador transações (timestamp)

	//loop percorrendo lista de operações IN
	for indOp, op := range sche.OperationsIn {
		if (op.tipo == "BT"){			
			novaTr := Transaction.newTransaction(timeTr)			
			listaTr = append(listaTr, novaTr...)
			timeTr++
			fmt.Printf("A transacao %d foi ativada", op.idTrans)
			//transformar operationsOut em schedule Out [!!!]
			fmt.Printf("Schedule Atual: %d", sche.ScheduleOut)
			continue
		}

		if (listaTr[op.idTrans].Status == "esperando"){
			//guardar na lista de espera
			continue
		}

		else if (listaTr[op.idTrans].Status == "ativa"){
			//se op = R então
			if (op.Tipo = "R"){
				//gera bloqueio
				newblock := Block.newBlock(op.Item, op.idTrans, "longa", "R")
				
				//verifica conflito bloqueio
				flagConflict := 0 //flag para conflito de bloqueios
				timeConflit := 0
				for _, block := range listaBlock {
					// verifica se existe um bloqueio de escrita sobre o mesmo item
					if (block.Item == newblock.item && block.TypeBlock == "W" && block.TrId != newblock.TrId) {
						flagConflict = 1
						timeConflict = block.TrId
						break
					}
				}
				
				if (flagConflict == 0){
					//inserir na lista de bloqueios concedidos
					listaBlock = append(listaBlock, newblock)
					//escalona operação inserindo em OperationsOut
					sche.OperationsIn[indOp].Exec = 1
					sche.OperationsOut = append(sche.OperationsOut, op)
				}
				
				//WAIT-DIE quando tem conflito
				if (flagConflict == 1){
					if (op.timeStamp < timeConflict) {
						//op espera [!!!]
						// inserir na lista WaitItem
					} else {
						//liberar bloqueios da transação concedidos
						//desescalonar operações da transação
						//status atualizado para abortada
						listaTr[op.idTrans].Status = "abortada"
					}
				}
				
			} else if (op.Tipo = "W"){ 
				//gera bloqueio
				newblock := Block.newBlock(op.Item, op.idTrans, "longa", "W")
				
				conflit := 0 //flag para conflito de bloqueios
				//verifica na lista de bloqueio
				for _, block := range listaBlock {
					// verifica se existe um bloqueio sobre o mesmo item
					if (block.Item == newblock.item  && block.TrId != newblock.TrId) {
						conflit = 1 //se sim seta conflict = 1
						break
					}
				}

				if (conflict := 0){
					//inserir na lista de bloqueios concedidos
					listaBlock = append(listaBlock, newblock)
					//escalona operação inserindo em OperationsOut
					sche.OperationsIn[indOp].Exec = 1
					sche.OperationsOut = append(sche.OperationsOut, op)
				}
				
				if (conflict := 1){
					//wait die  [!!!]
				}
				
			} else if (op.Tipo = "C"){ 

				//libera bloqueios da transação
				markedBlocksToRemoveListaBlock := make([]bool, len(listaBlock))
				markedBlocksToRemoveListaWait := make([]bool, len(listaBlock))
				for indBlock, block := range listaBlock {
					if (block.TrId == op.idTrans){
						
						//marca para remover block da lista
						markedBlocksToRemoveListaBlock[indBlock] = true
						
						//busca bloqueios em espera para serem concedidos
						//0 = encontra nada
						//1 = encontrou leitura
						//se 0 encontrou escrita concede => break
						//se 0 encontrou leitura concede => 1
						//se 1 encontrou escrita  => break
						//se 1 encontrou leitura concede 

						flagWait := 0 //se encontrar bloqueio de escrita seta para 1 e para
						for indWait, wait := range listaWait {
							if (block.Item == wait.Item){
								//já concedeu bloqueios de leitura
								if (wait.TypeBlock == "W" && flagWait == 1){
									break
								}

								//primeiro bloqueio é de escrita
								else if (wait.TypeBlock == "W" && flagWait == 0){
									//concede o bloqueio
									//marca para remover wait da lista
									markedBlocksToRemoveListaBlock[indBlock] = true

									
									break
								}

								//primeiro bloqueio é de leitura
								else if (wait.TypeBlock == "R" && flagWait == 0){
									flagWait = 1
									//concede bloqueio
								}

								//continua concedendo bloqueios de leitura
								else if (wait.TypeBlock == "R" && flagWait == 1){
									//concede bloqueio
								}
								
								
								
								//marca para remover wait da lista
								markedBlocksToRemoveListaWait[indWait] = true
								//conceder bloqueio
								listaBlock = append(listaBlock, wait)
								
								//transação muda status para ativa
								listaTr[wait.TrId].Status = "ativa"
								
								//procura operações da transação que foi ativada
								for i := 0; i < indOp; i++{
									if (sche.OperationsIn[i].idTrans == wait.TrId && sche.OperationsIn[i].Exec == 0){
										//verificar se existe conflito
										for _, verificaBlock in listaBlock{
											retorno = .verificaConflito(op)
										}
										
										//se não houver => escalonar operação
										if (retorno==0){
											sche.OperationsIn[i].Exec = 1
											sche.OperationsOut = append(sche.OperationsOut, sche.OperationsIn[i])
										}
										
									}
								}

								
								
								

							}
						}
						//remover lista wait aqui
						listaWait = removeBlocks(listaWait, markedBlocksToRemoveListaWait)
						
						//verifica transações esperando por bloqueios liberados [!!!]
					}
				}
				//remover lista block aqui
				listaBlock = removeBlocks(listaBlock, markedBlocksToRemoveListaBlock)

				//transação é concluída
				listaTr[op.idTrans].Status = "concluida"

				//escalona commit em OperationsOut
				sche.OperationsIn[indOp].Exec = 1
				sche.OperationsOut = append(sche.OperationsOut, op)

			}				
		}
	}		
}


func removeBlocks (listaBlock [] *Block, markedBlocksToRemove []bool) [] *Block {
	novaListaBlock := [] *Block{}
	for i, block := range listaBlock {
		// Cria um novo slice apenas com os itens não marcados
		if !markedBlocksToRemove[i] {
			novaListaBlock = append(result, listaBlock[i])
		} 
	}

	return novaListaBlock
}







