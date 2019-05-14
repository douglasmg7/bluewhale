{{template "base" .}}

{{ define "embedded-css"}}
<style type="text/css">
    .content {
        max-width: var(--width-small);
    }
    .title {
    }
    .title + h4 {
        margin-top: 0;
    }
</style>
{{end}}

{{define "title"}}Cadastro{{end}}

{{define "header"}}{{end}}

{{define "content"}}
    <div class="content">
        <h2>Dados da conta</h2>

        <h4 class="title">Nome</h4>
        <h4>{{.Name}}</h4>
        <a href="/user/change-name">Alterar</a>

        <h4 class="title">Email</h4>
        <h4>{{.Email}}</h4>
        <a href="/user/change-email">Alterar</a>

        <h4 class="title">Número de celular</h4>
        <h4>{{if .Mobile}} {{.Mobile}} {{else}} XXXX {{end}}</h4>
        <a href="/user/change-mobile">Alterar</a>

        <h4 class="title">Senha</h4>
        <h4>********</h4>
        <a href="/user/change-password">Alterar</a>

        <h4 class="title">RG</h4>
        <h4>{{.RG}}</h4>
        <a href="/user/change-rg">Alterar</a>

        <h4 class="title">CPF</h4>
        <h4>{{.CPF}}</h4>
        <a href="/user/change-cpf">Alterar</a>

        <h4 class="title">Apagar conta</h4>
        <a href="user/delete-account">Apagar</a>

        <!-- submit -->
        <a class="button" href="/">Sair</a>
    </div>
{{end}}
