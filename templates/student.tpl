{{template "base" .}}

{{define "title"}}Aluno{{end}}

{{define "body"}}
  <section class="section">
    <h1 class="title">Aluno</h2>
    <div class="container">
      <h2 class="subtitle">{{.Name}}</h2>
      <h2 class="subtitle">{{.Email}}</h2>
      <h2 class="subtitle">{{.Mobile}}</h2>
    </div>
  </section>
{{end}}