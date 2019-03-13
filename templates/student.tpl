{{template "base" .}}
{{define "title"}}Aluno{{end}}
{{define "main"}}
<div class="main">
  <section class="section">
    <div class="container">
      <p class="title">Aluno</p>
      <p class="subtitle">{{.Name}}</p>
      <p class="subtitle">{{.Email}}</p>
      <p class="subtitle">{{.Mobile}}</p>
    </div>
  </section>
</div>
{{end}}