package Scheduler

import (
	"Trab2/Operation"
	"Trab2/Block"
	"Trab2/Transaction"
	"Trab2/Item"
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
	ListItens [] *Item.ItemImpl{} // Itens do escalonamento
}


//INICIALIZANDO ESCALONADOR
func newScheduler(schedule string) {
	operationsin, listItens:= splitOperacoes(schedule)
	operationsout :=[]*Operation.OperationImpl{}

	return &SchedulerImpl{
		ScheduleIn:    schedule,
		ScheduleOut:   " ",
		OperationsIn:  operationsin,
		OperationsOut: operationsout,
		ListItens: listItens
	}
}

func splitOperacoes(escalonamento string) ([]*Operation.OperationImpl, []string) {
	estruturas := strings.Split(escalonamento, ")")
	listItens := []string{}
	seenItem := make(map[string]bool)

	operacoes := make([]Operation, 0)
	for _, estrutura := range estruturas {
		if estrutura != "" {
			op := parseOperacao(estrutura)
			operacoes = append(operacoes, op)
		}
	}

	for _, item := range operacoes {
		if !seenItem[item] {
			seenItem[item] = true
			listItens = append(listItens, item)
		}
	}

	return operacoes, listItens
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
	listaTrans := [] *Transaction.TransactionImpl{} //lista de transações
	timeTr := 0 //contador transações (timestamp)
	TrIdStr := strconv.Itoa(op.TrId)
	//loop percorrendo lista de operações IN
	for indOp, op := range sche.OperationsIn {
		if (op.tipo == "BT"){			
			novaTr := Transaction.newTransaction(timeTr)			
			listaTr = append(listaTr, novaTr...)
			timeTr++
			fmt.Printf("A transacao %d foi ativada", op.TrId)
			//transformar operationsOut em schedule Out [!!!]
			fmt.Printf("Schedule Atual: %d", sche.ScheduleOut)
			continue
		}

		if (listaTr[op.TrId].Status == "esperando"){
			continue
		}

		else if (listaTr[op.TrId].Status == "ativa"){
			//gera bloqueio
			//newblock := Block.newBlock(op.Item, op.TrId, "longa", "R")
			if (op.Tipo = "R"){
				for _, item in sche.ListItens{
					//item não tem bloqueio exclusivo (W) => leitura é concedida
					if (item.Label == op.Item && item.BlockExcl == "empty"){
						item.BlockComp = append(item.BlockComp, TrIdStr) //insere em block compartilhados
						sche.OperationsIn[indOp].Exec = 1 //flag executada para 1
						sche.OperationsOut = append(sche.OperationsOut, sche.OperationsIn[indOp]) //lista de escalonamento
					}
					
					//transação já tem bloqueio exclusivo do item, mas quer fazer leitura
					else if (item.Label == op.Item && item.BlockExcl == TrIdStr){
						sche.OperationsIn[indOp].Exec = 1 //flag executada para 1
						sche.OperationsOut = append(sche.OperationsOut, sche.OperationsIn[indOp]) //lista de escalonamento
					}

					//item tem bloqueio exclusivo de transação distinta => fazemos WAIT DIE
					else if (item.Label == op.Item && item.BlockExcl != "empty"){
						TStrblock := strconv.Atoi(item.BlockExcl)
						//WAIT
						if (op.TrId < TStrblock){
							item.WaitList = append(item.WaitList, TrIdStr) //insere na lista de espera
							listaTr[op.TrId].Status = "esperando"
						}
						//DIE
						else{
							//colocar flag execução para 0
							for _, abort in sche.OperationsIn{
								if (abort.TrId == op.TrId){
									abort.Exec = 0
								}
							}
							//remover operações de OperationsOut [!!!]

							//atualizar status transação
							listaTr[op.TrId].Status = "abortada"
						}
						
					}
				}
			} else if (op.Tipo = "W"){
				//gera bloqueio
				//newblock := Block.newBlock(op.Item, op.TrId, "longa", "W")
				for _, item in sche.ListItens{
					
					//item não tem bloqueio exclusivo (W) nem compartilhado (R) => escrita é concedida
					if (item.Label == op.Item && item.BlockExcl == "empty" && len(item.BlockComp)==0){
						item.BlockExcl = append(item.BlockExcl, TrIdStr) //insere em block exclusivo
						sche.OperationsIn[indOp].Exec = 1 //flag executada para 1
						sche.OperationsOut = append(sche.OperationsOut, sche.OperationsIn[indOp]) //lista de escalonamento
					}
					
					//item não tem bloqueio exclusivo(W) mas tem bloqueio compartilhado(R)
					else if (item.Label == op.Item && item.BlockExcl == "empty" && len(item.BlockComp)>0){
						convert := 0
						//transação já tem bloqueio de leitura e pode converter em de escrita
						for _,comp in BlockComp{
							if (TrIdStr==comp){
								convert = 1
								//remover comp de BlockComp [!!!]
								item.BlockExcl = TrIdStr
								sche.OperationsIn[indOp].Exec = 1 //flag executada para 1
								sche.OperationsOut = append(sche.OperationsOut, sche.OperationsIn[indOp]) //lista de escalonamento
								break
							}
						}

						//já existe bloqueio de leitura (incompatível com de escrita)
						if (!convert){
							TStrblock := strconv.Atoi(item.BlockExcl)
							//WAIT
							if (op.TrId < TStrblock){
								item.WaitList = append(item.WaitList, TrIdStr) //insere na lista de espera
								listaTr[op.TrId].Status = "esperando"
							}
							//DIE
							else{
								//colocar flag execução para 0
								for _, abort in sche.OperationsIn{
									if (abort.TrId == op.TrId){
										abort.Exec = 0
									}
								}
								//remover operações de OperationsOut [!!!]

								//atualizar status transação
								listaTr[op.TrId].Status = "abortada"
							}
						}
					}

					//outra transação já tem bloqueio exclusivo do item => fazemos WAIT DIE
					else if (item.Label == op.Item && item.BlockExcl != "empty" && item.BlockExcl != TrIdStr){
						TStrblock := strconv.Atoi(item.BlockExcl)
						//WAIT
						if (op.TrId < TStrblock){
							item.WaitList = append(item.WaitList, TrIdStr) //insere na lista de espera
							listaTr[op.TrId].Status = "esperando"
						}
						//DIE
						else{
							//colocar flag execução para 0
							for _, abort in sche.OperationsIn{
								if (abort.TrId == op.TrId){
									abort.Exec = 0
								}
							}
							//remover operações de OperationsOut [!!!]

							//atualizar status transação
							listaTr[op.TrId].Status = "abortada"
						}
					}

					//item tem bloqueio exclusivo da mesma transação 
					else if (item.Label == op.Item && item.BlockExcl == TrIdStr){
						sche.OperationsIn[indOp].Exec = 1 //flag executada para 1
						sche.OperationsOut = append(sche.OperationsOut, sche.OperationsIn[indOp]) //lista de escalonamento
					}
				}
				
				
			} else if (op.Tipo = "C"){	 
				for _, item in sche.ListItens{
					if (item.BlockExcl == TrIdStr){
						BlockExcl = "empty"
					}

					for _,comp in item.BlockComp{
						if (comp == TrIdStr){
							//remover comp
						}
					}

					//verificar wait list
					//primeira op de primeiro wait OK
					//verificar demais op de primeiro wait e para quando conflitar ou concluir
					//apartir do segundo wait  se algum conflitar na primeira op para verifica wait
					parouWait := 0
					//percorrendo lista FIFO wait list
					for indWaitLs := 0; i < len(item.WaitList); indWaitLs++{
						TrIdWait := strconv.Atoi(item.WaitList[indWaitLs])
						beginWait := 0
						//percorrendo lista de operações travadas da transação que acabou de ser liberada
						for indWaitTr := 0; i < indOp; indWaitTr++ {
							if (sche.OperationsIn[indWaitTr].TrId == TrIdWait && sche.OperationsIn[indWaitTr].Exec == 0 && beginWait == 0){		
								//verificar conflitos
								//se nao conflito
									//escalonar
									//beginWait = 1
								//se conflito 
									//wait die
									parouWait =1
									break
							}

							else if (sche.OperationsIn[indWaitTr].TrId == TrIdWait && sche.OperationsIn[indWaitTr].Exec == 0 && beginWait == 1){
								//verificar conflitos
								//se nao conflito
									//escalonar
								//se conflito 
									//wait die
									//parou
							}
						}
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







