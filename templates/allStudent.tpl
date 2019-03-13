{{template "base" .}}

{{define "title"}}Adicionar aluno{{end}}

{{define "main"}}
<div class="main">
  <section class="section">
    <div class="container">
      <h1 class="title">Alunos</h2>
      {{range .Students}}
        <h2 class="subtitle">
          <a href="/student/id/{{.Id}}">{{.Name}}</a>
        </h2>
      {{end}}
    </div>
  </section>
</div>
{{end}}