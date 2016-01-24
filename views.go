package main

const BaseTemplateStr = `
{{ define "base" }}
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Mailhook</title>
    <link href="/assets/css/bootstrap.min.css" rel="stylesheet">
    <link href="/assets/css/codemirror.css" rel="stylesheet">
    <link href="/assets/css/mailhook.css" rel="stylesheet">
    <link href="/assets/css/font-awesome.min.css" rel="stylesheet">
  </head>

  <body>

    <nav class="navbar navbar-inverse navbar-fixed-top">
      <div class="container">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="/">Mailhook</a>
        </div>
        <div id="navbar" class="collapse navbar-collapse">
          <ul class="nav navbar-nav pull-right">
            <li id="ghlogo"><a href="https://github.com/gophergala2016/mailhook"><i class="fa fa-github"></i></a></li>
          </ul>
        </div>
      </div>
    </nav>

    <div class="container">
			{{ template "content" . }}
    </div>
    <script src="/assets/js/jquery.min.js"></script>
    <script src="/assets/js/bootstrap.min.js"></script>
    <script src="/assets/js/codemirror.js"></script>
    <script src="/assets/js/jsmode.js"></script>
    <script src="/assets/js/mailhook.js"></script>
  </body>
</html>
{{ end }}
`

const AdminTemplateStr = `
{{ define "content" }}
<div class="page-header">
  <div class="row">
    <div class="col-lg-11"><h4>Rules</h4></div>
    <div class="col-lg-1"><a href="/new/" class="btn btn-primary">Add Rule</a></div>
  </div>

</div>
{{ if . }}
<table class="table table-hover">
	{{ range $key, $value := . }}
  <tr>
    <td><a href="/view/{{ $value.Id }}">{{ $value.Title }}</td>
		<td class="col-lg-1">
			<a rel="nofollow" href="/edit/{{ $value.Id }}">
				<span class="glyphicon glyphicon-pencil" aria-hidden="true"></span>
			</a>
    </td>
    <td class="col-lg-1">
			<a rel="nofollow" onclick="return confirm('Are you sure?');" href="/delete/{{ $value.Id }}">
				<span class="glyphicon glyphicon-trash" aria-hidden="true"></span>
			</a>
    </td>
	</tr>
	{{ end }}
</table>
{{ else }}
<div class="jumbotron text-center">  <p>You have not created any rule yet.</p> <p><a class="btn btn-primary btn-lg" href="/new/" role="button">Create Rule</a></p> </div>
{{ end }}
{{ end }}
`

const NewTemplateStr = `
{{ define "content" }}
<div class="page-header"><h4>Create Rule</h4></div>
<form class="form-horizontal" action="/create/" method="post">

  <div class="form-group">
    <input type="hidden" value={{.Id}} name="id">
    <label class="col-sm-2 control-label">Title</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" placeholder="Name your rule" name="title" value="{{.Title}}" required>
    </div>
  </div>

  <div class="form-group">
    <label class="col-sm-2 control-label">Script</label>
    <div class="col-sm-10">
      <textarea id="editor" class="form-control" rows="15" name="script">{{.Script}}</textarea>
    </div>
  </div>

  <div class="endpoints">
    <div class="form-group">
      <label class="col-sm-2 control-label">Endpoints</label>
      <div class="col-sm-9">
        <input type="url" class="form-control" name="endpoint_0">
      </div>
      <div class="col-sm-1">
        <a id="add-endpoint" class="btn btn-default"><span class="glyphicon glyphicon-plus" aria-hidden="true"></span></a>
      </div>
    </div>
  </div>

  <div class="form-group">
    <div class="col-sm-offset-2 col-sm-10">
      <input type="submit" class="btn btn-default" value="Save">
      <a class="btn btn-default" href="/">Cancel </a>
    </div>
  </div>

</form>
{{ end }}
`

const EditTemplateStr = `
{{ define "content" }}
<div class="page-header"><h4>Edit Rule</h4></div>
<form class="form-horizontal" action="/update/" method="post">

  <div class="form-group">
    <input type="hidden" value={{.Id}} name="id">
    <label class="col-sm-2 control-label">Title</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" placeholder="Name your rule" name="title" value="{{.Title}}" required>
    </div>
  </div>

  <div class="form-group">
    <label class="col-sm-2 control-label">Script</label>
    <div class="col-sm-10">
      <textarea id="editor" class="form-control" rows="15" name="script">{{.Script}}</textarea>
    </div>
  </div>

  {{ range $key, $value := .Endpoints }}
    <div class="form-group">
      <label class="col-sm-2 control-label">Endpoints</label>
      <div class="col-sm-10">
        <input type="url" class="form-control" name="endpoint_{{$key}}" value="{{$value}}">
      </div>
    </div>
  {{ end }}

  <div class="form-group">
    <div class="col-sm-offset-2 col-sm-10">
      <input type="submit" class="btn btn-default" value="Save">
      <a class="btn btn-default" href="/">Cancel </a>
    </div>
  </div>

</form>
{{ end }}
`

const ViewTemplateStr = `
{{ define "content" }}
<div class="page-header">
  <div class="row">
    <div class="col-lg-11"><h4>{{.Title}}</h4></div>
    <div class="col-lg-1"><a href="/" class="btn btn-primary">All Rules</a></div>
  </div>

</div>
<div class="row">
  <div class="col-lg-6"><pre>{{.Script}}</pre></div>
  <div class="col-lg-6">
    <h5> Endpoints </h5>
    <ul>
    {{ range $key, $value := .Endpoints }}
      <li>{{$value}}</li>
    {{ end }}
    </ul>
  </div>
</div>
{{ end }}
`
