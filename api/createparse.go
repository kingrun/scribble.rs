package api

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/scribble-rs/scribble.rs/game"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ParsePlayerName(value string) (string, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return trimmed, errors.New("vardo laukas negali būti tuščias.")
	}

	return trimmed, nil
}

func ParseLanguage(value string) (string, error) {
	toLower := strings.ToLower(strings.TrimSpace(value))
	for languageKey := range game.SupportedLanguages {
		if toLower == languageKey {
			return languageKey, nil
		}
	}

	return "", errors.New("pasirinkta kalba nepalaikoma")
}

func ParseDrawingTime(value string) (int, error) {
	result, parseErr := strconv.ParseInt(value, 10, 64)
	if parseErr != nil {
		return 0, errors.New("piešimo laikas turi susidaryti iš skaičių")
	}

	if result < game.LobbySettingBounds.MinDrawingTime {
		return 0, fmt.Errorf("piešimo laikas negali būti mažesnis nei %d", game.LobbySettingBounds.MinDrawingTime)
	}

	if result > game.LobbySettingBounds.MaxDrawingTime {
		return 0, fmt.Errorf("peišimo laikas negali būti didesnis nei %d", game.LobbySettingBounds.MaxDrawingTime)
	}

	return int(result), nil
}

func ParseRounds(value string) (int, error) {
	result, parseErr := strconv.ParseInt(value, 10, 64)
	if parseErr != nil {
		return 0, errors.New("raundai turi susidaryti iš skaičių")
	}

	if result < game.LobbySettingBounds.MinRounds {
		return 0, fmt.Errorf("raundų skaičius negali būti mažesnis nei %d", game.LobbySettingBounds.MinRounds)
	}

	if result > game.LobbySettingBounds.MaxRounds {
		return 0, fmt.Errorf("raundų skaičius negali būti didesnis nei %d", game.LobbySettingBounds.MaxRounds)
	}

	return int(result), nil
}

func ParseMaxPlayers(value string) (int, error) {
	result, parseErr := strconv.ParseInt(value, 10, 64)
	if parseErr != nil {
		return 0, errors.New("žaidėjų kiekis turi susidaryti iš skaičių")
	}

	if result < game.LobbySettingBounds.MinMaxPlayers {
		return 0, fmt.Errorf("maximalus žaidėjų skaičius turi būti ne mažesnins nei %d", game.LobbySettingBounds.MinMaxPlayers)
	}

	if result > game.LobbySettingBounds.MaxMaxPlayers {
		return 0, fmt.Errorf("maximalus žaidėjų skaičius turi būti ne didesnis nei %d", game.LobbySettingBounds.MaxMaxPlayers)
	}

	return int(result), nil
}

func ParseCustomWords(value string) ([]string, error) {
	trimmedValue := strings.TrimSpace(value)
	if trimmedValue == "" {
		return nil, nil
	}

	lowercaser := cases.Lower(language.English)
	result := strings.Split(trimmedValue, ",")
	for index, item := range result {
		trimmedItem := lowercaser.String(strings.TrimSpace(item))
		if trimmedItem == "" {
			return nil, errors.New("papildomų žodžių laukas negali būti tuščias")
		}
		result[index] = trimmedItem
	}

	return result, nil
}

func ParseClientsPerIPLimit(value string) (int, error) {
	result, parseErr := strconv.ParseInt(value, 10, 64)
	if parseErr != nil {
		return 0, errors.New("maximalus žaidėjų kiekis per IP turi susidaryti iš skaičių")
	}

	if result < game.LobbySettingBounds.MinClientsPerIPLimit {
		return 0, fmt.Errorf("žaidėjų per IP kiekis negali būti mažesnis nei %d", game.LobbySettingBounds.MinClientsPerIPLimit)
	}

	if result > game.LobbySettingBounds.MaxClientsPerIPLimit {
		return 0, fmt.Errorf("žaidėjų per IP kiekis negali būti didesnis nei %d", game.LobbySettingBounds.MaxClientsPerIPLimit)
	}

	return int(result), nil
}

func ParseCustomWordsChance(value string) (int, error) {
	result, parseErr := strconv.ParseInt(value, 10, 64)
	if parseErr != nil {
		return 0, errors.New("papildomų žodžių šansas vedamas skaičiais")
	}

	if result < 0 {
		return 0, errors.New("papildomų žodžių šansas negali būti mažesnis nei 0")
	}

	if result > 100 {
		return 0, errors.New("papildomų žodžių šansas negali būti didesnis nei 100")
	}

	return int(result), nil
}

func ParseBoolean(valueName string, value string) (bool, error) {
	if strings.EqualFold(value, "true") {
		return true, nil
	}

	if strings.EqualFold(value, "false") {
		return false, nil
	}

	if value == "" {
		return false, nil
	}

	return false, fmt.Errorf("the %s value must be a boolean value ('true' or 'false)", valueName)
}
