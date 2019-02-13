{{template "base" .}}

{{define "title"}}Cadastro{{end}}

{{define "body"}}
  <section class="section">
    <form class="container" action="/auth/signup" method="post" style="max-width:400px;">

      <h1 class="title">Cadastro</h2>

      <div class="field">
        <label for="name" class="label">Nome</label>
        <div class="control">
          <input class="input" type="text" placeholder="" id="name" name="name" value={{.Name.Value}}>
        </div>
        <p class="has-text-danger">{{.Name.Msg}}</p>
      </div>

      <div class="field">
        <label for="email" class="label">E-mail</label>
        <div class="control">
          <input class="input" type="text" placeholder="" id="email" name="email" value={{.Email.Value}}>
        </div>
        <p class="has-text-danger">{{.Email.Msg}}</p>
      </div>

      <div class="field">
        <label for="password" class="label">Senha</label>
        <div class="control">
          <input class="input" type="password" placeholder="" id="password" name="password" value={{.Password.Value}}>
        </div>
        <p class="has-text-danger">{{.Password.Msg}}</p>
      </div>

      <div class="field">
        <label for="passwordConfirm" class="label">Confirme a senha</label>
        <div class="control">
          <input class="input" type="password" placeholder="" id="passwordConfirm" name="passwordConfirm" value={{.PasswordConfirm.Value}}>
        </div>
        <p class="has-text-danger">{{.PasswordConfirm.Msg}}</p>
      </div>

      <div class="field is-grouped">
        <div class="control">
          <button class="button is-link">Cadastrar</button>
        </div>
        <div class="control">
          <button class="button is-text">Cancelar</button>
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