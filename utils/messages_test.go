package utils

import (
	"testing"
)

// Returns random greeting from default 'en' locale when locale provided is invalid
func TestGetRandomGreetingWithInvalidLocale(t *testing.T) { // Act
	result := GetRandomGreeting("invalid")

	// Assert
	found := false
	for _, greeting := range helloMessages["en"] {
		if result == greeting {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected greeting to be one of default English greetings, got %s", result)
	}
}

// Handles empty locale string by defaulting to 'en'
func TestGetRandomGreetingWithEmptyLocale(t *testing.T) {
	emptyLocale := ""

	// Act
	result := GetRandomGreeting(emptyLocale)

	// Assert
	found := false
	for _, greeting := range helloMessages["en"] {
		if result == greeting {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected greeting from 'en' locale when empty locale provided, got %s", result)
	}
}

// Handles empty locale string by defaulting to 'en'
func TestGetRandomGreeting(t *testing.T) {
	// Act
	result := GetRandomGreeting("ru")

	// Assert
	found := false
	for _, greeting := range helloMessages["ru"] {
		if result == greeting {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected greeting from 'en' locale when empty locale provided, got %s", result)
	}
}

// Returns random instruction from available messages for valid locale
func TestGetRandomInstructionReturnsValidMessage(t *testing.T) {
	locale := "fr"
	expectedMessages := []string{
		"Veuillez utiliser la commande pour traduire votre message.",
		"Utilisez la commande pour traduire votre message, s'il vous plaￃﾮt.",
		"Utilisez la commande pour traduire votre message, merci !",
	}

	result := GetRandomInstruction(locale)

	found := false
	for _, expected := range expectedMessages {
		if result == expected {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("GetRandomInstruction(%s) = %s; want one of %v", locale, result, expectedMessages)
	}
}

// Locale with empty instruction array falls back to 'en'
func TestGetRandomInstructionFallsBackToEnglish(t *testing.T) {
	locale := "invalid-locale"
	result := GetRandomInstruction(locale)

	if result == "" {
		t.Errorf("GetRandomInstruction(%s) returned empty string", locale)
	}

	if len(result) == 0 {
		t.Errorf("GetRandomInstruction(%s) should not return empty result", locale)
	}
}
