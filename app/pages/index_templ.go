// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Index() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"container\"><section id=\"top\" class=\"h-screen grid grid-rows-[10vh_1fr_10vh]\"><div class=\"flex flex-col items-center justify-center gap-2 row-start-2\"><span class=\"text-4xl text-black font-bold\">Congratulations!</span> <span class=\"text-2xl text-black\">You have set up the GoTTHS stack correctly. Check out the documentation of each component below!</span><div class=\"flex gap-2 items-center justify-center mt-4\"><sl-button size=\"large\" outline pill href=\"https://go.dev/\" target=\"_blank\">Go <iconify-icon icon=\"simple-icons:go\" slot=\"prefix\"></iconify-icon></sl-button> <sl-button size=\"large\" outline pill href=\"https://templ.guide/\" target=\"_blank\">Templ <iconify-icon icon=\"simple-icons:htmx\" slot=\"prefix\"></iconify-icon></sl-button> <sl-button size=\"large\" outline pill href=\"https://tailwindcss.com/\" target=\"_blank\">Tailwind CSS <iconify-icon icon=\"simple-icons:tailwindcss\" slot=\"prefix\"></iconify-icon></sl-button> <sl-button size=\"large\" outline pill href=\"https://htmx.org/\" target=\"_blank\">HTMX <iconify-icon icon=\"simple-icons:htmx\" slot=\"prefix\"></iconify-icon></sl-button> <sl-button size=\"large\" outline pill href=\"https://shoelace.style/\" target=\"_blank\">Shoelace <iconify-icon icon=\"fa6-brands:shoelace\" slot=\"prefix\"></iconify-icon></sl-button> <sl-button size=\"large\" outline pill href=\"https://github.com/air-verse/air\" target=\"_blank\">Air <iconify-icon icon=\"lucide:cloud\" slot=\"prefix\"></iconify-icon></sl-button> <sl-button size=\"large\" outline pill href=\"https://iconify.design/\" target=\"_blank\">Iconify <iconify-icon icon=\"simple-icons:iconify\" slot=\"prefix\"></iconify-icon></sl-button></div></div><div class=\"flex flex-col items-center justify-center gap-2 row-start-3\"><span class=\"text-black\">Also if you'd like, you can check me out :)</span><div class=\"flex gap-2 items-center justify-center\"><sl-button variant=\"default\" size=\"medium\" outline circle href=\"https://github.com/hastaka/\" target=\"_blank\"><iconify-icon icon=\"simple-icons:github\" label=\"Github\"></iconify-icon></sl-button> <sl-button variant=\"default\" size=\"medium\" outline circle href=\"https://www.linkedin.com/in/ahastak/\" target=\"_blank\"><iconify-icon icon=\"simple-icons:linkedin\" label=\"LinkedIn\"></iconify-icon></sl-button> <sl-button variant=\"default\" size=\"medium\" outline circle href=\"https://hastak.me/\" target=\"_blank\"><sl-icon src=\"/static/images/ah_logo.svg\" label=\"My Website\"></sl-icon></sl-button></div></div></section></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
