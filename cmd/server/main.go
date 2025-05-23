package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hvmello/goal-tracker-backend/internal/config"
)

func main() {
	// Inicializa conexão com o banco de dados
	db, err := config.NewDBConnection()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Executa as migrações automáticas
	if err := config.AutoMigrate(db); err != nil {
		log.Fatalf("Erro ao executar migrações: %v", err)
	}

	log.Println("Migrações executadas com sucesso!")

	// Configuração das rotas
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"status": "ok"}
		json.NewEncoder(w).Encode(response)
	})

	// Inicia o servidor
	log.Println("Servidor iniciando na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
