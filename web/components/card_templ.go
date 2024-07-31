// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

type imgInfo struct {
	src string
	alt string
}

type card struct {
	cardTitle string
	cardInfo  string
	image     imgInfo
}

var cards = []card{
	{
		cardTitle: "Student Visa",
		cardInfo:  "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Laboriosam numquam accusantium saepe aut facere a, ab expedita minus cum ipsa velit beatae illum, autem, quis placeat doloremque error debitis. Obcaecati eveniet repudiandae illum iusto, unde doloremque architecto ut porro beatae perspiciatis accusantium sapiente blanditiis, earum",
		image: imgInfo{
			src: "/static/media/images/main-page/temp-student-visa.webp",
			alt: "Student Visa Image",
		},
	},
	{
		cardTitle: "Visitor Visa",
		cardInfo:  "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Laboriosam numquam accusantium saepe aut facere a, ab expedita minus cum ipsa velit beatae illum, autem, quis placeat doloremque error debitis. Obcaecati eveniet repudiandae illum iusto, unde doloremque architecto ut porro beatae perspiciatis accusantium sapiente blanditiis, earum",
		image: imgInfo{
			src: "/static/media/images/main-page/temp-visitor-visa.webp",
			alt: "Student Visa Image",
		},
	},

	{
		cardTitle: "Student Visa",
		cardInfo:  "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Laboriosam numquam accusantium saepe aut facere a, ab expedita minus cum ipsa velit beatae illum, autem, quis placeat doloremque error debitis. Obcaecati eveniet repudiandae illum iusto, unde doloremque architecto ut porro beatae perspiciatis accusantium sapiente blanditiis, earum",
		image: imgInfo{
			src: "/static/media/images/main-page/temp-student-visa.webp",
			alt: "Student Visa Image",
		},
	},
	{
		cardTitle: "Visitor Visa",
		cardInfo:  "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Laboriosam numquam accusantium saepe aut facere a, ab expedita minus cum ipsa velit beatae illum, autem, quis placeat doloremque error debitis. Obcaecati eveniet repudiandae illum iusto, unde doloremque architecto ut porro beatae perspiciatis accusantium sapiente blanditiis, earum",
		image: imgInfo{
			src: "/static/media/images/main-page/temp-visitor-visa.webp",
			alt: "Student Visa Image",
		},
	},
}

func Cards() templ.Component {
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
		for _, card := range cards {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card-wrapper\"><div id=\"card-img-container\" class=\"blurred-img\"><img class=\"rounded-lg lg:rounded-xl\" src=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(card.image.src)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/components/card.templ`, Line: 54, Col: 62}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(card.image.alt)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/components/card.templ`, Line: 54, Col: 85}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"720\" height=\"720\" decoding=\"async\"></div><div id=\"card-info-container\" class=\"w-4/5 md:max-w-[calc(min(100%,34rem))]\"><h3 class=\"md:text-left pb-6 pt-6 md:pt-0\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(card.cardTitle)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/components/card.templ`, Line: 57, Col: 63}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3><p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(card.cardInfo)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/components/card.templ`, Line: 58, Col: 22}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}
