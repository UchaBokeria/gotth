// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Letter() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"21\" height=\"21\" viewBox=\"0 0 21 21\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"flex-grow-0 flex-shrink-0 w-5 h-5\" preserveAspectRatio=\"none\"><path d=\"M14.6641 17.6174H6.33073C3.83073 17.6174 2.16406 16.3674 2.16406 13.4507V7.61735C2.16406 4.70068 3.83073 3.45068 6.33073 3.45068H14.6641C17.1641 3.45068 18.8307 4.70068 18.8307 7.61735V13.4507C18.8307 16.3674 17.1641 17.6174 14.6641 17.6174Z\" stroke=\"white\" stroke-width=\"1.2\" stroke-miterlimit=\"10\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path> <path d=\"M14.6693 8.03418L12.0609 10.1175C11.2026 10.8008 9.79427 10.8008 8.93593 10.1175L6.33594 8.03418\" stroke=\"white\" stroke-width=\"1.2\" stroke-miterlimit=\"10\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}