package cli

import "testing"

func TestCommands(t *testing.T) {
	t.Run("default program with no additional params", func(t *testing.T) {
		got := NoParams()
		want := "TUI is not implemented yet."
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
	t.Run("text passed with no arguments", func(t *testing.T) {
		got := ParseText("Note text that is passed")
		want := "Note text that is passed"
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}
