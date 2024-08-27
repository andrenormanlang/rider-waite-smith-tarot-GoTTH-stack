// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "andrenormanlang/tarot-go-htmx/common"

func Home(cards []common.Card, selectedCards []common.Card, meanings []string, isShuffling bool) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Free Tarot Reading</title><link href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css\" rel=\"stylesheet\"><link href=\"https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css\" rel=\"stylesheet\"><link href=\"/static/css/styles.css\" rel=\"stylesheet\"><script src=\"/static/script/htmx.min.js\"></script></head><body class=\"bg-gray-100 text-center py-4\"><div class=\"container mx-auto px-2\"><h1 class=\"text-3xl font-bold mb-2\">Free Tarot Reading</h1><p class=\"text-gray-700 mb-4\">Get the answers you need with this 3-card Tarot spread.</p><input type=\"text\" placeholder=\"Enter your question or subject here (optional)\" id=\"tarot-question\" class=\"w-full max-w-lg mx-auto p-2 border border-gray-300 rounded mb-4\"><!-- Shuffle/Stop Shuffle Button -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if isShuffling {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button hx-get=\"/stop-shuffle\" hx-target=\"body\" class=\"bg-red-500 text-white px-4 py-2 rounded-lg\">Stop Shuffle</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button hx-get=\"/shuffle-cards\" hx-target=\"body\" class=\"bg-purple-500 text-white px-4 py-2 rounded-lg\">Shuffle Cards</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Shuffled Cards Container --><div id=\"shuffled-cards\" class=\"relative h-48 mt-4\"><!-- Loop through each card -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, card := range cards {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card-container\" style=\"--card-index: {index};\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !isShuffling {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button hx-get=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var2 string
				templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs("/select-card?card=" + card.Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home.templ`, Line: 36, Col: 81}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"body\" class=\"bg-transparent\"><img src=\"/images/CardBacks.png\" alt=\"Card Back\" class=\"w-full h-auto\"></button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"bg-gray-400 p-1 rounded-lg animate-shuffle\"><img src=\"/images/CardBacks.png\" alt=\"Card Back\" class=\"w-full h-auto\"></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><!-- Selected Cards Container --><div id=\"selected-cards\" class=\"mt-4 grid grid-cols-1 sm:grid-cols-3 gap-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, card := range selectedCards {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flip-card\"><div class=\"flip-card-inner\"><div class=\"flip-card-front\"><img src=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs("/images/" + card.Image)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home.templ`, Line: 54, Col: 69}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"{ card.Name }\" class=\"w-full h-auto\"></div><div class=\"flip-card-back\"><h2 class=\"text-md font-semibold\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(card.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home.templ`, Line: 57, Col: 81}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><p class=\"text-sm text-gray-700\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(card.MeaningUp)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home.templ`, Line: 58, Col: 85}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><button type=\"button\" class=\"btn btn-primary mt-2\" data-bs-toggle=\"modal\" data-bs-target=\"#cardModal\" onclick=\"populateModal({ card })\">View Full Description</button></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><!-- Modal for Extended Card Description --><div class=\"modal fade\" id=\"cardModal\" tabindex=\"-1\" aria-labelledby=\"cardModalLabel\" aria-hidden=\"true\"><div class=\"modal-dialog\"><div class=\"modal-content\"><div class=\"modal-header\"><h5 class=\"modal-title\" id=\"cardModalLabel\">Card Title</h5><button type=\"button\" class=\"btn-close\" data-bs-dismiss=\"modal\" aria-label=\"Close\"></button></div><div class=\"modal-body\"><p id=\"modal-desc\"></p></div><div class=\"modal-footer\"><button type=\"button\" class=\"btn btn-secondary\" data-bs-dismiss=\"modal\">Close</button></div></div></div></div><!-- Script to Reveal Cards --><script src=\"https://cdn.jsdelivr.net/npm/@popperjs/core@2.11.7/dist/umd/popper.min.js\"></script><script src=\"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.min.js\"></script><script src=\"/static/script/script.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
