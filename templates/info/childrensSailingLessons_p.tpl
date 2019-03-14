{{template "base_p" .}}
{{define "title"}}Aulas de vela para crianças{{end}}
{{define "main"}}
<style type="text/css">
    body {
        color: #777;
    }
    .header {
        margin: 0;
        padding-top: 2.5em;
        text-align: center;
        border-bottom: 1px solid #eee;
    }
    .header h1, .header h2 {
        font-weight: 500;
    }
    .header h1 {
        font-size: 3em;
        color: #333;
        margin: .2em 0;
    }
    .header h2 {
        color: #ccc;
        margin-top: 0;
    }
    .content {
        max-width: 800px;
        margin: 0 auto;
        padding: 0 2em;
        margin-bottom: 50px;
        line-height: 1.6em;
    }
    .content h2 {
        color: #888;
        font-weight: 500;
        margin: 50px 0 20px 0;
    }
</style>
<div id="layout">
    <div id="menu">
        
    </div>
    <div id="main">
        <div class="header">
            <h1>Escolinha de Optimist</h1>
            <h2>Um barco para crianças e jovens</h2>
        </div>
        <div class="content">
            <h2>O barco</h2>

            <p> Criança também veleja, e os pequenos tem uma classe de vela só para eles: a Optimist! </p>

            <p> 
                Ensinamos a arte de velejar de forma lúdica, trabalhando o respeito ao meio ambiente, 
                a competição saudável e o companheirismo.
                Brincando, as crianças aprendem a reconhecer as partes do barco, a observar o vento
                e a definir estratégias para velejar.
            </p>

            <h2>Benefícios da prática de vela</h2>
            <ul>
                <li>Melhora a autoconfiança</li>
                <li>Trabalha a disciplina e concentração</li>
                <li>Estimula a responsabilidade</li>
                <li>Desenvolve o trabalho em equipe</li>
                <li>Incentiva o respeito ao meio ambiente</li>
                <li>Trabalha a concentração e tomada de decisão</li>
                <li>Desenvolve a coordenação motora, o senso de direção e a noção espacial</li>
            </ul>
          
            <p>Agende uma aula experimental e conte conosco!</p>
        </div>
    </div>
</div>

{{end}}