{{template "base" .}}
{{define "title"}}Autenticação{{end}}
{{define "body"}}
  <section class="section">
    <form class="container has-text-centered" action="/auth/signin" method="post" style="max-width:300px;">
      <!-- title -->
      <h1 class="subtitle is-3">Autenticação</h2>
      <!-- message success head-->
      {{if .SuccessMsgHead}}
        <div class="notification is-success">
          {{.SuccessMsgHead}}
        </div>
      {{end}}
      <!-- message warn head-->
      {{if .WarnMsgHead}}
        <div class="notification is-danger">
          {{.WarnMsgHead}}
        </div>
      {{end}}
      <!-- email -->
      <div class="field">
        <div class="control has-icons-left">
          <input class="input {{if .Name.Msg}}is-danger{{end}}" type="text" placeholder="E-mail" id="email" name="email"  value={{.Email.Value}}>
          <span class="icon is-small is-left">
            <i class="fas fa-envelope"></i>
          </span>
        </div>
        <p class="help is-danger">{{.Email.Msg}}</p>
      </div>
      <!-- password -->
      <div class="field">
        <div class="control has-icons-left">
          <input class="input {{if .Name.Msg}}is-danger{{end}}" type="text" placeholder="Senha" id="password" name="password" value={{.Password.Value}}>
          <span class="icon is-small is-left">
            <i class="fas fa-key"></i>
          </span>
        </div>
      </div>
      <p class="help is-danger">{{.Password.Msg}}</p>
      <!-- submit -->
      <div class="field">
        <div class="control">
          <button type="submit" class="button is-info is-fullwidth">Entrar</button>
        </div>
      </div>
      <!-- reset password -->
      <div class="field">
        <div class="control has-text-centered">
          <a href="/auth/reset_password">Esqueceu a senha?</a>
        </div>
      </div>
      <!-- signup -->
      <div class="field">
        <div class="control has-text-centered">
          <p>Não tem cadastro?<a href="/auth/signup"> Criar cadastro</a></p>
        </div>
      </div>
      <!-- message success footer-->
      {{if .SuccessMsgFooter}}
        <div class="notification is-success">
          {{.SuccessMsgFooter}}
        </div>
      {{end}}
      <!-- message warn footer-->
      {{if .WarnMsgFooter}}
        <div class="notification is-danger">
          {{.WarnMsgFooter}}
        </div>
      {{end}}
    </form>
  </section>
{{end}}