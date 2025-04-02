package tarefas

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
	"path/filepath"
)

// Prioridade representa a prioridade de uma tarefa
type Prioridade string

const (
	Alta    Prioridade = "Alta"
	Media   Prioridade = "Média"
	Baixa   Prioridade = "Baixa"
)

// Tarefa representa uma tarefa no sistema
type Tarefa struct {
	ID        int       `json:"id"`
	Titulo    string    `json:"titulo"`
	Prioridade Prioridade `json:"prioridade"`
	Concluida bool      `json:"concluida"`
}

// GerenciadorDeTarefas gerencia as tarefas
type GerenciadorDeTarefas struct {
	mu       sync.Mutex
	tarefas  []Tarefa
	proximoID int
}

// NovoGerenciadorDeTarefas cria um novo gerenciador de tarefas
func NovoGerenciadorDeTarefas() *GerenciadorDeTarefas {
	return &GerenciadorDeTarefas{
		tarefas:  []Tarefa{},
		proximoID: 1,
	}
}

// TemTarefas verifica se existem tarefas no gerenciador
func (gt *GerenciadorDeTarefas) TemTarefas() bool {
	gt.mu.Lock()
	defer gt.mu.Unlock()
	return len(gt.tarefas) > 0
}

// ProcessarTarefaAsync processa uma tarefa de forma assíncrona
func (gt *GerenciadorDeTarefas) ProcessarTarefaAsync(titulo string, prioridade Prioridade, resultado chan<- string) {
	// Simula um processamento que leva tempo
	time.Sleep(2 * time.Second)
	
	gt.mu.Lock()
	defer gt.mu.Unlock()
	
	tarefa := Tarefa{
		ID:        gt.proximoID,
		Titulo:    titulo,
		Prioridade: prioridade,
	}
	gt.tarefas = append(gt.tarefas, tarefa)
	gt.proximoID++
	
	resultado <- fmt.Sprintf("Tarefa '%s' processada e adicionada com sucesso!", titulo)
}

// AdicionarTarefa adiciona uma nova tarefa
func (gt *GerenciadorDeTarefas) AdicionarTarefa(titulo string, prioridade Prioridade) {
	gt.mu.Lock()
	defer gt.mu.Unlock()
	tarefa := Tarefa{
		ID:        gt.proximoID,
		Titulo:    titulo,
		Prioridade: prioridade,
	}
	gt.tarefas = append(gt.tarefas, tarefa)
	gt.proximoID++
	fmt.Printf("Tarefa '%s' adicionada com sucesso!\n", titulo)
}

// ListarTarefas lista todas as tarefas
func (gt *GerenciadorDeTarefas) ListarTarefas() {
	gt.mu.Lock()
	defer gt.mu.Unlock()
	fmt.Println("Tarefas:")
	for _, tarefa := range gt.tarefas {
		status := "Pendente"
		if tarefa.Concluida {
			status = "Concluída"
		}
		fmt.Printf("- [%d] %s (%s) - Prioridade: %s\n", tarefa.ID, tarefa.Titulo, status, tarefa.Prioridade)
	}
}

// ConcluirTarefa marca uma tarefa como concluída
func (gt *GerenciadorDeTarefas) ConcluirTarefa(id int) {
	gt.mu.Lock()
	defer gt.mu.Unlock()
	for i := range gt.tarefas {
		if gt.tarefas[i].ID == id {
			gt.tarefas[i].Concluida = true
			fmt.Printf("Tarefa '%s' marcada como concluída!\n", gt.tarefas[i].Titulo)
			return
		}
	}
	fmt.Println("Tarefa não encontrada!")
}

// ExcluirTarefaConcluida exclui uma tarefa concluída
func (gt *GerenciadorDeTarefas) ExcluirTarefaConcluida(id int) {
	gt.mu.Lock()
	defer gt.mu.Unlock()
	for i, tarefa := range gt.tarefas {
		if tarefa.ID == id && tarefa.Concluida {
			gt.tarefas = append(gt.tarefas[:i], gt.tarefas[i+1:]...)
			fmt.Printf("Tarefa '%d' excluída com sucesso!\n", id)
			return
		}
	}
	fmt.Println("Tarefa não encontrada ou não está concluída!")
}

// EditarTarefa edita uma tarefa existente
func (gt *GerenciadorDeTarefas) EditarTarefa(id int, novoTitulo string) {
	gt.mu.Lock()
	defer gt.mu.Unlock()
	for i := range gt.tarefas {
		if gt.tarefas[i].ID == id {
			gt.tarefas[i].Titulo = novoTitulo
			fmt.Printf("Tarefa '%d' editada com sucesso!\n", id)
			return
		}
	}
	fmt.Println("Tarefa não encontrada!")
}

// FiltrarTarefasPorStatus filtra tarefas por status
func (gt *GerenciadorDeTarefas) FiltrarTarefasPorStatus(status string) {
	gt.mu.Lock()
	defer gt.mu.Unlock()
	fmt.Println("Tarefas filtradas:")
	for _, tarefa := range gt.tarefas {
		if (status == "pendente" && !tarefa.Concluida) || (status == "concluida" && tarefa.Concluida) {
			fmt.Printf("- [%d] %s\n", tarefa.ID, tarefa.Titulo)
		}
	}
}

// FiltrarTarefasPorPrioridade filtra tarefas por prioridade
func (gt *GerenciadorDeTarefas) FiltrarTarefasPorPrioridade(prioridade Prioridade) {
	gt.mu.Lock()
	defer gt.mu.Unlock()
	fmt.Println("Tarefas filtradas:")
	for _, tarefa := range gt.tarefas {
		if tarefa.Prioridade == prioridade {
			fmt.Printf("- [%d] %s\n", tarefa.ID, tarefa.Titulo)
		}
	}
}

// Estatisticas mostra estatísticas básicas
func (gt *GerenciadorDeTarefas) Estatisticas() {
	gt.mu.Lock()
	defer gt.mu.Unlock()
	pendentes := 0
	concluidas := 0
	for _, tarefa := range gt.tarefas {
		if tarefa.Concluida {
			concluidas++
		} else {
			pendentes++
		}
	}
	fmt.Printf("Total de Tarefas: %d\n", len(gt.tarefas))
	fmt.Printf("Tarefas Pendentes: %d\n", pendentes)
	fmt.Printf("Tarefas Concluídas: %d\n", concluidas)
}

func (gt *GerenciadorDeTarefas) SalvarTarefas(caminho string) error {
    gt.mu.Lock()
    defer gt.mu.Unlock()
    
    // Cria o diretório se não existir
    dir := filepath.Dir(caminho)
    if err := os.MkdirAll(dir, os.ModePerm); err != nil {
        return fmt.Errorf("erro ao criar diretório: %v", err)
    }
    
    file, err := os.Create(caminho)
    if err != nil {
        return fmt.Errorf("erro ao criar arquivo: %v", err)
    }
    defer file.Close()

    data, err := json.MarshalIndent(gt.tarefas, "", "  ")
    if err != nil {
        return fmt.Errorf("erro ao converter para JSON: %v", err)
    }

    _, err = file.Write(data)
    if err != nil {
        return fmt.Errorf("erro ao escrever no arquivo: %v", err)
    }
    
    fmt.Println("Tarefas salvas com sucesso em:", caminho)
    return nil
}

func (gt *GerenciadorDeTarefas) CarregarTarefas(caminho string) error {
    gt.mu.Lock()
    defer gt.mu.Unlock()

    // Verifica se o arquivo existe
    if _, err := os.Stat(caminho); os.IsNotExist(err) {
        fmt.Println("Arquivo de tarefas não encontrado. Iniciando com lista vazia.")
        return nil
    }

    file, err := os.Open(caminho)
    if err != nil {
        return fmt.Errorf("erro ao abrir arquivo: %v", err)
    }
    defer file.Close()

    var tarefas []Tarefa
    err = json.NewDecoder(file).Decode(&tarefas)
    if err != nil {
        return fmt.Errorf("erro ao decodificar JSON: %v", err)
    }

    gt.tarefas = tarefas
    if len(tarefas) > 0 {
        gt.proximoID = tarefas[len(tarefas)-1].ID + 1
    } else {
        gt.proximoID = 1
    }
    fmt.Println("Tarefas carregadas com sucesso!")
    return nil
}
