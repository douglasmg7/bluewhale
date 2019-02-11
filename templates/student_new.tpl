{{template "base" .}}

{{define "title"}}Adicionar aluno{{end}}

{{define "body"}}
  <section class="section">
    <form class="container" action="/student-save" method="post">

      <h1 class="title">Adicionar aluno</h2>

      <div class="field">
        <label for="name" class="label">Nome completo</label>
        <div class="control">
          <input class="input" type="text" placeholder="" id="name" name="name" value={{.Name}}>
        </div>
        <p class="help is-danger">{{.NameErr}}</p>
      </div>

      <div class="field">
        <label for="email" class="label">e-mail</label>
        <div class="control">
          <input class="input" type="text" placeholder="" id="email" name="email" value={{.Email}}>
        </div>
        <p class="help is-danger">{{.EmailErr}}</p>
      </div>

      <div class="field">
        <label for="mobile" class="label">Celular</label>
        <div class="control">
          <input class="input" type="text" placeholder="" id="mobile" name="mobile" value={{.Mobile}}>
        </div>
        <p class="help is-danger">{{.MobileErr}}</p>
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