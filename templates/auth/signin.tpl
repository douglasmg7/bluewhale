{{template "base" .}}
{{ define "embedded-css"}} {{end}}
{{define "title"}}Autenticação{{end}}

{{define "header"}}
<div class="header">
    <h1>Autenticação</h1>
    <h4></h4>
</div>
{{end}}

{{define "content"}}
<form class="content" action="/auth/signin" method="post">
    <!-- <h1>Autenticação</h2> -->

    <!-- Head messages. -->
    {{if .SuccessMsgHead}} <div> {{.SuccessMsgHead}} </div> {{end}}
    {{if .WarnMsgHead}} <div> {{.WarnMsgHead}} </div> {{end}}

    <!-- email -->
    <input class="{{if .Email.Msg}}is-danger{{end}}" type="text" placeholder="E-mail" id="email" name="email"  value={{.Email.Value}}>
    <span> <i class="fas fa-envelope"></i> </span>
    <p>{{.Email.Msg}}</p>

    <!-- password -->
    <input class="{{if .Password.Msg}}is-danger{{end}}" type="password" placeholder="Senha" id="password" name="password" value={{.Password.Value}}>
    <span> <i class="fas fa-key"></i> </span>
    <p>{{.Password.Msg}}</p>

    <!-- submit -->
    <button type="submit">Entrar</button>

    <!-- reset password -->
    <a href="/auth/reset_password">Esqueceu a senha?</a>

    <!-- signup -->
    <p>Não tem cadastro?<a href="/auth/signup"> Criar cadastro</a></p>

    <!-- Foot messages. -->
    {{if .SuccessMsgFooter}} <div> {{.SuccessMsgFooter}} </div> {{end}}
    {{if .WarnMsgFooter}} <div> {{.WarnMsgFooter}} </div> {{end}}
</form>
{{end}}