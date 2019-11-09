package marker_interface

import "fmt"

// Content : メソッドを定義していない、目印の役割だけ
type Content interface{}

func Accept(content Content) error {
	switch v := content.(type) {
	case string:
	case int:
		return nil
	default:
		return fmt.Errorf("type %v is not supported", v)
	}
	return nil
}
