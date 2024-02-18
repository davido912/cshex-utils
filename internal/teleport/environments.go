package teleport

import "fmt"

type Environment string

const (
	development Environment = "development"
	staging     Environment = "staging"
	production  Environment = "production"
)

func NewEnv(env string) (Environment, error) {
	switch Environment(env) {
	case development:
		return development, nil
	case staging:
		return staging, nil
	case production:
		return production, nil
	}
	return "", fmt.Errorf("unknown environment")
}
