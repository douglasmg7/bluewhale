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
          <input class="input {{if .Email.Msg}}is-danger{{end}}" type="text" placeholder="E-mail" id="email" name="email"  value={{.Email.Value}}>
          <span class="icon is-small is-left">
            <i class="fas fa-envelope"></i>
          </span>
        </div>
        <p class="has-text-danger has-text-left">{{.Email.Msg}}</p>
      </div>
      <!-- password -->
      <div class="field">
        <div class="control has-icons-left">
          <input class="input {{if .Password.Msg}}is-danger{{end}}" type="password" placeholder="Senha" id="password" name="password" value={{.Password.Value}}>
          <span class="icon is-small is-left">
            <i class="fas fa-key"></i>
          </span>
        </div>
        <p class="has-text-danger has-text-left">{{.Password.Msg}}</p>
      </div>
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