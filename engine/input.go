package engine

import "fmt"

type Keyboard struct {
	pressedKeys map[string]bool
}

func (k *Keyboard) Press(keyName string) {
	fmt.Println("Pressed key", keyName)
	k.pressedKeys[keyName] = true
}

func (k *Keyboard) Release(keyName string) {
	fmt.Println("Released key", keyName)
	k.pressedKeys[keyName] = false
}

func (k *Keyboard) IsPressed(keyName string) bool {
	if value, keyExists := k.pressedKeys[keyName]; keyExists {
		return value
	}
	return false
}

type Input struct {
	Keyboard *Keyboard
}

func NewInput() *Input {
	return &Input{
		Keyboard: &Keyboard{
			pressedKeys: make(map[string]bool),
		},
	}
}
