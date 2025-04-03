package main

import (
	"fmt"
	"os"
	"time"
	"gerenciador-tarefas/tarefas"
)

func main() {
	gerenciador := tarefas.NovoGerenciadorDeTarefas()

	const diretorio = "data"
    const nomeArquivo = "tarefas.json"
    const caminhoArquivo = diretorio + "/" + nomeArquivo

	err := gerenciador.CarregarTarefas(caminhoArquivo)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Erro ao carregar tarefas:", err)
	}

	resultadoProcessamento := make(chan string)

	go func() {
		for resultado := range resultadoProcessamento {
			fmt.Println(resultado)
		}
	}()

	defer close(resultadoProcessamento)

	for {
		temTarefas := gerenciador.TemTarefas()

		fmt.Println("\nEscolha uma opção:")
		fmt.Println("1. Adicionar tarefa")

		if temTarefas {
			fmt.Println("2. Listar tarefas")
			fmt.Println("3. Concluir tarefa")
			fmt.Println("4. Excluir tarefa concluída")
			fmt.Println("5. Editar tarefa")
			fmt.Println("6. Mostrar estatísticas")
			fmt.Println("7. Adicionar tarefa (processamento assíncrono)")
			fmt.Println("8. Sair")
		} else {
			fmt.Println("2. Adicionar tarefa (processamento assíncrono)")
			fmt.Println("3. Sair")
		}
		
        var opcao int
        fmt.Scan(&opcao)

        if !temTarefas {
            if opcao == 2 {
                opcao = 7 
            } else if opcao == 3 {
                opcao = 8 
            } else if opcao > 3 {
                fmt.Println("Opção inválida!")
                continue
            }
        }

        switch opcao {
        case 1:
            fmt.Print("Digite o título da tarefa: ")
            var titulo string
            fmt.Scan(&titulo)
            fmt.Print("Digite a prioridade (1 - Alta, 2 - Média, 3 - Baixa): ")
            var opcaoPrioridade int
            fmt.Scan(&opcaoPrioridade)
            
            var prioridadeTarefa tarefas.Prioridade
            switch opcaoPrioridade {
            case 1:
                prioridadeTarefa = tarefas.Alta
            case 2:
                prioridadeTarefa = tarefas.Media
            case 3:
                prioridadeTarefa = tarefas.Baixa
            default:
                fmt.Println("Prioridade inválida!")
                continue
            }
            
            gerenciador.AdicionarTarefa(titulo, prioridadeTarefa)
            if err := gerenciador.SalvarTarefas(caminhoArquivo); err != nil {
                fmt.Println("Erro ao salvar tarefas:", err)
            }

        case 2:
            gerenciador.ListarTarefas()
            fmt.Println("\nOpções adicionais para listar tarefas:")
            fmt.Println("1. Filtrar por status")
            fmt.Println("2. Filtrar por prioridade")
            fmt.Println("3. Voltar ao menu principal")
            var opcaoListar int
            fmt.Scan(&opcaoListar)
            switch opcaoListar {
            case 1:
                fmt.Print("Digite o status para filtrar (1 - pendente, 2 - concluida): ")
                var opcaoStatus int
                fmt.Scan(&opcaoStatus)
                
                var status string
                switch opcaoStatus {
                case 1:
                    status = "pendente"
                case 2:
                    status = "concluida"
                default:
                    fmt.Println("Status inválido!")
                    continue
                }
                
                gerenciador.FiltrarTarefasPorStatus(status)
                
            case 2:
                fmt.Print("Digite a prioridade para filtrar (1 - Alta, 2 - Média, 3 - Baixa): ")
                var opcaoPrioridade int
                fmt.Scan(&opcaoPrioridade)
                
                var prioridadeTarefa tarefas.Prioridade
                switch opcaoPrioridade {
                case 1:
                    prioridadeTarefa = tarefas.Alta
                case 2:
                    prioridadeTarefa = tarefas.Media
                case 3:
                    prioridadeTarefa = tarefas.Baixa
                default:
                    fmt.Println("Prioridade inválida!")
                    continue
                }
                
                gerenciador.FiltrarTarefasPorPrioridade(prioridadeTarefa)
                
            case 3:
                continue
            default:
                fmt.Println("Opção inválida!")
            }

        case 3:
            fmt.Print("Digite o ID da tarefa a ser concluída: ")
            var id int
            fmt.Scan(&id)
            gerenciador.ConcluirTarefa(id)
            if err := gerenciador.SalvarTarefas(caminhoArquivo); err != nil {
                fmt.Println("Erro ao salvar tarefas:", err)
            }

        case 4:
            fmt.Print("Digite o ID da tarefa a ser excluída: ")
            var id int
            fmt.Scan(&id)
            gerenciador.ExcluirTarefaConcluida(id)
            if err := gerenciador.SalvarTarefas(caminhoArquivo); err != nil {
                fmt.Println("Erro ao salvar tarefas:", err)
            }

        case 5:
            fmt.Print("Digite o ID da tarefa a ser editada: ")
            var id int
            fmt.Scan(&id)
            fmt.Print("Digite o novo título: ")
            var novoTitulo string
            fmt.Scan(&novoTitulo)
            gerenciador.EditarTarefa(id, novoTitulo)
            if err := gerenciador.SalvarTarefas(caminhoArquivo); err != nil {
                fmt.Println("Erro ao salvar tarefas:", err)
            }

        case 6:
            gerenciador.Estatisticas()

        case 7:
            fmt.Print("Digite o título da tarefa: ")
            var titulo string
            fmt.Scan(&titulo)
            fmt.Print("Digite a prioridade (1 - Alta, 2 - Média, 3 - Baixa): ")
            var opcaoPrioridade int
            fmt.Scan(&opcaoPrioridade)
            
            var prioridadeTarefa tarefas.Prioridade
            switch opcaoPrioridade {
            case 1:
                prioridadeTarefa = tarefas.Alta
            case 2:
                prioridadeTarefa = tarefas.Media
            case 3:
                prioridadeTarefa = tarefas.Baixa
            default:
                fmt.Println("Prioridade inválida!")
                continue
            }
            
            fmt.Println("Processando tarefa em segundo plano...")
            go gerenciador.ProcessarTarefaAsync(titulo, prioridadeTarefa, resultadoProcessamento)
            
            
            go func() {
                time.Sleep(3 * time.Second)
                if err := gerenciador.SalvarTarefas(caminhoArquivo); err != nil {
                    fmt.Println("Erro ao salvar tarefas:", err)
                }
            }()

        case 8:
            fmt.Println("Saindo...")
            return

        default:
            fmt.Println("Opção inválida!")
        }
    }
}
