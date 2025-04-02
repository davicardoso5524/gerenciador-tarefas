# Gerenciador de Tarefas

## Descrição
Este é um sistema de Gerenciamento de Tarefas escrito em **Go**, que permite adicionar, listar, concluir, excluir e editar tarefas. Também suporta processamento assíncrono para adição de tarefas em segundo plano.

Este projeto foi desenvolvido para a disciplina de Linguagens de Programação do curso de DSI - Desenvolvimento de Sistemas de Informação no campus de Itabaiana.

Feito por Davi Cardoso, Joao Keweni, Jose Adryel e Marcos Yago.

## Funcionalidades
- **Adicionar Tarefa**: Adiciona uma nova tarefa com prioridade (Alta, Média, Baixa).
- **Listar Tarefas**: Exibe todas as tarefas cadastradas.
- **Filtrar Tarefas**: Permite filtrar por status ou prioridade.
- **Concluir Tarefa**: Marca uma tarefa como concluída.
- **Excluir Tarefa**: Remove uma tarefa concluída.
- **Editar Tarefa**: Permite alterar o título de uma tarefa existente.
- **Estatísticas**: Mostra estatísticas sobre as tarefas.
- **Processamento Assíncrono**: Permite adicionar tarefas em segundo plano.
- **Persistência de Dados**: Tarefas são salvas em um arquivo JSON para manutenção dos dados entre execuções.

## Requisitos
- Go 1.16 ou superior.

## Como Executar
1. Clone o repositório:
   ```sh
   git clone https://github.com/seu-usuario/gerenciador-tarefas.git
   ```
2. Acesse a pasta do projeto:
   ```sh
   cd gerenciador-tarefas
   ```
3. Execute o programa:
   ```sh
   go run main.go
   ```

## Estrutura do Projeto
```
/gerenciador-tarefas
├── main.go           # Ponto de entrada do programa
├── tarefas/          # Pacote responsável pelo gerenciamento de tarefas
│   ├── tarefas.go    # Implementação das funcionalidades de tarefas
├── data/             # Diretório onde as tarefas são salvas em JSON
└── README.md         # Documentação do projeto
```
