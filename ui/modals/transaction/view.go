package transaction

import (
	"github.com/algorandfoundation/hack-tui/internal"
	"github.com/algorandfoundation/hack-tui/ui/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
)

func (m ViewModel) View() string {
	if m.Participation == nil {
		return "No key selected"
	}
	if m.ATxn == nil {
		return "Loading..."
	}
	// TODO: Refactor ATxn to Interface
	txn, err := m.ATxn.ProduceQRCode()
	if err != nil {
		return "Something went wrong"
	}
	link, _ := internal.ToLoraDeepLink(m.State.Status.Network, m.Active, *m.Participation)
	render := lipgloss.JoinVertical(
		lipgloss.Center,
		"Scan the QR code with your wallet",
		qrStyle.Render(txn),
		style.WithHyperlink("click here to open in Lora", link),
	)

	width := lipgloss.Width(render)
	height := lipgloss.Height(render)

	if width > m.Width || height > m.Height {
		return lipgloss.JoinVertical(
			lipgloss.Center,
			style.Red.Render(ansi.Wordwrap("QR Code too large to display... Please adjust terminal dimensions or font size.", m.Width, " ")),
			"",
			style.WithHyperlink("or click here to open in Lora", link),
		)
	}

	return render
}
