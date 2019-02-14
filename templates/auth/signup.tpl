{{template "base" .}}
{{define "title"}}Cadastro{{end}}
{{define "body"}}
  <section class="section">
    <form class="container has-text-centered" action="/auth/signup" method="post" style="max-width:300px;">
      <!-- title -->
      <h1 class="subtitle is-3">Cadastro</h2>
      <!-- name -->
      <div class="field">
        <div class="control has-icons-left">
          <input class="input {{if .Name.Msg}}is-danger{{end}}" type="text" id="name" name="name" placeholder="Nome" value={{.Name.Value}}>
          <span class="icon is-small is-left">
            <i class="fas fa-user"></i>
          </span>
        </div>
        <p class="has-text-danger has-text-left">{{.Name.Msg}}</p>
      </div>
      <!-- email -->
      <div class="field">
        <div class="control has-icons-left">
          <input class="input {{if .Email.Msg}}is-danger{{end}}" type="text" id="email" name="email" placeholder="E-mail" value={{.Email.Value}}>
          <span class="icon is-small is-left">
            <i class="fas fa-envelope"></i>
          </span>
        </div>
        <p class="has-text-danger has-text-left">{{.Email.Msg}}</p>
      </div>
      <!-- password -->
      <div class="field">
        <div class="control has-icons-left">
          <input class="input {{if .Password.Msg}}is-danger{{end}}" type="password" id="password" name="password" placeholder="Senha" value={{.Password.Value}}>
          <span class="icon is-small is-left">
            <i class="fas fa-key"></i>
          </span>
        </div>
        <p class="has-text-danger has-text-left">{{.Password.Msg}}</p>
      </div>
      <!-- confirm password -->
      <div class="field">
        <div class="control has-icons-left">
          <input class="input {{if .PasswordConfirm.Msg}}is-danger{{end}}" type="password" id="passwordConfirm" name="passwordConfirm" placeholder="Confirme a senha" value={{.PasswordConfirm.Value}}>
          <span class="icon is-small is-left">
            <i class="fas fa-check"></i>
          </span>
        </div>
        <p class="has-text-danger has-text-left">{{.PasswordConfirm.Msg}}</p>
      </div>
      <!-- submit -->
      <div class="field">
        <div class="control">
          <button type="submit" class="button is-info is-fullwidth">Cadastrar</button>
        </div>
      </div>
      <!-- success message -->
      {{if .SuccessMsg}}
        <div class="notification is-success">
          {{.SuccessMsg}}
        </div>
      {{end}}
      <!-- warn message -->
      {{if .WarnMsg}}
      <div class="notification is-danger">
        {{.WarnMsg}}
      </div>
      {{end}}
    </form>
  </section>
{{end}}