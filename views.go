package main

const BaseTemplateStr = `
{{ define "base" }}
<html>
  <body>
	{{ template "content" . }}
  </body>
</html>
{{ end }}
`

const AdminTemplateStr = `
{{ define "content" }}
	<ul>
	{{ range $key, $value := . }}
	    <li> <a href="/view/{{ $value.Id }}">{{ $value.Title }} </a> | 
                 <a href="/edit/{{ $value.Id }}"> edit </a>| 
	         <a rel="nofollow" onclick="return confirm('Are you sure?');" href="/delete/{{ $value.Id }}">delete</a> </li>
	{{ end }}
	</ul>
{{ end }}
`
