{{template "base" .}}
{{ define "embedded-css"}} {{end}}
{{define "title"}}Cadastro{{end}}

{{define "header"}}
<div class="header">
    <h1>Cadastro</h1>
    <h4></h4>
</div>
{{end}}

{{define "content"}}
    <form class="content" action="/auth/signup" method="post">
        <!-- <h1 class="subtitle is-3">Cadastro</h2> -->

        <!-- name -->
        <input class="{{if .Name.Msg}}is-danger{{end}}" type="text" id="name" name="name" placeholder="Nome" value={{.Name.Value}}>
        <span> <i class="fas fa-user"></i> </span>
        <p>{{.Name.Msg}}</p>

        <!-- email -->
        <input class="{{if .Email.Msg}}is-danger{{end}}" type="text" id="email" name="email" placeholder="E-mail" value={{.Email.Value}}>
        <span> <i class="fas fa-envelope"></i> </span>
        <p>{{.Email.Msg}}</p>

        <!-- password -->
        <input class="{{if .Password.Msg}}is-danger{{end}}" type="password" id="password" name="password" placeholder="Senha" value={{.Password.Value}}>
        <span> <i class="fas fa-key"></i> </span>
        <p class="has-text-danger has-text-left">{{.Password.Msg}}</p>

        <!-- confirm password -->
        <input class="{{if .PasswordConfirm.Msg}}is-danger{{end}}" type="password" id="passwordConfirm" name="passwordConfirm" placeholder="Confirme a senha" value={{.PasswordConfirm.Value}}>
        <span> <i class="fas fa-check"></i> </span>
        <p>{{.PasswordConfirm.Msg}}</p>

        <!-- submit -->
        <button type="submit">Cadastrar</button>

        <!-- Foot message. -->
        {{if .SuccessMsg}} <div> {{.SuccessMsg}} </div> {{end}}
        {{if .WarnMsg}} <div> {{.WarnMsg}} </div> {{end}}
    </form>
{{end}}