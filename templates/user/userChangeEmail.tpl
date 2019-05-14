{{template "base" .}}

{{ define "embedded-css"}}
<style type="text/css">
    .content {
        max-width: var(--width-small);
    }
</style>
{{end}}

{{define "title"}}Alteração do email{{end}}

{{define "content"}}
<form class="content" action="/user/change/email" method="post">
    <h2 class="title">Alteração do email</h2>

    <!-- Email -->
    <label for="email">Novo email</label>
    <input type="text" id="email" name="email"  value={{.Email.Value}}>
    <p class="error"> {{.Email.Msg}} </p>

    <!-- Password -->
    <label for="password">Senha</label>
    <input type="password" id="password" name="password">
    <p class="error"> {{.Password.Msg}} </p>

    <!-- submit -->
    <input type="submit" value="Alterar">
</form>
{{end}}
