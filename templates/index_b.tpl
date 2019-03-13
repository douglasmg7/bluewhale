{{template "base_b" .}}

{{define "title"}}Página inicial{{end}}

{{define "main_b"}}
<div class="main container">
    <!-- Presentation. -->
    <div class="row my-5">
        <div class="col-sm">
            <h1 class="title is-1 text-dark">Escola de Vela Velejar!</h1>
            <h2 class="subtitle is-3 text-black-50">Aqui você dará o seu primeiro grande passo no mundo da vela, seja ele para navegar outros oceanos ou como esporte.</h2>
        </div>
        <div class="col-sm">
            <img class="img-fluid rounded-lg" src="/static/img/main1.jpg">
        </div>
    </div>
    <!-- Sub presentation. -->
    <div class="row my-5">
        <div class="col-sm mb-5">
            <div class="card shadow">
                <img class="card-img-top" src="/static/img/sail_school.jpg" alt="Placeholder image">
                <div class="card-body">
                    <h4 class="card-title">Aula de vela para crianças</h4>
                    <p class="card-text text-black-50">Criança também veleja, e os pequenos tem uma classe de vela só para eles: a Optimist!</p>
                </div>
            </div>
        </div>
        <div class="col-sm mb-5">
            <div class="card shadow">
                <img class="card-img-top" src="/static/img/sail3.jpg" alt="Placeholder image">
                <div class="card-body">
                    <h4 class="card-title">Aula de vela para adultos</h4>
                    <p class="card-text text-black-50">Promovemos a iniciação de adultos na vela em diferentes barcos, para que cada aluno tenha uma rica experiência em diferentes embarcações.</p>
                </div>
            </div>
        </div>
        <div class="col-sm mb-5">
            <div class="card shadow">
                <img class="card-img-top" src="/static/img/sail_school2.jpg" alt="Placeholder image">
                <div class="card-body">
                    <h4 class="card-title">Passeio e aluguel de veleiros</h4>
                    <p class="card-text text-black-50">É possível desfrutar da vela sem exatamente saber velejar. Nossa escola oferece passeios em veleiros na Lagoa dos Ingleses para até 6 pessoas, em barcos cabinados.</p>
                </div>
            </div>
        </div>
    </div>
</div>
{{end}}



