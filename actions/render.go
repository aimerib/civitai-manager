package actions

import (
	"civitai/public"
	"civitai/templates"
	"html/template"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/helpers/forms"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.plush.html",

		// fs.FS containing templates
		TemplatesFS: templates.FS(),

		// fs.FS containing assets
		AssetsFS: public.FS(),

		// Add template helpers here:
		Helpers: render.Helpers{
			// for non-bootstrap form helpers uncomment the lines
			// below and import "github.com/gobuffalo/helpers/forms"
			forms.FormKey:    forms.Form,
			forms.FormForKey: forms.FormFor,
			"alertColorClass": func(k string) string {
				switch k {
				case "success":
					return "bg-green-100 border border-green-400 text-green-700"
				case "error":
					return "bg-red-100 border border-red-400 text-red-700"
				case "info":
					return "bg-blue-100 border border-blue-400 text-blue-700"
				case "warning":
					return "bg-yellow-100 border border-yellow-400 text-yellow-700"
				default:
					return "bg-gray-100 border border-gray-400 text-gray-700"
				}
			},
			"closeIconColorClass": func(k string) string {
				switch k {
				case "success":
					return "text-green-500 hover:text-green-800"
				case "error":
					return "text-red-500 hover:text-red-800"
				case "info":
					return "text-blue-500 hover:text-blue-800"
				case "warning":
					return "text-yellow-500 hover:text-yellow-800"
				default:
					return "text-gray-500 hover:text-gray-800"
				}
			},
			"flashIcon": func(flashType string) template.HTML {
				var icon string
				switch flashType {
				case "success":
					icon = `<svg class="h-6 w-6 text-green-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
							  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
							</svg>`
				case "error":
					icon = `<svg class="h-6 w-6 text-red-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
							  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
							</svg>`
				case "warning":
					icon = `<svg class="h-6 w-6 text-yellow-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
							  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
							</svg>`
				default:
					icon = `<svg class="h-6 w-6 text-blue-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
							  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
							</svg>`
				}
				return template.HTML(icon)
			},
			"flashTitle": func(flashType string) string {
				switch flashType {
				case "success":
					return "Success"
				case "error":
					return "Error"
				case "warning":
					return "Warning"
				default:
					return "Information"
				}
			},
		},
	})
}
