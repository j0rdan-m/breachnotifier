package checker

import "fmt"

func GetChecker(checkerType string) (Checker, error) {
	switch checkerType {
	case "leakcheck":
		return NewLeakCheck(), nil
	default:
		return nil, fmt.Errorf("type de checker inconnu : %s", checkerType)
	}
}
