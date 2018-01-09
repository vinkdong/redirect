package server

import (
	"net/http"
	"net/url"
	"fmt"
	template2 "html/template"
	"github.com/containous/traefik/log"
	"bytes"
	"github.com/VinkDong/gox/vtime"
	"regexp"
)

func redirect(w http.ResponseWriter, r *http.Request) {

	parameters := r.URL.Query()
	config := Context.Config

	value := url.Values{}
	tempValue := make(map[string]string)
	for k, v := range parameters {
		vk := config.Static.Parser(k)
		value[config.Static.Parser(vk)] = v
		if len(v) >0 {
			tempValue[vk] = v[0]
		}
	}

	for _, v := range config.Static.Time.Keys {
		vt := config.Static.Time

		if checkSkip(value.Get(v)) {
			tempValue[v] = value.Get(v)
			value.Set(v, value.Get(v))
		} else {
			from := &vtime.Time{
				Format: vt.From.Format,
				Unit:   vt.From.Unit,
				Value:  value.Get(v),
			}
			to := &vtime.Time{
				Format: vt.To.Format,
				TZ:     vt.To.TZ,
			}
			from.Transfer(to)
			tempValue[v] = to.Value
			value.Set(v, to.Value)
		}
	}
	for k, v := range config.Static.Template {
		templ, err := template2.New(k).Parse(v)
		var tpl bytes.Buffer
		if err != nil {
			log.Println(err)
		}
		templ.Execute(&tpl, tempValue)
		value.Add(k, tpl.String())
	}

	destinationMapper := config.Destination

	var b bytes.Buffer

	for k, v := range value {
		b.Write([]byte(fmt.Sprintf("%s=%s&", k, v[0])))
	}

	uri := fmt.Sprintf("%s/?%s", destinationMapper["all"], b.String())
	w.Write([]byte(fmt.Sprintf(`
<script>
window.location = "%s";
</script>
`, uri)))
}

func checkSkip(from string) bool {
	for _,s := range Context.Config.Static.Time.Skip{
		match, _ := regexp.MatchString(s, from)
		if match{
			return match
		}
	}
	return false
}
