{{template "base" .}}

{{define "title"}}Adicionar aluno{{end}}

{{define "body"}}
  <section class="section">
    <form class="container" action="/student/new" method="post">

      <h1 class="title">Adicionar aluno</h2>

      <div class="field">
        <label for="name" class="label">Nome completo</label>
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
        <label for="mobile" class="label">Celular</label>
        <div class="control">
          <input class="input" type="text" placeholder="" id="mobile" name="mobile" value={{.Mobile.Value}}>
        </div>
        <p class="help is-danger">{{.Mobile.Msg}}</p>
      </div>

      <div class="field is-grouped">
        <div class="control">
          <button class="button is-link">Adicionar</button>
        </div>
        <div class="control">
          <button class="button is-text">Cancelar</button>
        </div>
      </div>

    </form>
  </section>
{{end}}