{{define "base"}}
<!doctype html>
<html lang='en'>
    <head>
        <!-- <meta Content-Type='text/html' charset='utf-8'> -->
        <meta charset='utf-8'>
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <title>{{template "title" .}}</title>
        <!-- <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.7.1/css/all.css" integrity="sha384-fnmOCqbTlWIlj8LyTjo7mOUStjsKC4pOpQbqyi7RrhN7udi9RwhKkMHpvLbHG9Sr" crossorigin="anonymous"> -->
        <link rel='stylesheet' href='/static/css/bulma.min.css'>
        <link rel='stylesheet' href='/static/fontawesome/css/all.min.css'>
        <!-- <link rel='shortcut icon' href='/static/img/favicon.ico' type='image/x-icon'> -->
        <link rel='stylesheet' href='https://fonts.googleapis.com/css?family=Ubuntu+Mono:400,700'>
    </head>
    <body>
        {{template "nav" .}}
        {{template "body" .}}
        <!-- {{template "footer" .}} -->
    </body>
</html>
{{end}}