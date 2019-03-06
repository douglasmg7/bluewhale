{{template "base" .}}

{{define "title"}}Adicionar aluno{{end}}

{{define "body"}}
  <section class="section">
    <h1 class="title">Alunos</h2>
    <div class="container">
      {{range .Names}}
        <h2 class="subtitle">
          <a href="">{{.}}</a>
        </h2>
      {{end}}
    </div>
  </section>
{{end}}