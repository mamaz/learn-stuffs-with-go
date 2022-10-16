package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Person struct {
	ID   string
	Name string
}

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.DurationFieldUnit = time.Millisecond
	person := Person{ID: "12", Name: "Mamazo"}

	personJSON, _ := json.Marshal(person)

	log.Info().RawJSON("Person", personJSON).Dur("DurationMs", time.Duration(200*time.Millisecond)).Msg("Person")
	log.Error().Stack().Err(fmt.Errorf("some error")).Msg("Error nih!")
}
