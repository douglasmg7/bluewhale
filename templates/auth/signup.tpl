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
          <input class="input" type="text" id="name" name="name" placeholder="Nome" value={{.Name.Value}}>
          <span class="icon is-small is-left">
            <i class="fas fa-user"></i>
          </span>
        </div>
        <p class="has-text-danger">{{.Name.Msg}}</p>
      </div>
      <!-- email -->
      <div class="field">
        <div class="control has-icons-left">
          <input class="input" type="text" id="email" name="email" placeholder="E-mail" value={{.Email.Value}}>
          <span class="icon is-small is-left">
            <i class="fas fa-envelope"></i>
          </span>
        </div>
        <p class="has-text-danger">{{.Email.Msg}}</p>
      </div>
      <!-- password -->
      <div class="field">
        <div class="control has-icons-left">
          <input class="input" type="password" id="password" name="password" placeholder="Senha" value={{.Password.Value}}>
          <span class="icon is-small is-left">
            <i class="fas fa-key"></i>
          </span>
        </div>
        <p class="has-text-danger">{{.Password.Msg}}</p>
      </div>
      <!-- confirm password -->
      <div class="field">
        <div class="control has-icons-left">
          <input class="input" type="password" id="passwordConfirm" name="passwordConfirm" placeholder="Confirme a senha" value={{.PasswordConfirm.Value}}>
          <span class="icon is-small is-left">
            <i class="fas fa-check"></i>
          </span>
        </div>
        <p class="has-text-danger">{{.PasswordConfirm.Msg}}</p>
      </div>

      <div class="field">
        <div class="control">
          <button type="submit" class="button is-info is-fullwidth">Cadastrar</button>
        </div>
      </div>

      {{if .SuccessMsg}}
        <div class="notification is-success">
          {{.SuccessMsg}}
        </div>
      {{end}}

      {{if .WarnMsg}}
      <div class="notification is-danger">
        {{.WarnMsg}}
      </div>
      {{end}}

    </form>
  </section>
{{end}}