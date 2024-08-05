package actions

import (
	"civitai/public"
	"civitai/templates"

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
		},
	})
}
