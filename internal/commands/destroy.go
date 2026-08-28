package commands

import (
	"fmt"

	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
	"github.com/fatih/color"
	"github.com/johnathantam/TodoTerminal/internal/storage"
)

func confirmDestroyByGUI() (bool, error) {
	var confirmed bool

	theme := huh.ThemeFunc(func(isDark bool) *huh.Styles {
		styles := huh.ThemeBase(isDark)
		styles.Focused.FocusedButton = styles.Focused.FocusedButton.
			Background(lipgloss.Color("#FF5555")). // red — the selected option
			Foreground(lipgloss.Color("#FFFFFF"))
		styles.Focused.BlurredButton = styles.Focused.BlurredButton.
			Background(lipgloss.Color("#333333")). // dark grey — the unselected option
			Foreground(lipgloss.Color("#AAAAAA"))
		return styles
	})

	err := huh.NewConfirm().
		Title("Delete .todoterminal?").
		Description("This will permanently delete all projects and tasks. This cannot be undone.").
		Affirmative("Yes, delete everything").
		Negative("Cancel").
		Value(&confirmed).
		WithTheme(theme).
		Run()
	if err != nil {
		return false, err
	}
	return confirmed, nil
}

func Destroy(args []string) error {
	paths, err := storage.FindLocationOfAppPaths()
	if err != nil {
		return err
	}

	confirmed, err := confirmDestroyByGUI()
	if err != nil {
		return err
	}
	if !confirmed {
		fmt.Println("Cleanup cancelled.")
		return nil
	}

	if err := storage.RemoveAppStructure(paths.AppDirectoryPath); err != nil {
		return err
	}

	color.New(color.FgGreen).Println("TodoTerminal project deleted.")

	return nil
}
