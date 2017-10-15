package main

const defaultTemplate = `<!doctype html>
<html>
  <head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>{{.Title}}</title>
  </head>
  <body>
    <section>
      <button onclick="external.invoke_('back')">Back</button>
      <button onclick="external.invoke_('forward')">Forward</button>
      <button onclick="external.invoke_('reload')">Reload</button>
      <button onclick="external.invoke_('home')">Home</button>
      <input id="uri" type="text" />
      <button onclick="external.invoke_('open:'+document.getElementById('uri').value)">
        Go
      </button>
    </section>
    <section>
      <pre>
      {{range .Lines}} {{if .Link}}({{.Type}}) <a class="{{ .Type }}" href="{{.Link}}">{{.Text}}</a>{{else}}      {{.Text}}{{end}}
      {{end}}</pre>
    </section>
<script src="https://code.jquery.com/jquery-3.1.0.slim.min.js" integrity="sha256-cRpWjoSOw5KcyIOaZNo4i6fZ9tKPhYYb6i5T9RSVJG8=" crossorigin="anonymous"></script>
<script type="text/javascript">
$(document).ready(function () {
  $(".QRY").click(function (e) {
  e.preventDefault();
  var query = prompt("Please enter required input: ", "");
  if (query != null) {
    window.location = e.target.href + "?" + query;
  }
  });
});
</script>
</body>
</html>`
