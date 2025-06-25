package translators

import (
	"encoding/json"
	"fmt"
)

func Convert[Source any, Dest any](source Source) (*Dest, error) {
	// Passo 1: Converter a struct de origem para []byte JSON.
	bytes, err := json.Marshal(source)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter origem para JSON: %w", err)
	}

	// Passo 2: Criar um ponteiro para o tipo de destino.
	dest := new(Dest)

	// Passo 3: Converter o []byte JSON para a struct de destino.
	if err := json.Unmarshal(bytes, dest); err != nil {
		return nil, fmt.Errorf("erro ao converter JSON para destino: %w", err)
	}

	return dest, nil
}
