{{template "base" .}}

{{define "title"}}Aguardando confirmação{{end}}

{{define "body"}}
  <section class="section">
    <form class="container" action="/student/save" method="post" style="max-width:400px;">

      <h1 class="title">Solicitação</h2>


      <div class="field is-grouped">
        <div class="control">
          <button class="button is-link">Cadastrar</button>
        </div>
      </div>

    </form>
  </section>
{{end}}