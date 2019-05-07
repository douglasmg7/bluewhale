{{template "base" .}}

{{ define "embedded-css"}}
<style type="text/css">
    a.reset-pass {
        display: block;
        margin: .2em 0 1em 0;
    }
    p {
        margin-bottom: 0;
    }
</style>
{{end}}

{{define "title"}}Autenticação{{end}}

{{define "header"}}
<div class="header">
    <h1>Entrar</h1>
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
    <label for="email">E-mail</label>
    <input type="text" id="email" name="email"  value={{.Email.Value}}>
    <p class="error"> {{.Email.Msg}} </p>

    <!-- password -->
    <label for="password">Senha</label>
    <input type="password" id="password" name="password" value={{.Password.Value}}>
    <p class="error">{{.Password.Msg}}</p>

    <!-- submit -->
    <input type="submit" value="Entrar">

    <!-- reset password -->
    <a class="reset-pass" href="/auth/reset_password">Esqueceu a senha?</a>

    <!-- signup -->
    <p>Não tem cadastro? </p>
    <a class="signup" href="/auth/signup">Criar cadastro</a>

    <!-- Foot messages. -->
    {{if .SuccessMsgFooter}} <div> {{.SuccessMsgFooter}} </div> {{end}}
    {{if .WarnMsgFooter}} <div> {{.WarnMsgFooter}} </div> {{end}}
</form>
{{end}}
