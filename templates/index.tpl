{{template "base" .}}

{{define "title"}}Página inicial{{end}}

{{define "main"}}
<style type="text/css">
  .has-rounded-border {
    border-radius: 10px;
  }
  .card-image figure img {
    border-radius: 10px 10px 0 0;
  }
  .card {
    border-radius: 10px
  }
</style>
<div class="main">
  <!-- Presentation. -->
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column">
          <h1 class="title is-1 has-text-grey-dark">Escola de Vela Velejar!</h2>
          <h2 class="subtitle is-3 has-text-grey">Aqui você dará o seu primeiro grande passo no mundo da vela, seja ele para navegar outros oceanos ou como esporte.</h2>
        </div>
        <div class="column">
          <figure class="figure">
            <img class="has-rounded-border" src="/static/img/main1.jpg">
          </figure>
        </div>
      </div>
    </div>
  </section>
  <!-- Sub presentation. -->
  <section class="section">
    <div class="container">
      <div class="columns">
        <div class="column">
          <div class="card">
            <div class="card-image">
              <figure>
                <img src="/static/img/sail_school.jpg" alt="Placeholder image">
              </figure>
            </div>
            <div class="card-content">
              <p class="title is-4">Aula de vela para crianças</h3>
              <p class="subtitle is-6">Criança também veleja, e os pequenos tem uma classe de vela só para eles: a Optimist!</p>
            </div>
          </div>
        </div>
        <div class="column">
          <div class="card">
            <div class="card-image">
              <figure class="image">
                <img src="/static/img/sail3.jpg" alt="Placeholder image">
              </figure>
            </div>
            <div class="card-content">
              <p class="title is-4">Aula de vela para adultos</h3>
              <p class="subtitle is-6">Promovemos a iniciação de adultos na vela em diferentes barcos, para que cada aluno tenha uma rica experiência em diferentes embarcações.</p>
            </div>
          </div>        
        </div>
        <div class="column">
          <div class="card">
            <div class="card-image">
              <figure>
                <img src="/static/img/sail_school2.jpg" alt="Placeholder image">
              </figure>
            </div>
            <div class="card-content">
              <p class="title is-4">Passeio e aluguel de veleiros</h3>
              <p class="subtitle is-6">É possível desfrutar da vela sem exatamente saber velejar. Nossa escola oferece passeios em veleiros na Lagoa dos Ingleses para até 6 pessoas, em barcos cabinados.</p>
            </div>
          </div>        
        </div>
      </div>
    </div>
  </section>
</div>
{{end}}