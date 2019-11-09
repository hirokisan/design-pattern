package empty_interface

import "fmt"

func Accept(content interface{}) error {
	switch v := content.(type) {
	case string:
	case int:
		return nil
	default:
		return fmt.Errorf("type %v is not supported", v)
	}
	return nil
}
