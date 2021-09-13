package secretreader

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/alessio/shellescape.v1"
)

type SecretStore interface {
	Get(key string) (string, error)
}

type SecretReader struct {
	secretStore SecretStore
}

func NewSecretReader(secStore SecretStore) *SecretReader {
	return &SecretReader{
		secretStore: secStore,
	}

}

func (sr *SecretReader) GenerateSource(file string) error {
	// Read file
	keys := sr.readFile(file)
	// Get secret values
	secrets := map[string]string{}
	for _, key := range keys {
		key = strings.Trim(key, "{{")
		key = strings.Trim(key, "}}")
		val, err := sr.secretStore.Get(key)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("error getting %s", key))
		}

		secrets[key] = strings.TrimSpace(val)
	}
	// Populate source
	output, err := sr.populateSource(file, secrets)

	if err != nil {
		return err
	}

	fmt.Printf("%s\n", output)
	return nil
}

var reg = regexp.MustCompile(`{{(.*?)}}`)

func (sr *SecretReader) readFile(file string) []string {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return reg.FindAllString(string(body), -1)
}

func (sr *SecretReader) populateSource(file string, secrets map[string]string) (string, error) {
	body, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	temp := string(body)
	temp = strings.ReplaceAll(temp, "{{", "")
	temp = strings.ReplaceAll(temp, "}}", "")

	for key, val := range secrets {
		val = shellescape.Quote(val)
		temp = strings.ReplaceAll(temp, key, val)
	}

	return temp, nil

}
