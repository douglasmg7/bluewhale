{{ define "embedded-css"}}
<style type="text/css">
    .content {
        margin-top: 2em;
    }
</style>
{{end}}

{{template "base" .}}

{{define "title"}}Esola de vela Ventos Gerais{{end}}

{{define "header"}} {{end}}

{{define "content"}}
<div class="content">
    <!-- Presentation. -->
    <div class="row">
        <div class="column">
            <h1>Escola de Vela Ventos Gerais</h1>
            <h3>Seu primeiro grande passo no mundo da vela, seja ele para navegar outros oceanos ou como esporte.</h3>
        </div>
        <div class="column">
            <img src="/static/img/main1.jpg">
        </div>
    </div>
    <!-- Sub presentation. -->
    <div class="row">
        <div class="column">
            <img src="/static/img/sail_school.jpg" alt="Placeholder image">
            <div>
                <h4>Aula de vela para crianças</h4>
                <p class="card-text text-black-50">Criança também veleja, e os pequenos tem uma classe de vela só para eles: a Optimist!</p>
            </div>
        </div>
        <div class="column">
            <img src="/static/img/sail3.jpg" alt="Placeholder image">
            <div>
                <h4>Aula de vela para adultos</h4>
                <p class="card-text text-black-50">Promovemos a iniciação de adultos na vela em diferentes barcos, para que cada aluno tenha uma rica experiência em diferentes embarcações.</p>
            </div>
        </div>
        <div class="column">
            <img src="/static/img/sail_school2.jpg" alt="Placeholder image">
            <div>
                <h4>Passeio e aluguel de veleiros</h4>
                <p class="card-text text-black-50">É possível desfrutar da vela sem exatamente saber velejar. Nossa escola oferece passeios em veleiros na Lagoa dos Ingleses para até 6 pessoas, em barcos cabinados.</p>
            </div>
        </div>
    </div>
</div>
{{end}}








