package render

import (
	"github.com/labstack/echo"
	"github.com/hobo-go/echo-mw/multitemplate"
	"path/filepath"
	"strings"
	"html/template"
	"fmt"
	"net/http"
	"errors"
)

func Render() echo.MiddlewareFunc {
	return render()
}

func render() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc{
		return func(c echo.Context) error {
			if err := next(c); err != nil{
				c.Error(err)
			}

			tmpl, context, err := getContext(c)
			fmt.Println(tmpl)
			fmt.Println(context)

			if err == nil{
				c.Render(http.StatusOK, "views/"+tmpl+".html", context)
			} else {
				c.Logger().Errorf("Render Error: %v, tmpl %v, content %v", err, tmpl, context)
			}

			return nil
		}
	}
}

func getContext(c echo.Context) (tmpl string, context map[string]interface{}, err error) {
	tmplName := c.Get("tmpl")
	tmplNameValue, isString := tmplName.(string)
	tmplData := c.Get("data")

	//模板未定义
	if !isString {
		return "", nil, errors.New("No tmpl defined!")
	}

	//公共模板数据
	commonDatas := getCommonContext(c)

	//模板数据
	if tmplData != nil {
		contextData, isMap := tmplData.(map[string]interface{})

		if isMap {
			for key, value := range commonDatas {
				contextData[key] = value
			}
			return tmplNameValue, contextData, nil
		}
	}

	return tmplNameValue, commonDatas, nil
}

func getCommonContext(c echo.Context) map[string]interface{}{
	//例如auth
	return nil
}
func LoadTemplates() echo.Renderer {
	return loadTemplatesDefault("views")
}

func loadTemplatesDefault(templateDir string) *multitemplate.Render {
	r := multitemplate.New()

	layoutDir := templateDir
	layouts, err := filepath.Glob(layoutDir + "*/*.html" )
	fmt.Println("Layouts=", layouts)
	if err != nil {
		panic(err.Error())
	}

	tmpl := template.Must(template.ParseFiles(layouts...))
	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		//files := append(includes, layout)

		tmplName := strings.TrimPrefix(layout, layoutDir)
		tmplName = strings.TrimSuffix(tmplName, ".html")
		//log.DebugPrint("Tmpl add " + tmplName)
		fmt.Println("Tmpl add " + tmplName)
		r.Add(tmplName, tmpl)
	}
	return &r
}