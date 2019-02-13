{{template "base" .}}

{{define "title"}}Cadastrar{{end}}

{{define "body"}}
  <section class="section">
    <form class="container" action="/student/save" method="post" style="max-width:400px;">

      <h1 class="title">Cadastrar</h2>

      <div class="field">
        <label for="name" class="label">Nome</label>
        <div class="control">
          <input class="input" type="text" placeholder="" id="name" name="name" value={{.Name.Value}}>
        </div>
        <p class="help is-danger">{{.Name.Msg}}</p>
      </div>

      <div class="field">
        <label for="email" class="label">E-mail</label>
        <div class="control">
          <input class="input" type="text" placeholder="" id="email" name="email" value={{.Email.Value}}>
        </div>
        <p class="help is-danger">{{.Email.Msg}}</p>
      </div>

      <div class="field">
        <label for="password" class="label">Senha</label>
        <div class="control">
          <input class="input" type="password" placeholder="" id="password" name="password" value={{.Password.Value}}>
        </div>
        <p class="help is-danger">{{.Password.Msg}}</p>
      </div>

      <div class="field">
        <label for="passwordConfirm" class="label">Confirme a senha</label>
        <div class="control">
          <input class="input" type="password" placeholder="" id="passwordConfirm" name="passwordConfirm" value={{.PasswordConfirm.Value}}>
        </div>
        <p class="help is-danger">{{.PasswordConfirm.Msg}}</p>
      </div>

      <div class="field is-grouped">
        <div class="control">
          <button class="button is-link">Cadastrar</button>
        </div>
        <div class="control">
          <button class="button is-text">Cancelar</button>
        </div>
      </div>

    </form>
  </section>
{{end}}