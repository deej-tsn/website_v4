// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func SortMenu() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"select_wrap\"><label for=\"sortBy\">Sort By:</label> <select name=\"sort\" id=\"sortBy\" hx-get=\"/blog/posts\" hx-target=\"#loader\" hx-swap=\"innerHTML\" hx-trigger=\" load, change\"><option value=\"newest\" selected>Newest to Oldest</option> <option value=\"oldest\">Oldest to Newest</option></select></div><script type=\"text/javascript\">\n        const sortBy = document.getElementById(\"sortBy\");\n        const localStorageSort = localStorage.getItem(\"sortBy\");\n        if(localStorageSort === null){\n            sortBy.value = \"newest\";\n        }else{\n            sortBy.value = localStorageSort;\n        }\n\n\n        sortBy.addEventListener(\"change\", function() {\n            const newSort = sortBy.value;\n            console.log(newSort)\n            localStorage.setItem(\"sortBy\", newSort);\n        });\n  \n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
