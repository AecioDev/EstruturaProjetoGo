package main

import (
	"log"

	"meu-projeto/config"
	"meu-projeto/internal/api/server"
	"meu-projeto/internal/db"
)

func main() {
	// Carregar configurações meu-projeto
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// Inicializar banco de dados
	database, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Inicializar e executar o servidor
	s := server.NewServer(cfg, database)
	if err := s.Run(); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
