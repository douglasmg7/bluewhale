{{template "base" .}}

{{ define "embedded-css"}}
<style type="text/css">
    .content {
        max-width: var(--width-small);
    }
</style>
{{end}}

{{define "title"}}Alteração do nome{{end}}

{{define "content"}}
<form class="content" action="/user/change/name" method="post">
    <h2 class="title">Alteração do nome</h2>

    <label for="name">Novo nome</label>
    <input type="text" id="name" name="name"  value={{.Name.Value}}>
    <p class="error"> {{.Name.Msg}} </p>

    <!-- submit -->
    <input type="submit" value="Alterar">
</form>
{{end}}
