package main

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// uuidGenerator is a component that displays a UUID generator page.
type uuidGenerator struct {
	app.Compo
	generatedUUID  string
	copyButtonText string
}

// OnMount is called when the component is mounted.
// We generate a UUID and set the initial button text when the app loads.
func (g *uuidGenerator) OnMount(ctx app.Context) {
	g.copyButtonText = "Copy"
	g.generateNewUUID(ctx)
}

// Render builds the UI of the component.
func (g *uuidGenerator) Render() app.UI {
	// The main div now has classes for both light and dark modes.
	return app.Div().Class("bg-white dark:bg-gray-900 text-gray-800 dark:text-gray-200 min-h-screen font-sans flex items-center justify-center transition-colors duration-300").Body(
		app.Div().Class("container mx-auto p-8 max-w-2xl text-center").Body(

			// Header Section
			app.Header().Class("mb-8").Body(
				app.H1().Class("text-5xl md:text-6xl font-bold text-gray-900 dark:text-white mb-2").Text("UUID Generator"),
				app.P().Class("text-xl md:text-2xl text-cyan-600 dark:text-cyan-400").Text("Generate new UUIDs with a single click."),
			),

			// Generator Section
			app.Main().Class("space-y-8").Body(
				app.Section().Body(
					// Display area for the generated UUID
					app.Div().Class("bg-gray-100 dark:bg-gray-800 p-4 rounded-lg shadow-lg flex items-center justify-between").Body(
						// Replaced the input with a span for a cleaner, non-editable display.
						app.Span().
							Class("text-xl text-gray-900 dark:text-white w-full text-left font-mono tracking-wider").
							ID("uuid-output").
							Text(g.generatedUUID),
						// Copy to clipboard button
						app.Button().
							Class("bg-cyan-600 dark:bg-cyan-500 hover:bg-cyan-500 dark:hover:bg-cyan-400 text-white font-bold py-2 px-4 rounded transition-colors ml-4 whitespace-nowrap").
							Text(g.copyButtonText).
							OnClick(g.copyToClipboard),
					),
				),

				// Action Button
				app.Section().Body(
					app.Button().
						Class("bg-green-600 dark:bg-green-500 hover:bg-green-500 dark:hover:bg-green-400 text-white font-bold py-3 px-8 rounded-full text-xl transition-transform transform hover:scale-105").
						Text("Generate New UUID").
						OnClick(g.onClick),
				),
			),
		),
	)
}

// onClick is the handler for the "Generate New UUID" button.
func (g *uuidGenerator) onClick(ctx app.Context, e app.Event) {
	g.generateNewUUID(ctx)
}

// generateNewUUID generates a new UUID and updates the component state.
func (g *uuidGenerator) generateNewUUID(ctx app.Context) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		log.Println("Error generating UUID:", err)
		g.generatedUUID = "Error generating UUID"
	} else {
		g.generatedUUID = newUUID.String()
	}
	log.Println("Generated UUID:", g.generatedUUID)
	ctx.Update() // Re-render the component to display the new UUID
}

// copyToClipboard copies the generated UUID to the user's clipboard and provides visual feedback.
func (g *uuidGenerator) copyToClipboard(ctx app.Context, e app.Event) {
	// This Javascript code now reads the innerText of the span.
	app.Raw(`
		const textToCopy = document.getElementById('uuid-output').innerText;
		const textArea = document.createElement('textarea');
		textArea.style.position = 'fixed';
		textArea.value = textToCopy;
		document.body.appendChild(textArea);
		textArea.focus();
		textArea.select();
		try {
			document.execCommand('copy');
		} catch (err) {
			console.error('Unable to copy', err);
		}
		document.body.removeChild(textArea);
	`)

	// Provide visual feedback by changing the button text.
	g.copyButtonText = "Copied!"
	ctx.Update()

	// Reset the button text after 2 seconds.
	ctx.After(2*time.Second, func(ctx app.Context) {
		g.copyButtonText = "Copy"
		ctx.Update()
	})
}

// main is the application's entry point.
func main() {
	// Associate the uuidGenerator component with the root path.
	app.Route("/", func() app.Composer { return &uuidGenerator{} })

	// Run the application.
	app.RunWhenOnBrowser()

	// Generate the static website files for deployment.
	err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "UUID Generator",
		Description: "A simple UUID generator built with go-app.",
		Styles: []string{
			// Using Tailwind CSS for styling.
			"https://cdn.tailwindcss.com/3.4.1/tailwind.min.css",
		},

		// Replace "your-github-username.github.io/your-repo-name" with your actual GitHub Pages URL.
		Resources: app.GitHubPages("uuid.calvinbrown.dev"),
	})

	if err != nil {
		log.Fatal(err)
	}
}
