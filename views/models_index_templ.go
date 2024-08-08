// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"civitai-manager/models"
	"fmt"
)

func ModelsIndex(models []models.Model) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"mx-auto\"><div class=\"grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, model := range models {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"bg-background-card cursor-pointer rounded-lg shadow-md overflow-hidden transition-transform duration-300 hover:scale-105\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/models/%d", model.ID))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/models_index.templ`, Line: 14, Col: 49}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#content\" hx-push-url=\"true\" hx-swap=\"innerHTML\"><div class=\"relative\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(model.ModelVersions) > 0 && len(model.ModelVersions[0].Images) > 0 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"media-container\" data-src=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(model.ModelVersions[0].Images[0].URL)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/models_index.templ`, Line: 23, Col: 55}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" data-alt=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(model.Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/models_index.templ`, Line: 24, Col: 29}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><!-- Content will be inserted here by JavaScript --></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if !model.Checked {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"absolute top-2 right-2 bg-green-500 text-white rounded-full p-1\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-4 w-4\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M5 13l4 4L19 7\"></path></svg></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"p-4\"><h2 class=\"text-lg font-semibold truncate\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(model.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/models_index.templ`, Line: 49, Col: 61}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(model.ModelVersions) > 0 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-sm text-gray-600\">Published: ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(model.ModelVersions[0].PublishedAt.Format("2006-01-02"))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/models_index.templ`, Line: 52, Col: 76}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = modelGridScript().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func modelGridScript() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_modelGridScript_a333`,
		Function: `function __templ_modelGridScript_a333(){(function() {
		function loadImage(url) {
			return fetch(url)
				.then((response) => response.blob())
				.then((blob) => URL.createObjectURL(blob))
				.catch((error) => console.error("Error loading image:", error));
		}

		function setupMediaContainers() {
			const containers = document.querySelectorAll(".media-container");
			containers.forEach((container) => {
				const url = container.dataset.src;
				const alt = container.dataset.alt;

				loadImage(url).then((objectUrl) => {
					const img = new Image();
					img.onload = function () {
						container.innerHTML = ` + "`" + `<img src="${objectUrl}" alt="${alt}" class="w-full h-48 object-cover" style="object-position: left -20px">` + "`" + `;
					};
					img.onerror = function () {
						const video = document.createElement("video");
						video.onloadedmetadata = function () {
							container.innerHTML = ` + "`" + `
								<video class="w-full h-48 object-cover" controls style="object-position: left -20px">
									<source src="${url}" type="video/mp4">
									Your browser does not support the video tag.
								</video>
							` + "`" + `;
						};
						video.onerror = function () {
							container.innerHTML = ` + "`" + `<p>Unsupported media type</p>` + "`" + `;
						};
						video.src = url;
					};
					img.src = objectUrl;
				});
			});
		}

		if ("serviceWorker" in navigator) {
			window.addEventListener("load", function () {
				navigator.serviceWorker.register('/sw.js').then(
					function (registration) {
						console.log(
							"ServiceWorker registration successful with scope: ",
							registration.scope
						);
					},
					function (err) {
						console.log("ServiceWorker registration failed: ", err);
					}
				);
			});
		}
		setupMediaContainers();
	})();

	// htmx.on("htmx:load", setupMediaContainers);
	// htmx.on("htmx:beforeSwap", cleanupPage);
}`,
		Call:       templ.SafeScript(`__templ_modelGridScript_a333`),
		CallInline: templ.SafeScriptInline(`__templ_modelGridScript_a333`),
	}
}
